package workwx

import (
	"encoding/json"
	"net/url"
)

type accessTokenReq struct {
	CorpID     string
	CorpSecret string
}

var _ urlValuer = accessTokenReq{}

func (x accessTokenReq) intoURLValues() url.Values {
	return url.Values{
		"corpid":     {x.CorpID},
		"corpsecret": {x.CorpSecret},
	}
}

type CommonResp struct {
	ErrCode int64  `json:"errcode" header:"Error-Code"`
	ErrMsg  string `json:"errmsg" header:"Error-Msg"`
}

// IsOK 响应体是否为一次成功请求的响应
//
// 实现依据: https://work.weixin.qq.com/api/doc#10013
//
// > 企业微信所有接口，返回包里都有errcode、errmsg。
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *CommonResp) IsOK() bool {
	return x.ErrCode == 0
}

func (x *CommonResp) TryIntoErr() *ClientError {
	if x.IsOK() {
		return nil
	}

	return &ClientError{
		Code: x.ErrCode,
		Msg:  x.ErrMsg,
	}
}

type intoBodyer[T any] struct {
	data T
}

func newIntoBodyer[T any](data T) intoBodyer[T] {
	return intoBodyer[T]{data: data}
}
func (i intoBodyer[T]) intoBody() ([]byte, error) {
	result, err := json.Marshal(i.data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type accessTokenResp struct {
	CommonResp

	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}

type jsAPITicketAgentConfigReq struct{}

var _ urlValuer = jsAPITicketAgentConfigReq{}

func (x jsAPITicketAgentConfigReq) intoURLValues() url.Values {
	return url.Values{
		"type": {"agent_config"},
	}
}

type jsAPITicketReq struct{}

var _ urlValuer = jsAPITicketReq{}

func (x jsAPITicketReq) intoURLValues() url.Values {
	return url.Values{}
}

type jsAPITicketResp struct {
	CommonResp

	Ticket        string `json:"ticket"`
	ExpiresInSecs int64  `json:"expires_in"`
}

type suiteAccessTokenReq struct {
	SuiteId     string `json:"suite_id"`
	SuiteSecret string `json:"suite_secret"`
	SuiteTicket string `json:"suite_ticket"`
}

type suiteAccessTokenResp struct {
	CommonResp

	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresInSecs    int64  `json:"expires_in"`
}

type providerAccessTokenReq struct {
	CorpId         string `json:"corpid"`
	ProviderSecret string `json:"provider_secret"`
}

type providerAccessTokenResp struct {
	ProviderAccessToken string `json:"provider_access_token"`
	ExpiresIn           int    `json:"expires_in"`
}
