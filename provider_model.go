package workwx

type GetCustomizedAuthUrlReq struct {
	State          string   `json:"state"`
	TemplateIdList []string `json:"templateid_list"`
}

type GetCustomizedAuthUrlResp struct {
	QrcodeUrl string `json:"qrcode_url"`
	ExpiresIn int    `json:"expires_in"`
}

type AppQrcodeStyle int

const (
	AppQrcodeStyle0 AppQrcodeStyle = iota //带说明外框的二维码，适合于实体物料
	AppQrcodeStyle1                       // 带说明外框的二维码，适合于屏幕类
	AppQrcodeStyle2                       // 不带说明外框（小尺寸）
	AppQrcodeStyle3                       // 不带说明外框（中尺寸）
	AppQrcodeStyle4                       // 不带说明外框（大尺寸）
)

type AppQrcodeResultType int

const (
	AppQrcodeResultTypeBuffer AppQrcodeResultType = iota + 1 // 返回二维码 buffer
	AppQrcodeResultTypeUrl                                   // 返回二维码 url
)

type GetAppQrCodeReq struct {
	SuiteId    string              `json:"suite_id"`    // 第三方应用id
	State      string              `json:"state"`       // 场景值，可为空
	Style      AppQrcodeStyle      `json:"style"`       // 二维码类型，可为空，默认为不带说明外框小尺寸
	ResultType AppQrcodeResultType `json:"result_type"` // 二维码返回类型，可为空，默认为 buffer
}

type GetAppQrCodeResp struct {
	Qrcode string `json:"qrcode"`
}
