package workwx

import "context"

// GetCustomizedAuthUrl 代开发应用获取授权链接
func (c *App) GetCustomizedAuthUrl(ctx context.Context, state string, templateIds []string) (GetCustomizedAuthUrlResp, error) {
	type Response struct {
		CommonResp
		GetCustomizedAuthUrlResp
	}

	var resp Response

	if err := c.executeWXApiJSONPost("/cgi-bin/service/get_customized_auth_url", newIntoBodyer(GetCustomizedAuthUrlReq{
		State:          state,
		TemplateIdList: templateIds,
	}), &resp, true); err != nil {
		return resp.GetCustomizedAuthUrlResp, err
	}

	if err := resp.TryIntoErr(); err != nil {
		return resp.GetCustomizedAuthUrlResp, err
	}

	return resp.GetCustomizedAuthUrlResp, nil
}

func (c *App) GetAppLicenseInfo(ctx context.Context, corpId, suiteId string) (GetAppLicenseInfoResp, error) {
	type Response struct {
		CommonResp
		GetAppLicenseInfoResp
	}

	var (
		resp Response
		req  = GetAppLicenseInfoReq{
			CorpId:  corpId,
			SuiteId: suiteId,
		}
	)

	if err := c.executeWXApiJSONPost("/cgi-bin/license/get_app_license_info", newIntoBodyer(req), &resp, true); err != nil {
		return resp.GetAppLicenseInfoResp, err
	}

	return resp.GetAppLicenseInfoResp, nil
}
