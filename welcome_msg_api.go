package workwx

import (
	"encoding/json"
	"time"
)

type SendWelcomeMsgReq struct {
	//附件，最多可添加9个附件
	Attachments []Attachments `json:"attachments,omitempty"`
	Text        Text          `json:"text"`
	// WelcomeCode 通过 添加外部联系人事件 推送给企业的发送欢迎语的凭证，有效期为 20秒，必填
	WelcomeCode string `json:"welcome_code"`
}

var _ bodyer = SendWelcomeMsgReq{}

func (x SendWelcomeMsgReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// sendWelcomeMsgResp 发送新客户欢迎语响应
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92137#发送新客户欢迎语
type sendWelcomeMsgResp struct {
	CommonResp
}

var _ bodyer = sendWelcomeMsgResp{}

func (x sendWelcomeMsgResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execSendWelcomeMsg 发送新客户欢迎语
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92137#发送新客户欢迎语
func (c *App) execSendWelcomeMsg(req SendWelcomeMsgReq) (sendWelcomeMsgResp, error) {
	var resp sendWelcomeMsgResp

	var i = 0
	for i < 3 {
		i++
		err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/send_welcome_msg", req, &resp, true)
		if err != nil {
			return sendWelcomeMsgResp{}, err
		}

		/*
			https://developer.work.weixin.qq.com/document/path/92137
			如果欢迎语已成功下发给客户，应用再调用该接口时，将返回41051（externaluser has started chatting）。
			每次添加新客户时可能有多个企业自建应用/第三方应用收到带有welcome_code的回调事件，当多个应用同时调用发送欢迎语接口时，
			仅最先调用接口的应用可以成功，其他应用返回41096（welcome msg is being distributed）错误。
			请注意，返回41096错误码表示当前已有其他应用正在发送欢迎语，不代表欢迎语已经下发成功，存在下发失败的可能性，故调用方可以重试；
			如果返回了41051，则代表欢迎语已成功下发，无需重试。为了保证用户体验，请根据实际业务需求，合理配置应用可见范围。
		*/
		if resp.ErrCode == 41096 {
			time.Sleep(time.Millisecond * 100)
			continue
		} else if resp.ErrCode == 41051 {
			break
		}

		if bizErr := resp.TryIntoErr(); bizErr != nil {
			return sendWelcomeMsgResp{}, bizErr
		}

		if err == nil {
			break
		}
	}

	return resp, nil

}
