package workwx

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const mediaFieldName = "media"

// Media 欲上传的素材

type Media struct {
	filename string
	filesize int64
	stream   io.Reader
}

// NewMediaFromFile 从操作系统级文件创建一个欲上传的素材对象
func NewMediaFromFile(f *os.File) (*Media, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &Media{
		filename: stat.Name(),
		filesize: stat.Size(),
		stream:   f,
	}, nil
}

// NewMediaFromBuffer 从内存创建一个欲上传的素材对象
func NewMediaFromBuffer(filename string, buf []byte) (*Media, error) {
	stream := bytes.NewReader(buf)
	return &Media{
		filename: filename,
		filesize: int64(len(buf)),
		stream:   stream,
	}, nil
}

func NewMediaFromHttpResp(filename string, response *http.Response) *Media {
	return &Media{
		filename: filename,
		filesize: response.ContentLength,
		stream:   response.Body,
	}
}

func (m *Media) writeTo(w *multipart.Writer) error {
	wr, err := w.CreateFormFile(mediaFieldName, m.filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(wr, m.stream)
	if err != nil {
		return err
	}

	return nil
}

// UploadPermanentImageMedia 上传永久图片素材
func (c *App) UploadPermanentImageMedia(media *Media) (url string, err error) {
	url, err = c.mediaUploadImg(media)
	if err != nil {
		return "", err
	}

	return url, nil
}

const (
	tempMediaTypeImage = "image"
	tempMediaTypeVoice = "voice"
	tempMediaTypeVideo = "video"
	tempMediaTypeFile  = "file"
)

type MediaUploader func(media *Media) (*MediaUploadResult, error)

// UploadTempImageMedia 上传临时图片素材
func (c *App) UploadTempImageMedia(media *Media) (*MediaUploadResult, error) {
	result, err := c.mediaUpload(tempMediaTypeImage, media)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UploadTempVoiceMedia 上传临时语音素材
func (c *App) UploadTempVoiceMedia(media *Media) (*MediaUploadResult, error) {
	result, err := c.mediaUpload(tempMediaTypeVoice, media)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UploadTempVideoMedia 上传临时视频素材
func (c *App) UploadTempVideoMedia(media *Media) (*MediaUploadResult, error) {
	result, err := c.mediaUpload(tempMediaTypeVideo, media)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UploadTempFileMedia 上传临时文件素材
func (c *App) UploadTempFileMedia(media *Media) (*MediaUploadResult, error) {
	result, err := c.mediaUpload(tempMediaTypeFile, media)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FetchMediaInfo 通过header获取媒体资源的信息
func (c *App) FetchMediaInfo(ctx context.Context, mediaId string) (mediaInfo MediaInfoRsp, err error) {
	var (
		req = FetchMediaReq{MediaID: mediaId}
		rsp struct {
			CommonResp
			MediaInfoRsp
		}
	)

	if mediaInfo, err = c.executeWXApiHead(ctx, "/cgi-bin/media/get", req, &rsp, true); err != nil {
		return MediaInfoRsp{}, err
	} else if bizErr := rsp.TryIntoErr(); bizErr != nil {
		return MediaInfoRsp{}, bizErr
	}

	return mediaInfo, nil
}

func (c *App) DownloadMedia(mediaId string) (*http.Response, error) {
	return c.execGet("/cgi-bin/media/get", FetchMediaReq{MediaID: mediaId}, true)
}

func (c *App) RangeDownloadMedia(ctx context.Context, mediaId string, writer io.Writer) error {
	panic("implement me")
}
