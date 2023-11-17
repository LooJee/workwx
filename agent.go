package workwx

func (c *App) GetAgentInfo(agentId int64) (info AgentInfo, err error) {
	rsp, err := c.execGetAgentInfo(AgentInfoReq{AgentId: agentId})
	if err != nil {
		return
	}

	return rsp.AgentInfo, nil
}

func (c *App) GetAgentPermission() (permission GetAgentPermission, err error) {
	type Response struct {
		CommonResp
		GetAgentPermission
	}

	var resp Response

	if err := c.executeWXApiJSONPost("/cgi-bin/agent/get_permissions", newIntoBodyer(struct{}{}), &resp, true); err != nil {
		return permission, err
	} else if err := resp.TryIntoErr(); err != nil {
		return permission, err
	}

	return resp.GetAgentPermission, nil
}
