package workwx

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/loojee/workwx/internal/lowlevel/encryptor"
	"github.com/loojee/workwx/internal/lowlevel/envelope"
	"github.com/loojee/workwx/internal/lowlevel/httpapi"
	"github.com/loojee/workwx/internal/lowlevel/signature"
)

type CallBackHandler struct {
	token     string
	encryptor *encryptor.WorkWXEncryptor
	ep        *envelope.Processor
}

func NewCBHandler(token string, encodingAESKey string) (*CallBackHandler, error) {
	enc, err := encryptor.NewWorkWXEncryptor(encodingAESKey)
	if err != nil {
		return nil, err
	}

	ep, err := envelope.NewProcessor(token, encodingAESKey)
	if err != nil {
		return nil, err
	}

	return &CallBackHandler{token: token, encryptor: enc, ep: ep}, nil
}

func (cb *CallBackHandler) GetCallBackMsg(r *http.Request) (*RxMessage, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusInternalServerError)
		return nil, err
	}

	//logger.Debugf("read body : %+v", string(body))

	// 验签
	// 解析Xml
	ev, err := cb.ep.HandleIncomingMsg(r.URL, body)
	if err != nil {
		return nil, err
	}

	message, err := fromEnvelope(ev.Msg)
	if err != nil {
		return nil, err
	}

	fmt.Println(message)
	return message, nil
}

// EchoTestHandler
// wx后台配置服务器ip，回显
func (cb *CallBackHandler) EchoTestHandler(rw http.ResponseWriter, r *http.Request) {
	url := r.URL

	if !signature.VerifyHTTPRequestSignature(cb.token, url, "") {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	adapter := httpapi.URLValuesForEchoTestAPI(url.Query())
	args, err := adapter.ToEchoTestAPIArgs()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := cb.encryptor.Decrypt([]byte(args.EchoStr))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(payload.Msg)
}
