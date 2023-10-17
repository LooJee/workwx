package customer_service

import (
	"github.com/mitchellh/mapstructure"
)

type MessageType string // 消息类型

const (
	MessageTypeText                MessageType = "text"                  // 文本消息
	MessageTypeImage               MessageType = "image"                 // 图片消息
	MessageTypeVoice               MessageType = "voice"                 // 语音消息
	MessageTypeVideo               MessageType = "video"                 // 视频消息
	MessageTypeFile                MessageType = "file"                  // 文件消息
	MessageTypeLocation            MessageType = "location"              // 位置消息
	MessageTypeLink                MessageType = "link"                  // 链接消息
	MessageTypeBusinessCard        MessageType = "business_card"         // 名片消息
	MessageTypeMiniProgram         MessageType = "miniprogram"           // 小程序消息
	MessageTypeMsgMenu             MessageType = "msgmenu"               // 菜单消息
	MessageTypeChannelsShopProduct MessageType = "channels_shop_product" // 视频号商品消息
	MessageTypeChannelsShopOrder   MessageType = "channels_shop_order"   // 视频号订单消息
	MessageTypeMergedMsg           MessageType = "merged_msg"            // 聊天记录消息
	MessageTypeChannels            MessageType = "channels"              // 视频号消息
	MessageTypeMeeting             MessageType = "meeting"               // 会议消息
	MessageTypeSchedule            MessageType = "schedule"              // 日程消息
	MessageTypeEvent               MessageType = "event"                 // 事件消息
)

type MsgItem struct {
	MsgType             MessageType                 `json:"msgtype"`
	Text                *TextMessage                `json:"text,omitempty"`
	Image               *ImageMessage               `json:"image,omitempty"`
	Voice               *VoiceMessage               `json:"voice,omitempty"`
	Video               *VideoMessage               `json:"video,omitempty"`
	File                *FileMessage                `json:"file,omitempty"`
	Location            *LocationMessage            `json:"location,omitempty"`
	Link                *LinkMessage                `json:"link,omitempty"`
	BusinessCard        *BusinessCardMessage        `json:"business_card,omitempty"`
	MiniProgram         *MiniProgramMessage         `json:"miniprogram,omitempty"`
	MsgMenu             *MsgMenuMessage             `json:"msgmenu,omitempty"`
	ChannelsShopProduct *ChannelsShopProductMessage `json:"channels_shop_product,omitempty"`
	ChannelsShopOrder   *ChannelsShopOrderMessage   `json:"channels_shop_order,omitempty"`
	MergedMsg           *MergedMsgMessage           `json:"merged_msg,omitempty"`
	Event               *EventMessage               `json:"event"`
}

type TextMessage struct {
	Content string `json:"content"`           // 文本内容
	MenuId  string `json:"menu_id,omitempty"` // 客户点击菜单消息，触发的回复消息中附带的菜单ID
}

type ImageMessage struct {
	MediaId string `json:"media_id"` // 图片消息媒体ID
}

type VoiceMessage struct {
	MediaId string `json:"media_id"` // 语音消息媒体ID
}

type VideoMessage struct {
	MediaId string `json:"media_id"` // 视频消息媒体ID
}

type FileMessage struct {
	MediaId string `json:"media_id"` // 文件消息媒体ID
}

type LocationMessage struct {
	Latitude  float64 `json:"latitude"`  // 地理位置纬度
	Longitude float64 `json:"longitude"` // 地理位置经度
	Name      string  `json:"name"`      // 位置名
	Address   string  `json:"address"`   // 地址
}

type LinkMessage struct {
	Title        string `json:"title"`             // 标题
	Description  string `json:"description"`       // 描述
	Url          string `json:"url"`               // 链接
	PicUrl       string `json:"pic_url,omitempty"` // 缩略图链接
	ThumbMediaId string `json:"thumb_media_id"`    // 缩略图媒体ID
}

type BusinessCardMessage struct {
	UserId string `json:"user_id"` // 用户ID
}

type MiniProgramMessage struct {
	AppId        string `json:"appid"`          // 小程序id
	Title        string `json:"title"`          // 标题
	PagePath     string `json:"pagepath"`       // 页面路径
	ThumbMediaId string `json:"thumb_media_id"` // 缩略图媒体ID
}

type MsgMenuType string

const (
	MsgMenuTypeClick       MsgMenuType = "click"       // 回复菜单
	MsgMenuTypeView        MsgMenuType = "view"        // 超链接菜单
	MsgMenuTypeMiniprogram MsgMenuType = "miniprogram" // 小程序菜单
)

type MsgMenuClick struct {
	Id      string `json:"id"`      // 菜单ID
	Content string `json:"content"` // 菜单内容
}

type MsgMenuView struct {
	Url     string `json:"url"`     // 超链接地址
	Content string `json:"content"` // 菜单内容
}

type MsgMenuMiniprogram struct {
	Appid    string `json:"appid"`    // 小程序ID
	Pagepath string `json:"pagepath"` // 小程序路径
	Content  string `json:"content"`  // 菜单内容
}

type MsgMenuItem struct {
	Type        MsgMenuType         `json:"type"` // 菜单类型
	Click       *MsgMenuClick       `json:"click"`
	View        *MsgMenuView        `json:"view"`
	Miniprogram *MsgMenuMiniprogram `json:"miniprogram"`
}

// MsgMenuMessage 菜单消息
type MsgMenuMessage struct {
	HeadContent string        `json:"head_content"` // 起始文本
	List        []MsgMenuItem `json:"list"`         // 菜单项配置
	TailContent string        `json:"tail_content"` // 结束文本
}

// ChannelsShopProductMessage 视频号商品信息
type ChannelsShopProductMessage struct {
	ProductId     string `json:"product_id"`      // 商品ID
	HeadImage     string `json:"head_image"`      // 商品图片
	Title         string `json:"title"`           // 商品标题
	SalesPrice    string `json:"sales_price"`     // 商品价格，以分为单位
	ShopNickname  string `json:"shop_nickname"`   // 店铺名称
	ShopHeadImage string `json:"shop_head_image"` // 店铺头像
}

// ChannelsShopOrderMessage 视频号订单信息
type ChannelsShopOrderMessage struct {
	OrderId       string `json:"order_id"`       // 订单ID
	ProductTitles string `json:"product_titles"` // 商品标题
	PriceWording  string `json:"price_wording"`  // 订单价格描述
	State         string `json:"state"`          // 订单状态
	ImageUrl      string `json:"image_url"`      // 订单缩略图
	ShopNickname  string `json:"shop_nickname"`  // 店铺名称
}

type MergedMsgItem struct {
	SendTime   uint32 `json:"send_time"`   // 发送时间
	MsgType    string `json:"msgtype"`     // 消息类型
	SenderName string `json:"sender_name"` // 发送者名称
	MsgContent string `json:"msg_content"` // 消息内容，Json字符串，结构可参考本文档消息类型说明
}

// MergedMsgMessage 聊天记录消息
type MergedMsgMessage struct {
	Title string          `json:"title"` // 聊天记录标题
	Item  []MergedMsgItem `json:"item"`
}

type EventType string

const (
	EventTypeEnterSession                  EventType = "enter_session"                     // 用户进入会话事件
	EventTypeMsgSendFail                   EventType = "msg_send_fail"                     // 消息发送失败事件
	EventTypeServicerStatusChange          EventType = "servicer_status_change"            // 接待人员接待状态变更事件
	EventTypeSessionStatusChange           EventType = "session_status_change"             // 会话状态变更事件
	EventTypeUserRecallMsg                 EventType = "user_recall_msg"                   // 用户撤回消息事件
	EventTypeServicerRecallMsg             EventType = "servicer_recall_msg"               // 接待人员撤回消息事件
	EventTypeRejectCustomerMsgSwitchChange EventType = "reject_customer_msg_switch_change" // 拒收客户消息变更事件
)

// EventMessage 事件消息

type EventMessage map[string]interface{}

func (e EventMessage) EventType() EventType {
	return EventType(e["event_type"].(string))
}

type EventMessageDecoder[T any] struct {
	msg EventMessage
}

func NewEventMessage[T any](msg EventMessage) *EventMessageDecoder[T] {
	return &EventMessageDecoder[T]{msg: msg}
}

func (d EventMessageDecoder[T]) Decode() (value T, err error) {
	err = mapstructure.Decode(d.msg, &value)

	return
}

// EventEnterSession 用户进入会话事件
type EventEnterSession struct {
	EventType      string `json:"event_type" mapstructure:"event_type"`
	OpenKfid       string `json:"open_kfid" mapstructure:"open_kfid"`
	ExternalUserid string `json:"external_userid" mapstructure:"external_userid"`
	Scene          string `json:"scene" mapstructure:"scene"`
	SceneParam     string `json:"scene_param" mapstructure:"scene_param"`
	WelcomeCode    string `json:"welcome_code" mapstructure:"welcome_code"`
	WechatChannels struct {
		Nickname string `json:"nickname" mapstructure:"nickname"`
		Scene    int    `json:"scene" mapstructure:"scene"`
	} `json:"wechat_channels" mapstructure:"wechat_channels"`
}

// EventSendFail 消息发送失败事件
type EventSendFail struct {
	EventType      string `json:"event_type" mapstructure:"event_type"`
	OpenKfid       string `json:"open_kfid" mapstructure:"open_kfid"`
	ExternalUserid string `json:"external_userid" mapstructure:"external_userid"`
	FailMsgid      string `json:"fail_msgid" mapstructure:"fail_msgid"`
	FailType       int    `json:"fail_type" mapstructure:"fail_type"`
}

// EventServicerStatusChange 接待人员接待状态变更事件
type EventServicerStatusChange struct {
	EventType      string `json:"event_type" mapstructure:"event_type"`
	ServicerUserid string `json:"servicer_userid" mapstructure:"servicer_userid"`
	Status         int    `json:"status" mapstructure:"status"`
	StopType       int    `json:"stop_type" mapstructure:"stop_type"`
	OpenKfid       string `json:"open_kfid" mapstructure:"open_kfid"`
}

// EventSessionStatusChange 会话状态变更事件
type EventSessionStatusChange struct {
	EventType         string `json:"event_type" mapstructure:"event_type"`
	OpenKfid          string `json:"open_kfid" mapstructure:"open_kfid"`
	ExternalUserid    string `json:"external_userid" mapstructure:"external_userid"`
	ChangeType        int    `json:"change_type" mapstructure:"change_type"`
	OldServicerUserid string `json:"old_servicer_userid" mapstructure:"old_servicer_userid"`
	NewServicerUserid string `json:"new_servicer_userid" mapstructure:"new_servicer_userid"`
	MsgCode           string `json:"msg_code" mapstructure:"msg_code"`
}

// EventUserRecallMsg 用户撤回消息事件
type EventUserRecallMsg struct {
	EventType      string `json:"event_type" mapstructure:"event_type"`
	OpenKfid       string `json:"open_kfid" mapstructure:"open_kfid"`
	ExternalUserid string `json:"external_userid" mapstructure:"external_userid"`
	RecallMsgid    string `json:"recall_msgid" mapstructure:"recall_msgid"`
}

// EventServicerRecallMsg 接待人员撤回消息事件
type EventServicerRecallMsg struct {
	EventType      string `json:"event_type" mapstructure:"event_type"`
	OpenKfid       string `json:"open_kfid" mapstructure:"open_kfid"`
	ExternalUserid string `json:"external_userid" mapstructure:"external_userid"`
	RecallMsgid    string `json:"recall_msgid" mapstructure:"recall_msgid"`
	ServicerUserid string `json:"servicer_userid" mapstructure:"servicer_userid"`
}

// EventRejectCustomerMsgSwitchChange 拒收客户消息变更事件
type EventRejectCustomerMsgSwitchChange struct {
	EventType      string `json:"event_type"`
	ServicerUserid string `json:"servicer_userid"`
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
	RejectSwitch   int    `json:"reject_switch"`
}

type EnterSessionContext struct { // 48小时内最后一次进入会话的上下文信息。请求的need_enter_session_context参数设置为1才返回
	Scene          string        `json:"scene"`       // 进入会话的场景值，获取客服账号链接开发者自定义的场景值
	SceneParam     string        `json:"scene_param"` // 进入会话的自定义参数，获取客服账号链接返回的url，开发者按规范拼接的scene_param参数
	WechatChannels WechatChannel `json:"wechat_channels"`
}

type WechatChannel struct {
	Nickname     string `json:"nickname"`      // 视频号名称，视频号场景值为1、2、3时返回此项
	ShopNickname string `json:"shop_nickname"` // 视频号小店名称，视频号场景值为4、5时返回此项
	Scene        int    `json:"scene"`         // 视频号场景值。1：视频号主页，2：视频号直播间商品列表页，3：视频号商品橱窗页，4：视频号小店商品详情页，5：视频号小店订单页
}
