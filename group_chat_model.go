package workwx

import (
	"encoding/json"
)

type GetGroupChatReq struct {
	ChatId   string `json:"chat_id"`
	NeedName int8   `json:"need_name"`
}

// GetGroupChatReq 获取客户群详情请求
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92122#获取客户群详情
var _ bodyer = GetGroupChatReq{}

func (x GetGroupChatReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetGroupChatResp 获取客户群详情响应
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92122#获取客户群详情
type GetGroupChatResp struct {
	CommonResp
	GroupChat `json:"group_chat"`
}

type GroupChat struct {
	AdminList []struct {
		Userid string `json:"userid"`
	} `json:"admin_list"`
	ChatID     string `json:"chat_id"`
	CreateTime int    `json:"create_time"`
	MemberList []struct {
		Invitor struct {
			Userid string `json:"userid"`
		} `json:"invitor"`
		JoinScene int    `json:"join_scene"`
		JoinTime  int    `json:"join_time"`
		Type      int    `json:"type"`
		Unionid   string `json:"unionid"`
		Userid    string `json:"userid"`
		Name      string `json:"name"`
		// 在群聊中的名称
		GroupNickname string `json:"group_nickname"`
		State         string `json:"state"`
	} `json:"member_list"`
	Name   string `json:"name"`
	Notice string `json:"notice"`
	Owner  string `json:"owner"`
}

var _ bodyer = GetGroupChatResp{}

func (x GetGroupChatResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GroupChatJoinWayScene 群活码场景
// 1-在小程序中联系
// 2-通过二维码联系
type GroupChatJoinWayScene int

const (
	// GroupChatJoinWaySceneMicroApp 在小程序中联系
	GroupChatJoinWaySceneMicroApp GroupChatJoinWayScene = 1
	// GroupChatJoinWaySceneQrcode 通过二维码联系
	GroupChatJoinWaySceneQrcode GroupChatJoinWayScene = 2
)

type GroupChatJoinWayAutoCreate int

const (
	GroupChatJoinWayAutoCreateFalse GroupChatJoinWayAutoCreate = 0
	GroupChatJoinWayAutoCreateTrue  GroupChatJoinWayAutoCreate = 1
)

type GroupChatAddJoinWayReq struct {
	Scene          GroupChatJoinWayScene      `json:"scene"`
	Remark         string                     `json:"remark"`
	AutoCreateRoom GroupChatJoinWayAutoCreate `json:"auto_create_room"`
	RoomBaseName   string                     `json:"room_base_name"`
	RoomBaseId     int                        `json:"room_base_id"`
	ChatIdList     []string                   `json:"chat_id_list"`
	State          string                     `json:"state"`
}

var _ bodyer = GroupChatAddJoinWayReq{}

func (x GroupChatAddJoinWayReq) intoBody() ([]byte, error) {
	return intoJsonBody(x)
}

type GroupChatAddJoinWayResp struct {
	CommonResp
	ConfigId string `json:"config_id"`
}

type GroupChatGetJoinWayReq struct {
	ConfigId string `json:"config_id"`
}

var _ bodyer = GroupChatGetJoinWayReq{}

func (x GroupChatGetJoinWayReq) intoBody() ([]byte, error) {
	return intoJsonBody(x)
}

type GroupChatGetJoinWay struct {
	ConfigId       string                     `json:"config_id"`
	Scene          int                        `json:"scene"`
	Remark         string                     `json:"remark"`
	AutoCreateRoom GroupChatJoinWayAutoCreate `json:"auto_create_room"`
	RoomBaseName   string                     `json:"room_base_name"`
	RoomBaseId     int                        `json:"room_base_id"`
	ChatIdList     []string                   `json:"chat_id_list"`
	QrCode         string                     `json:"qr_code"`
	State          string                     `json:"state"`
}

type GroupChatGetJoinWayResp struct {
	CommonResp
	JoinWay GroupChatGetJoinWay `json:"join_way"`
}

type GroupChatUpdateJoinWayReq struct {
	ConfigId       string                     `json:"config_id"`
	Scene          int                        `json:"scene"`
	Remark         string                     `json:"remark"`
	AutoCreateRoom GroupChatJoinWayAutoCreate `json:"auto_create_room"`
	RoomBaseName   string                     `json:"room_base_name"`
	RoomBaseId     int                        `json:"room_base_id"`
	ChatIdList     []string                   `json:"chat_id_list"`
	State          string                     `json:"state"`
}

func (x GroupChatUpdateJoinWayReq) intoBody() ([]byte, error) {
	return intoJsonBody(x)
}

type GroupChatDelJoinWayReq struct {
	ConfigId string `json:"config_id"`
}

func (x GroupChatDelJoinWayReq) intoBody() ([]byte, error) {
	return intoJsonBody(x)
}
