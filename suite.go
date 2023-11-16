package workwx

import (
	"context"
)

func (c *App) GetPermanentCode(ctx context.Context, authCode string) (GetPermanentCodeResp, error) {
	type Response struct {
		CommonResp
		GetPermanentCodeResp
	}
	var resp Response
	if err := c.executeWXApiJSONPost("/cgi-bin/service/get_permanent_code", newIntoBodyer(GetPermanentCodeReq{AuthCode: authCode}), &resp, true); err != nil {
		return GetPermanentCodeResp{}, err
	}

	if err := resp.TryIntoErr(); err != nil {
		return GetPermanentCodeResp{}, err
	}

	return resp.GetPermanentCodeResp, nil
}
