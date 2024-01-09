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

func (c *App) GetAuthInfo(ctx context.Context, corpId, permanentCode string) (GetAuthInfoResp, error) {
	type Response struct {
		CommonResp
		GetAuthInfoResp
	}
	var resp Response
	if err := c.executeWXApiJSONPost("/cgi-bin/service/get_auth_info", newIntoBodyer(GetAuthInfoReq{AuthCorpId: corpId, PermanentCode: permanentCode}), &resp, true); err != nil {
		return GetAuthInfoResp{}, err
	}

	if err := resp.TryIntoErr(); err != nil {
		return GetAuthInfoResp{}, err
	}

	return resp.GetAuthInfoResp, nil
}

// 获取应用二维码，仅限第三方应用使用
func (c *App) GetAppQrcode(ctx context.Context, req GetAppQrcodeReq) (GetAppQrcodeResp, error) {
	type Response struct {
		CommonResp
		GetAppQrcodeResp
	}
	var resp Response
	if err := c.executeWXApiJSONPost("/cgi-bin/service/get_app_qrcode", newIntoBodyer(req), &resp, true); err != nil {
		return GetAppQrcodeResp{}, err
	}

	if err := resp.TryIntoErr(); err != nil {
		return GetAppQrcodeResp{}, err
	}

	return resp.GetAppQrcodeResp, nil
}
