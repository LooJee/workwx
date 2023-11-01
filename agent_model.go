package workwx

type AgentInfo struct {
	AgentId        int      `json:"agentid"`         // 企业应用id
	Name           string   `json:"name"`            // 企业应用名称
	SquareLogoUrl  string   `json:"square_logo_url"` // 企业应用方形头像
	Description    string   `json:"description"`     // 企业应用详情
	AllowUserInfos struct { //企业应用可见范围（人员），不为空时表示单独针对人员设置可见范围
		User []struct {
			Userid string `json:"userid"`
		} `json:"user"`
	} `json:"allow_userinfos"`
	AllowParties struct { // 企业应用可见范围（部门）
		PartyId []int `json:"partyid"`
	} `json:"allow_partys"`
	AllowTags struct { // 企业应用可见范围（标签）
		TagId []int `json:"tagid"`
	} `json:"allow_tags"`
	Close                   int    `json:"close"`                     // 企业应用是否被停用。0：未被停用；1：被停用
	RedirectDomain          string `json:"redirect_domain"`           // 企业应用可信域名
	ReportLocationFlag      int    `json:"report_location_flag"`      // 企业应用是否打开地理位置上报 0：不上报；1：进入会话上报；
	IsReportEnter           int    `json:"isreportenter"`             // 是否上报用户进入应用事件。0：不接收；1：接收
	HomeUrl                 string `json:"home_url"`                  // 应用主页url
	CustomizedPublishStatus int    `json:"customized_publish_status"` // 代开发自建应用返回该字段，表示代开发发布状态。0：待开发（企业已授权，服务商未创建应用）；1：开发中（服务商已创建应用，未上线）；2：已上线（服务商已上线应用且不存在未上线版本）；3：存在未上线版本（服务商已上线应用但存在未上线版本）
}
