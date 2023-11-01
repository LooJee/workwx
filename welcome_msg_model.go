package workwx

// SendWelcomeMsgReq
// 发送新客户欢迎语请求
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92137#发送新客户欢迎语

type Image struct {
	// MediaID 图片的media_id，可以通过 素材管理 接口获得
	MediaID string `json:"media_id,omitempty"`
	// PicURL 图片的链接，仅可使用 上传图片接口得到的链接
	PicURL string `json:"pic_url,omitempty" comment:"图片url"`
	// Title 图片名
	Title string `json:"title,omitempty"`
}

type Link struct {
	// Desc 图文消息的描述，最长为512字节
	Desc string `json:"desc,omitempty" comment:"网页描述"`
	// PicURL 图文消息封面的url
	PicURL string `json:"picurl,omitempty"`
	// Title 图文消息标题，最长为128字节，必填
	Title string `json:"title" comment:"网页标题"`
	// URL 图文消息的链接，必填
	URL string `json:"url" comment:"网页链接"`
}

type MiniProgram struct {
	// Appid 小程序appid，必须是关联到企业的小程序应用，必填
	Appid string `json:"appid" comment:"小程序 appid"`
	// Page 小程序page路径，必填
	Page string `json:"page" comment:"小程序路径"`
	// PicMediaID 小程序消息封面的mediaid，封面图建议尺寸为520*416，必填
	PicMediaID string `json:"pic_media_id" comment:"小程序封面图"`
	// Title 小程序消息标题，最长为64字节，必填
	Title string `json:"title" comment:"小程序标题"`
}

// Video 视频附件
type Video struct {
	// MediaID 视频的media_id，可以通过 素材管理 接口获得，必填
	MediaId string `json:"media_id,omitempty"`
	Url     string `json:"url" comment:"视频链接"`
	Title   string `json:"title"`
}

// File 文件附件
type File struct {
	MediaId string `json:"media_id,omitempty"`
	Url     string `json:"url" comment:"文件链接"`
	Title   string `json:"title"`
}

type Attachments struct {
	MsgType     string       `json:"msgtype" binding:"oneof=image file video miniprogram link"` // 附件类型，可选image、link、miniprogram或者video，必填
	Image       *Image       `json:"image,omitempty" binding:"required_if=MsgType image" comment:"图片"`
	Link        *Link        `json:"link,omitempty" binding:"required_if=MsgType link" comment:"网页"`
	MiniProgram *MiniProgram `json:"miniprogram,omitempty" binding:"required_if=MsgTpe miniprogram" comment:"小程序"`
	Video       *Video       `json:"video,omitempty" binding:"required_if=MsgType video" comment:"视频"`
	File        *File        `json:"file,omitempty" binding:"required_if=MsgType file" comment:"文件"`
}

type Text struct {
	// Content 消息文本内容,最长为4000字节
	Content string `json:"content,omitempty"`
}

const (
	AttachmentImage       = "image"
	AttachmentLink        = "link"
	AttachmentMiniProgram = "miniprogram"
	AttachmentVideo       = "video"
	AttachmentFile        = "file"
)
