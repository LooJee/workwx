package workwx

func (c *App) GetAgentInfo(agentId int64) (info AgentInfo, err error) {
	rsp, err := c.execGetAgentInfo(AgentInfoReq{AgentId: agentId})
	if err != nil {
		return
	}

	return rsp.AgentInfo, nil
}
