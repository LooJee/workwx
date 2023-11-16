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
