package workwx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/mozillazg/go-httpheader"
	"github.com/pkg/errors"
)

// WorkWX 企业微信客户端
type WorkWX struct {
	opts options
	// CorpID 企业 ID，必填
	CorpID string
}

// App 企业微信客户端（分应用）
type App struct {
	*WorkWX
	// CorpSecret 应用的凭证密钥，必填
	CorpSecret string
	// AgentID 应用 ID，必填
	AgentID                int64
	accessToken            *token
	jsapiTicket            *token
	jsapiTicketAgentConfig *token
}

// New 构造一个 WorkWX 客户端对象，需要提供企业 ID
func New(corpID string, opts ...CtorOption) *WorkWX {
	optionsObj := defaultOptions()

	for _, o := range opts {
		o.applyTo(&optionsObj)
	}

	return &WorkWX{
		opts:   optionsObj,
		CorpID: corpID,
	}
}

// WithApp 构造本企业下某自建 app 的客户端
func (c *WorkWX) WithApp(corpSecret string, agentID int64) *App {
	app := App{
		WorkWX:                 c,
		CorpSecret:             corpSecret,
		AgentID:                agentID,
		accessToken:            &token{mutex: &sync.RWMutex{}},
		jsapiTicket:            &token{mutex: &sync.RWMutex{}},
		jsapiTicketAgentConfig: &token{mutex: &sync.RWMutex{}},
	}
	app.accessToken.setGetTokenFunc(app.getAccessToken)
	app.jsapiTicket.setGetTokenFunc(app.getJSAPITicket)
	app.jsapiTicketAgentConfig.setGetTokenFunc(app.getJSAPITicketAgentConfig)
	app.SpawnAccessTokenRefresher()
	return &app
}

func (c *App) composeWXApiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	base, err := url.Parse(c.opts.WxAPIHost)
	if err != nil {
		// TODO: error_chain
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", c.opts.WxAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

func (c *App) composeWXURLWithToken(path string, req interface{}, withAccessToken bool) *url.URL {
	wxApiURL := c.composeWXApiURL(path, req)

	if !withAccessToken {
		return wxApiURL
	}

	q := wxApiURL.Query()
	q.Set("access_token", c.accessToken.getToken())
	wxApiURL.RawQuery = q.Encode()

	return wxApiURL
}

func (c *App) executeWXApiGet(path string, req urlValuer, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	//defer logger.Debugf("url: %s, req: %+v, resp: %+v", urlStr, req, objResp)

	resp, err := c.opts.restyCli.R().Get(urlStr)
	if err != nil {
		return err
	}

	bodyResp := resp.Body()
	err = json.Unmarshal(bodyResp, &objResp)
	return err
}

func (c *App) execGet(path string, req urlValuer, withAccessToken bool) (hrsp *http.Response, err error) {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	//defer logger.Debugf("url: %s, req: %+v, resp: %+v", urlStr, req, objResp)

	resp, err := c.opts.restyCli.R().SetDoNotParseResponse(true).Get(urlStr)
	if err != nil {
		return nil, err
	}

	return resp.RawResponse, nil
}

// 微信端接收的参数中一个数组里包含有多种类型，强类型语言无法支持，只能在前端拼接成str直接传到wx
func (c *App) executeWXApiJSONPostWithBytesReq(path string, req []byte, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	// resp, err := c.opts.HTTP.Post(urlStr, "application/json", bytes.NewReader(req))
	resp, err := c.opts.restyCli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		Post(urlStr)
	if err != nil {
		// TODO: error_chain
		return err
	}

	err = json.Unmarshal(resp.Body(), &objResp)

	return err
}

func (c *App) executeWXApiJSONPost(path string, req bodyer, objResp interface{}, withAccessToken bool) error {
	// defer util.FuncTracer("path", path, "req", req, "resp", objResp)()
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	//defer logger.Debugf("url: %s, req: %+v, resp: %+v", urlStr, req, objResp)

	body, err := req.intoBody()
	if err != nil {
		// TODO: error_chain
		return err
	}

	resp, err := c.opts.restyCli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(urlStr)
	if err != nil {
		// TODO: error_chain
		return err
	}

	err = json.Unmarshal(resp.Body(), &objResp)
	return err
}

func (c *App) executeWXApiHead(ctx context.Context, path string, req urlValuer, objResp interface{}, withAccessToken bool) (mediaInfo MediaInfoRsp, err error) {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	resp, err := c.opts.restyCli.R().
		SetContext(ctx).
		Head(urlStr)

	if err != nil {
		return
	}

	err = httpheader.Decode(resp.Header(), &mediaInfo)
	if err != nil {
		return
	}
	if resp.StatusCode() != http.StatusOK {
		err = json.Unmarshal(resp.Body(), &objResp)
	}
	return
}

func (c *App) executeWXApiMediaUpload(path string, req mediaUploader, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	//defer logger.Debugf("url: %s, req: %+v, resp: %+v", urlStr, req, objResp)

	m := req.getMedia()
	resp, err := c.opts.restyCli.R().
		SetFileReader("media", m.filename, m.stream).
		Post(urlStr)
	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(resp.Body(), &objResp)
	return err
}

func (c *App) GetToken() (token string, err error) {
	token = c.accessToken.getToken()
	if token == "" {
		err = errors.New("invalid conf")
		return
	}
	return
}
