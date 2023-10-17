package workwx

import (
	"net/url"

	"github.com/loojee/workwx/customer_service"
)

// 添加客服账号 Models

type AddCustomerServiceAccountReq struct {
	Name    string `json:"name"`     // 客服名称
	MediaId string `json:"media_id"` // 客服头像临时素材
}

type AddCustomerServiceAccountRsp struct {
	OpenKfid string `json:"open_kfid"` // 新创建的客服账号ID
}

// 删除客服账号 Models

type DelCustomerServiceAccountReq struct {
	OpenKfid string `json:"open_kfid"`
}

// 修改客服账号 Models

type UpdateCustomerServiceAccountReq struct {
	OpenKfid string `json:"open_kfid"` // 要修改的客服账号ID。
	Name     string `json:"name"`      // 新的客服名称，如不需要修改可不填。不多于16个字符
	MediaId  string `json:"media_id"`  // 新的客服头像临时素材，如不需要修改可不填。不多于128个字节
}

// 获取客服账号列表 Models

type FetchCustomerServiceAccountsReq struct {
	Offset uint32 `json:"offset"` // 分页，偏移量, 默认为0
	Limit  uint32 `json:"limit"`  // 分页，预期请求的数据量，默认为100，取值范围 1 ~ 100
}

type CustomerServiceAccountItem struct {
	OpenKfid        string `json:"open_kfid"`        // 客服账号ID
	Name            string `json:"name"`             // 客服名称
	Avatar          string `json:"avatar"`           // 客服头像URL
	ManagePrivilege bool   `json:"manage_privilege"` // 当前调用接口的应用身份，是否有该客服账号的管理权限（编辑客服账号信息、分配会话和收发消息）。组件应用不返回此字段
}

type FetchCustomerServiceAccountRsp struct {
	AccountList []CustomerServiceAccountItem `json:"account_list"`
}

// 获取客服账号链接 Models

type FetchCustomerServiceAccountContactWayReq struct {
	OpenKfid string `json:"open_kfid"`
	Scene    string `json:"scene"`
}

type FetchCustomerServiceAccountContactWayRsp struct {
	Url string `json:"url"`
}

// 添加接待人员 Models

type AddCustomerServiceServicerReq struct {
	OpenKfid         string   `json:"open_kfid"`          // 客服账号ID
	UseridList       []string `json:"userid_list"`        // 接待人员userid列表。可填充个数：0 ~ 100。超过100个需分批调用。
	DepartmentIdList []int    `json:"department_id_list"` // 接待人员部门id列表。可填充个数：0 ~ 100。超过100个需分批调用。
	// userid_list和department_id_list至少需要填其中一个
}

type AddCustomerServiceServicerItem struct {
	Userid       string `json:"userid,omitempty"` // 接待人员的userid，当userid为空时，表示返回的是部门
	DepartmentId int    `json:"department_id"`    // 接待人员部门的id
	CommonResp
}

type AddCustomerServiceServicerRsp struct {
	ResultList []AddCustomerServiceServicerItem `json:"result_list"`
}

// 删除接待人员 Models

type DelCustomerServiceServicerReq struct {
	OpenKfid         string   `json:"open_kfid"`          // 客服账号ID
	UseridList       []string `json:"userid_list"`        // 接待人员userid列表。可填充个数：0 ~ 100。超过100个需分批调用。
	DepartmentIdList []int    `json:"department_id_list"` // 接待人员部门id列表。可填充个数：0 ~ 100。超过100个需分批调用。
}

type DelCustomerServiceServicerItem AddCustomerServiceServicerItem

type DelCustomerServiceServicerRsp struct {
	ResultList []DelCustomerServiceServicerItem `json:"result_list"`
}

// 获取接待人员列表 Models

type FetchCustomerServiceServicersReq struct {
	OpenKfid string // 客服账号ID
}

func (x FetchCustomerServiceServicersReq) intoURLValues() url.Values {
	return url.Values{
		"open_kfid": {x.OpenKfid},
	}
}

type FetchCustomerServiceServicersItem struct {
	Userid       string `json:"userid,omitempty"` // 接待人员的userid,当userid为空时，表示返回的是部门
	DepartmentId int    `json:"department_id"`    // 接待人员部门的id
	Status       int    `json:"status"`           // 接待人员的接待状态。0:接待中,1:停止接待。
	StopType     uint   `json:"stop_type"`        // 接待人员的接待状态为「停止接待」的子类型。0:停止接待,1:暂时挂起
}

type FetchCustomerServiceServicersRsp struct {
	ServicerList []FetchCustomerServiceServicersItem `json:"servicer_list"`
}

// 获取会话状态 Models

type FetchCustomerServiceStateReq struct {
	OpenKfid       string `json:"open_kfid"`       // 客服账号ID
	ExternalUserid string `json:"external_userid"` // 微信客户的external_userid
}

type FetchCustomerServiceStateRsp struct {
	ServiceState   int    `json:"service_state"`   // 当前的会话状态
	ServicerUserid string `json:"servicer_userid"` // 接待人员的userid。
}

// 变更会话状态 Models

type TransCustomerServiceStateReq struct {
	OpenKfid       string `json:"open_kfid"`       // 客服账号ID
	ExternalUserid string `json:"external_userid"` // 微信客户的external_userid
	ServiceState   int    `json:"service_state"`   // 变更的目标状态，状态定义和所允许的变更可参考概述中的流程图和表格
	ServicerUserid string `json:"servicer_userid"` // 接待人员的userid。
}

type TransCustomerServiceStateRsp struct {
	MsgCode string `json:"msg_code"` //用于发送响应事件消息的code，将会话初次变更为service_state为2和3时，返回回复语code，service_state为4时，返回结束语code。可用该code调用发送事件响应消息接口给客户发送事件响应消息
}

// 接收消息 Models

type SyncCustomerServiceMsgReq struct {
	Cursor      string `json:"cursor"`
	Token       string `json:"token"`
	Limit       int    `json:"limit"`
	VoiceFormat int    `json:"voice_format"`
	OpenKfid    string `json:"open_kfid"`
}

type SyncCustomerServiceMsgItem struct {
	Msgid          string `json:"msgid"`
	SendTime       int    `json:"send_time"`
	Origin         int    `json:"origin"`
	OpenKfid       string `json:"open_kfid,omitempty"`
	ExternalUserid string `json:"external_userid,omitempty"`
	ServicerUserid string `json:"servicer_userid,omitempty"`

	customer_service.MsgItem
}

type SyncCustomerServiceMsgRsp struct {
	NextCursor string                       `json:"next_cursor"`
	MsgList    []SyncCustomerServiceMsgItem `json:"msg_list"`
	HasMore    int                          `json:"has_more"`
}

// 发送消息 Models

type SendCustomerServiceMsgReq struct {
	Touser   string `json:"touser"`
	OpenKfid string `json:"open_kfid"`
	Msgid    string `json:"msgid,omitempty"` // 可不填

	Msgtype     customer_service.MessageType         `json:"msgtype"`
	Text        *customer_service.TextMessage        `json:"text,omitempty"`
	Image       *customer_service.ImageMessage       `json:"image,omitempty"`
	Voice       *customer_service.VoiceMessage       `json:"voice,omitempty"`
	Video       *customer_service.VideoMessage       `json:"video,omitempty"`
	File        *customer_service.FileMessage        `json:"file,omitempty"`
	Link        *customer_service.LinkMessage        `json:"link,omitempty"` // 图文链接
	MiniProgram *customer_service.MiniProgramMessage `json:"mini_program"`   // 小程序消息
	MsgMenu     *customer_service.MsgMenuMessage     `json:"msg_menu"`       // 菜单消息
	Location    *customer_service.LocationMessage    `json:"location"`       // 地理位置消息
}

type SendCustomerServiceMsgRsp struct {
	MsgId string `json:"msgid"` // 消息ID
}

// 发送欢迎语等事件响应消息 Models

type SendMsgOnCustomerServiceEventReq struct {
	Code  string `json:"code"`  // 事件响应消息对应的code。通过事件回调下发，仅可使用一次。
	Msgid string `json:"msgid"` // 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。 不多于32字节

	Msgtype customer_service.MessageType     `json:"msgtype"`            // 消息类型。对不同的msgtype，有相应的结构描述，详见消息类型
	Text    *customer_service.TextMessage    `json:"text,omitempty"`     // 文本消息
	MsgMenu *customer_service.MsgMenuMessage `json:"msg_menu,omitempty"` // 菜单消息
}

type SendMsgOnCustomerServiceEventRsp struct {
	Msgid string `json:"msgid"` // 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。 不多于32字节
}

// FetchCustomerServiceCustomersReq 获取客户基础信息
// https://developer.work.weixin.qq.com/document/path/95159

type FetchCustomerServiceCustomersReq struct {
	ExternalUseridList      []string `json:"external_userid_list"`       // external_userid列表，可填充个数：1 ~ 100。超过100个需分批调用。
	NeedEnterSessionContext int      `json:"need_enter_session_context"` // 是否需要返回客户48小时内最后一次进入会话的上下文信息。0-不返回 1-返回。默认不返回
}

type FetchCustomerServiceCustomersItem struct {
	ExternalUserid      string   `json:"external_userid"` // 微信客户的external_userid
	Nickname            string   `json:"nickname"`        // 微信昵称
	Avatar              string   `json:"avatar"`          // 微信头像。
	Gender              int      `json:"gender"`          // 性别。第三方不可获取，统一返回0
	Unionid             string   `json:"unionid"`         // unionid，需要绑定微信开发者账号才能获取到，查看绑定方法。第三方不可获取
	EnterSessionContext struct { // 48小时内最后一次进入会话的上下文信息。请求的need_enter_session_context参数设置为1才返回
		Scene          string   `json:"scene"`       // 进入会话的场景值，获取客服账号链接开发者自定义的场景值
		SceneParam     string   `json:"scene_param"` // 进入会话的自定义参数，获取客服账号链接返回的url，开发者按规范拼接的scene_param参数
		WechatChannels struct { // 进入会话的视频号信息，从视频号进入会话才有值
			Nickname     string `json:"nickname"`      // 视频号名称，视频号场景值为1、2、3时返回此项
			ShopNickname string `json:"shop_nickname"` // 视频号小店名称，视频号场景值为4、5时返回此项
			Scene        int    `json:"scene"`         // 视频号场景值。1：视频号主页，2：视频号直播间商品列表页，3：视频号商品橱窗页，4：视频号小店商品详情页，5：视频号小店订单页
		} `json:"wechat_channels"`
	} `json:"enter_session_context"`
}

type FetchCustomerServiceCustomersRsp struct {
	CustomerList          []FetchCustomerServiceCustomersItem `json:"customer_list"`
	InvalidExternalUserId []string                            `json:"invalid_external_userid"`
}

// 获取「客户数据统计」企业汇总数据 Models

type GetCustomerServiceCorpStatisticReq struct {
	OpenKfid  string `json:"open_kfid"`  // 客服账号ID。不传入时返回的数据为企业维度汇总的数据
	StartTime int    `json:"start_time"` // 起始日期的时间戳，填这一天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天
	EndTime   int    `json:"end_time"`   // 结束日期的时间戳，填这一天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天
}

type CustomerServiceCorpStatisticItem struct {
	SessionCnt                int `json:"session_cnt"`                  // 咨询会话数。客户发过消息并分配给接待人员或智能助手的客服会话数，转接不会产生新的会话
	CustomerCnt               int `json:"customer_cnt"`                 // 咨询客户数。在会话中发送过消息的客户数量，若客户多次咨询只计算一个客户
	CustomerMsgCnt            int `json:"customer_msg_cnt"`             // 咨询消息总数。客户在会话中发送的消息的数量
	UpgradeServiceCustomerCnt int `json:"upgrade_service_customer_cnt"` // 升级服务客户数。通过「升级服务」功能成功添加专员或加入客户群的客户数，若同一个客户添加多个专员或客户群，只计算一个客户。在2022年3月10日以后才会有对应统计数据
	AiSessionReplyCnt         int `json:"ai_session_reply_cnt"`         // 智能回复会话数。客户发过消息并分配给智能助手的咨询会话数。通过API发消息或者开启智能回复功能会将客户分配给智能助手
	AiTransferRate            int `json:"ai_transfer_rate"`             // 转人工率。一个自然日内，客户给智能助手发消息的会话中，转人工的会话的占比。
	AiKnowledgeHitRate        int `json:"ai_knowledge_hit_rate"`        // 知识命中率。一个自然日内，客户给智能助手发送的消息中，命中知识库的占比。只有在开启了智能回复原生功能并配置了知识库的情况下，才会产生该项统计数据。当api托管了会话分配，智能回复原生功能失效。若不返回，代表没有向配置知识库的智能接待助手发送消息，该项无法计算
	MsgRejectedCustomerCnt    int `json:"msg_rejected_customer_cnt"`    // 被拒收消息的客户数。被接待人员设置了“不再接收消息”的客户数
}

type GetCustomerServiceCorpStatisticRsp struct {
	StatisticList []struct {
		StartTime uint32                           `json:"start_time"`
		Statistic CustomerServiceCorpStatisticItem `json:"statistic"`
	} `json:"statistic_list"`
}

// 获取「客户数据统计」接待人员明细数据 Models

type GetCustomerServiceServicerStatisticReq struct {
	OpenKfid  string `json:"open_kfid"`  // 客服账号ID。不传入时返回的数据为接待人员维度汇总的数据
	Servicer  string `json:"servicer"`   // 接待人员的userid。
	StartTime int    `json:"start_time"` // 起始日期的时间戳，填这一天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天
	EndTime   int    `json:"end_time"`   // 结束日期的时间戳，填这一天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天
}

type CustomerServiceServicerStatisticItem struct {
	SessionCnt                         int `json:"session_cnt"`
	CustomerCnt                        int `json:"customer_cnt"`
	CustomerMsgCnt                     int `json:"customer_msg_cnt"`
	ReplyRate                          int `json:"reply_rate"`
	FirstReplyAverageSec               int `json:"first_reply_average_sec"`
	SatisfactionInvestgateCnt          int `json:"satisfaction_investgate_cnt"`
	SatisfactionParticipationRate      int `json:"satisfaction_participation_rate"`
	SatisfiedRate                      int `json:"satisfied_rate"`
	MiddlingRate                       int `json:"middling_rate"`
	DissatisfiedRate                   int `json:"dissatisfied_rate"`
	UpgradeServiceCustomerCnt          int `json:"upgrade_service_customer_cnt"`
	UpgradeServiceMemberInviteCnt      int `json:"upgrade_service_member_invite_cnt"`
	UpgradeServiceMemberCustomerCnt    int `json:"upgrade_service_member_customer_cnt"`
	UpgradeServiceGroupchatInviteCnt   int `json:"upgrade_service_groupchat_invite_cnt"`
	UpgradeServiceGroupchatCustomerCnt int `json:"upgrade_service_groupchat_customer_cnt"`
	MsgRejectedCustomerCnt             int `json:"msg_rejected_customer_cnt"`
}

type GetCustomerServiceServicerStatisticRsp struct {
	StatisticList []struct {
		StartTime uint32                           `json:"start_time"`
		Statistic CustomerServiceCorpStatisticItem `json:"statistic"`
	} `json:"statistic_list"`
}
