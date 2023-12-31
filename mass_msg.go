package workwx

// AddMsgTemplate 创建企业群发
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92135#创建企业群发
func (c *App) AddMsgTemplate(req AddMsgTemplateReq) (msgID string, failedList []string, err error) {
	var resp addMsgTemplateResp
	resp, err = c.execAddMsgTemplate(req)
	if err != nil {
		return "", nil, err
	}
	return resp.MsgID, resp.FailList, nil
}

// GetGroupMsgSendResultExternalContact 获取企业群发成员执行结果
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取企业群发成员执行结果
func (c *App) GetGroupMsgSendResultExternalContact(req GetGroupMsgSendResultExternalContactReq) (res GetGroupMsgSendResultExternalContactResp, err error) {
	var resp GetGroupMsgSendResultExternalContactResp
	resp, err = c.execGetGroupMsgSendResultExternalContact(req)
	if err != nil {
		return
	}
	// ok = resp.IsOK()
	return resp, nil
}

// GetGroupMsgTaskExternalContact 获取群发成员发送任务列表
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发成员发送任务列表
func (c *App) GetGroupMsgTaskExternalContact(req ReqGetGroupmsgTaskExternalcontact) (resp GetGroupMsgTaskExternalContactResp, err error) {
	resp, err = c.execGetGroupmsgTaskExternalcontact(req)
	if err != nil {
		return resp, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return resp, bizErr
	}
	return
}

// GetGroupmsgListV2Externalcontact 获取群发记录列表
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发记录列表
func (c *App) GetGroupmsgListV2Externalcontact(req getGroupMsgListV2ExternalContactReq) (ok bool, err error) {
	var resp getGroupMsgListV2ExternalContactResp
	resp, err = c.execGetGroupmsgListV2Externalcontact(req)
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return
}

type StopSendMsgExternalContactReq struct {
	Msgid string `json:"msgid"`
}

// StopSendMsgExternalContact 停止群发任务
func (c *App) StopSendMsgExternalContact(MsgID string) (err error) {
	var resp struct {
		CommonResp
	}
	err = c.executeWXApiJSONPost("/cgi-bin/externalcontact/cancel_groupmsg_send",
		newIntoBodyer(StopSendMsgExternalContactReq{Msgid: MsgID}), &resp, true)
	if err != nil {
		return err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return bizErr
	}
	return nil
}
