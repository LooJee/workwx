package workwx

import "net/url"

type GetPermanentCodeReq struct {
	AuthCode string `json:"auth_code"`
}

type DealerCorpInfo struct {
	CorpId   string `json:"corpid"`
	CorpName string `json:"corp_name"`
}

type AuthCorpInfo struct {
	CorpId            string `json:"corpid"`
	CorpName          string `json:"corp_name"`
	CorpType          string `json:"corp_type"`
	CorpSquareLogoUrl string `json:"corp_square_logo_url"`
	CorpUserMax       int    `json:"corp_user_max"`
	CorpAgentMax      int    `json:"corp_agent_max"`
	CorpFullName      string `json:"corp_full_name"`
	VerifiedEndTime   int    `json:"verified_end_time"`
	SubjectType       int    `json:"subject_type"`
	CorpWxqrcode      string `json:"corp_wxqrcode"`
	CorpScale         string `json:"corp_scale"`
	CorpIndustry      string `json:"corp_industry"`
	CorpSubIndustry   string `json:"corp_sub_industry"`
	Location          string `json:"location"`
}

type AuthInfoAgentPrivilege struct {
	Level      int      `json:"level"`
	AllowParty []int    `json:"allow_party"`
	AllowUser  []string `json:"allow_user"`
	AllowTag   []int    `json:"allow_tag"`
	ExtraParty []int    `json:"extra_party"`
	ExtraUser  []string `json:"extra_user"`
	ExtraTag   []int    `json:"extra_tag"`
}

type AuthInfoAgentSharedFrom struct {
	CorpId    string `json:"corpid"`
	ShareType int    `json:"share_type"`
}

type AuthInfoAgent struct {
	AgentId          int                     `json:"agentid"`
	Name             string                  `json:"name"`
	RoundLogoUrl     string                  `json:"round_logo_url"`
	SquareLogoUrl    string                  `json:"square_logo_url"`
	Appid            int                     `json:"appid"`
	AuthMode         int                     `json:"auth_mode,omitempty"`
	IsCustomizedApp  bool                    `json:"is_customized_app,omitempty"`
	AuthFromThirdapp bool                    `json:"auth_from_thirdapp,omitempty"`
	Privilege        AuthInfoAgentPrivilege  `json:"privilege,omitempty"`
	SharedFrom       AuthInfoAgentSharedFrom `json:"shared_from"`
}

type AuthUserInfo struct {
	UserId     string `json:"userid"`
	OpenUserId string `json:"open_userid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
}

type RegisterCodeInfo struct {
	RegisterCode string `json:"register_code"`
	TemplateId   string `json:"template_id"`
	State        string `json:"state"`
}

type GetPermanentCodeResp struct {
	AccessToken   string `json:"access_token"`
	ExpiresIn     int    `json:"expires_in"`
	PermanentCode string `json:"permanent_code"`
	State         string `json:"state"`
	CorpAuthInfo
}

type CorpAuthInfo struct {
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       struct {
		Agent []AuthInfoAgent `json:"agent"`
	} `json:"auth_info"`
	AuthUserInfo     AuthUserInfo     `json:"auth_user_info"`
	RegisterCodeInfo RegisterCodeInfo `json:"register_code_info"`
}

type GetAuthInfoResp struct {
	CorpAuthInfo
}

type GetAuthInfoReq struct {
	AuthCorpId    string `json:"auth_corpid"`
	PermanentCode string `json:"permanent_code"`
}

type GetAppQrcodeReq struct {
	SuiteId    string `json:"suite_id"`
	State      string `json:"state,omitempty"`
	Style      int    `json:"style"`
	ResultType int    `json:"result_type"`
}

type GetAppQrcodeResp struct {
	Qrcode string `json:"qrcode"` // 二维码URL地址
}

type GetCorpTokenReq struct {
	AuthCorpId    string `json:"auth_corpid"`
	PermanentCode string `json:"permanent_code"`
}

type GetCorpTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type SearchContactReq struct {
	AuthCorpid     string `json:"auth_corpid"`
	QueryWord      string `json:"query_word"`
	QueryType      int    `json:"query_type"`
	QueryRange     int    `json:"query_range"`
	Agentid        int    `json:"agentid"`
	Limit          int    `json:"limit"`
	FullMatchField int    `json:"full_match_field"`
	Cursor         string `json:"cursor"`
}

type SearchContactResp struct {
	QueryResult SearchContactQueryResult `json:"query_result"`
	IsLast      bool                     `json:"is_last"`
	NextCursor  string                   `json:"next_cursor"`
}

type SearchContactQueryResult struct {
	User  SearchContactQueryResultUser `json:"user"`
	Party struct {
		DepartmentId []int `json:"department_id"`
	} `json:"party"`
	DismissUser SearchContactQueryResultUser `json:"dismiss_user"`
}

type SearchContactQueryResultUser struct {
	Userid     []string `json:"userid"`
	OpenUserid []string `json:"open_userid"`
}

type GetUserInfoThirdReq struct {
	Code string `json:"code"`
}

var _ urlValuer = GetUserInfoThirdReq{}

func (x GetUserInfoThirdReq) intoURLValues() url.Values {
	return url.Values{
		"code": {x.Code},
	}
}

type GetUserInfoThirdResp struct {
	Corpid     string `json:"corpid"`
	Userid     string `json:"userid"`
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int    `json:"expires_in"`
	OpenUserid string `json:"open_userid"`
}
