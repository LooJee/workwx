package workwx

import "context"

// AddCustomerServiceAccount 添加客服账号
// https://developer.work.weixin.qq.com/document/path/94662
func (c *App) AddCustomerServiceAccount(ctx context.Context, req AddCustomerServiceAccountReq) (rsp AddCustomerServiceAccountRsp, err error) {
	var out struct {
		CommonResp
		AddCustomerServiceAccountRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/account/add", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.AddCustomerServiceAccountRsp, nil
}

// DelCustomerServiceAccount 删除客服账号
// https://developer.work.weixin.qq.com/document/path/94663
func (c *App) DelCustomerServiceAccount(ctx context.Context, req DelCustomerServiceAccountReq) (err error) {
	var out CommonResp

	err = c.executeWXApiJSONPost("/cgi-bin/kf/account/del", newIntoBodyer(req), &out, true)
	if err != nil {
		return err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return bizErr
	}

	return nil
}

// UpdateCustomerServiceAccount 修改客服账号
// https://developer.work.weixin.qq.com/document/path/94663
func (c *App) UpdateCustomerServiceAccount(ctx context.Context, req UpdateCustomerServiceAccountReq) (err error) {
	var out CommonResp

	err = c.executeWXApiJSONPost("/cgi-bin/kf/account/update", newIntoBodyer(req), &out, true)
	if err != nil {
		return err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return bizErr
	}

	return nil
}

// FetchCustomerServiceAccounts 获取客服账号列表
// https://developer.work.weixin.qq.com/document/path/94661
func (c *App) FetchCustomerServiceAccounts(ctx context.Context, req FetchCustomerServiceAccountsReq) (rsp FetchCustomerServiceAccountRsp, err error) {
	var out struct {
		CommonResp
		FetchCustomerServiceAccountRsp
	}
	err = c.executeWXApiJSONPost("/cgi-bin/kf/account/list", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	}
	if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.FetchCustomerServiceAccountRsp, nil
}

// FetchAllCustomerServiceAccounts 获取所有的客服账号列表
func (c *App) FetchAllCustomerServiceAccounts(ctx context.Context) (data []CustomerServiceAccountItem, err error) {
	var (
		offset uint32 = 0
		limit  uint32 = 100
	)

	for {
		rsp, err := c.FetchCustomerServiceAccounts(ctx, FetchCustomerServiceAccountsReq{
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return data, err
		}
		data = append(data, rsp.AccountList...)
		if len(rsp.AccountList) < int(limit) {
			break
		}
		offset++
	}

	return
}

// FetchCustomerServiceAccountContactWay 获取客服账号链接
// https://developer.work.weixin.qq.com/document/path/94665
func (c *App) FetchCustomerServiceAccountContactWay(ctx context.Context, req FetchCustomerServiceAccountContactWayReq) (rsp FetchCustomerServiceAccountContactWayRsp, err error) {
	var out struct {
		CommonResp
		FetchCustomerServiceAccountContactWayRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/add_contact_way", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.FetchCustomerServiceAccountContactWayRsp, nil
}

// AddCustomerServiceServicer 添加接待人员
// https://developer.work.weixin.qq.com/document/path/94646
func (c *App) AddCustomerServiceServicer(ctx context.Context, req AddCustomerServiceServicerReq) (rsp AddCustomerServiceServicerRsp, err error) {
	var out struct {
		CommonResp
		AddCustomerServiceServicerRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/servicer/add", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.AddCustomerServiceServicerRsp, nil
}

// DelCustomerServiceServicer 删除接待人员
// https://developer.work.weixin.qq.com/document/path/94647
func (c *App) DelCustomerServiceServicer(ctx context.Context, req DelCustomerServiceServicerReq) (rsp DelCustomerServiceServicerRsp, err error) {
	var out struct {
		CommonResp
		DelCustomerServiceServicerRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/servicer/del", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.DelCustomerServiceServicerRsp, nil
}

// FetchCustomerServiceServicers 获取接待人员列表
// https://developer.work.weixin.qq.com/document/path/94645
func (c *App) FetchCustomerServiceServicers(ctx context.Context, req FetchCustomerServiceServicersReq) (rsp FetchCustomerServiceServicersRsp, err error) {
	var out struct {
		CommonResp
		FetchCustomerServiceServicersRsp
	}

	err = c.executeWXApiGet("/cgi-bin/kf/servicer/list", req, &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.FetchCustomerServiceServicersRsp, nil
}

// FetchCustomerServiceState 获取会话状态
// https://developer.work.weixin.qq.com/document/path/94669
func (c *App) FetchCustomerServiceState(ctx context.Context, req FetchCustomerServiceStateReq) (rsp FetchCustomerServiceStateRsp, err error) {
	var out struct {
		CommonResp
		FetchCustomerServiceStateRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/service_state/get", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.FetchCustomerServiceStateRsp, nil
}

// TransCustomerServiceState 变更会话状态
// https://developer.work.weixin.qq.com/document/path/94669#%E5%8F%98%E6%9B%B4%E4%BC%9A%E8%AF%9D%E7%8A%B6%E6%80%81
func (c *App) TransCustomerServiceState(ctx context.Context, req TransCustomerServiceStateReq) (rsp TransCustomerServiceStateRsp, err error) {
	var out struct {
		CommonResp
		TransCustomerServiceStateRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/service_state/trans", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.TransCustomerServiceStateRsp, nil
}

// SyncCustomerServiceMsg 接收消息和事件
// https://developer.work.weixin.qq.com/document/path/94670
func (c *App) SyncCustomerServiceMsg(ctx context.Context, req SyncCustomerServiceMsgReq) (rsp SyncCustomerServiceMsgRsp, err error) {
	var out struct {
		CommonResp
		SyncCustomerServiceMsgRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/sync_msg", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.SyncCustomerServiceMsgRsp, nil
}

// SendCustomerServiceMsg 发送客服消息
// https://developer.work.weixin.qq.com/document/path/94677
func (c *App) SendCustomerServiceMsg(ctx context.Context, req SendCustomerServiceMsgReq) (rsp SendCustomerServiceMsgRsp, err error) {
	var out struct {
		CommonResp
		SendCustomerServiceMsgRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/send_msg", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.SendCustomerServiceMsgRsp, nil
}

// SendCustomerServiceMsgOnEvent 发送事件响应消息
// https://developer.work.weixin.qq.com/document/path/95122
func (c *App) SendCustomerServiceMsgOnEvent(ctx context.Context, req SendMsgOnCustomerServiceEventReq) (rsp SendMsgOnCustomerServiceEventRsp, err error) {
	var out struct {
		CommonResp
		SendMsgOnCustomerServiceEventRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/send_msg_on_event", newIntoBodyer(req), &out, true)
	if err != nil {
		return
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.SendMsgOnCustomerServiceEventRsp, nil
}

// FetchCustomerServiceCustomers 获取客户基础信息
// https://developer.work.weixin.qq.com/document/path/95159
func (c *App) FetchCustomerServiceCustomers(ctx context.Context, req FetchCustomerServiceCustomersReq) (rsp FetchCustomerServiceCustomersRsp, err error) {
	var out struct {
		CommonResp
		FetchCustomerServiceCustomersRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/customer/batchget", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.FetchCustomerServiceCustomersRsp, nil
}

// GetCustomerServiceCorpStatistic 获取「客户数据统计」企业汇总数据
// https://developer.work.weixin.qq.com/document/path/95489
func (c *App) GetCustomerServiceCorpStatistic(ctx context.Context, req GetCustomerServiceCorpStatisticReq) (rsp GetCustomerServiceCorpStatisticRsp, err error) {
	var out struct {
		CommonResp
		GetCustomerServiceCorpStatisticRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/get_corp_statistic", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.GetCustomerServiceCorpStatisticRsp, nil
}

// GetCustomerServiceServicerStatistic 获取「客户数据统计」接待人员明细数据
// https://developer.work.weixin.qq.com/document/path/95490
func (c *App) GetCustomerServiceServicerStatistic(ctx context.Context, req GetCustomerServiceServicerStatisticReq) (rsp GetCustomerServiceServicerStatisticRsp, err error) {
	var out struct {
		CommonResp
		GetCustomerServiceServicerStatisticRsp
	}

	err = c.executeWXApiJSONPost("/cgi-bin/kf/get_servicer_statistic", newIntoBodyer(req), &out, true)
	if err != nil {
		return rsp, err
	} else if bizErr := out.TryIntoErr(); bizErr != nil {
		return rsp, bizErr
	}

	return out.GetCustomerServiceServicerStatisticRsp, nil
}
