package workwx

import (
	"net/url"
	"strconv"
)

// AgentInfoRsp 获取企业应用详情
type AgentInfoRsp struct {
	CommonResp
	AgentInfo
}

type AgentInfoReq struct {
	AgentId int64
}

var _ urlValuer = AgentInfoReq{}

func (x AgentInfoReq) intoURLValues() url.Values {
	return url.Values{
		"agentid": {strconv.Itoa(int(x.AgentId))},
	}
}

// execGetAgentInfo 获取应用详情
func (c *App) execGetAgentInfo(req AgentInfoReq) (AgentInfoRsp, error) {
	var resp AgentInfoRsp
	err := c.executeWXApiGet("/cgi-bin/agent/get", req, &resp, true)
	if err != nil {
		return AgentInfoRsp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return AgentInfoRsp{}, bizErr
	}

	return resp, nil
}
