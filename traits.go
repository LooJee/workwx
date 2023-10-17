package workwx

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// urlValuer 可转化为 url.Values 类型的 trait
type urlValuer interface {
	intoURLValues() url.Values
}

// Bodyer 可转化为 API 请求体的 trait
type bodyer interface {
	intoBody() ([]byte, error)
}

// mediaUploader 携带 *Media 对象，可转化为 multipart 文件上传请求体的 trait
type mediaUploader interface {
	getMedia() *Media
}

func intoJsonBody(x interface{}) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}
