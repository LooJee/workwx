package workwx

import (
	"github.com/pkg/errors"
)

// ListGroupChat 获取客户群列表
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92120#获取客户群列表
func (c *App) ListGroupChat(req ListGroupChatReq) (ListGroupChatResp, error) {
	var resp ListGroupChatResp
	resp, err := c.execListGroupChat(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetGroupChat 获取客户群详情
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92122#获取客户群详情
func (c *App) GetGroupChat(req GetGroupChatReq) (GetGroupChatResp, error) {
	resp, err := c.execGetGroupChat(req)
	if err != nil {
		return GetGroupChatResp{}, err
	}
	return resp, err
}

// GroupChatAddJoinWay 配置客户群进群方式
// 文档：https://developer.work.weixin.qq.com/document/path/92229#%E9%85%8D%E7%BD%AE%E5%AE%A2%E6%88%B7%E7%BE%A4%E8%BF%9B%E7%BE%A4%E6%96%B9%E5%BC%8F
func (c *App) GroupChatAddJoinWay(req GroupChatAddJoinWayReq) (string, error) {
	resp, err := c.execGroupChatAddJoinWay(req)
	if err != nil {
		return "", err
	}

	return resp.ConfigId, err
}

// GroupChatGetJoinWay 获取客户群进群方式
// 文档：https://developer.work.weixin.qq.com/document/path/92229#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E7%BE%A4%E8%BF%9B%E7%BE%A4%E6%96%B9%E5%BC%8F%E9%85%8D%E7%BD%AE
func (c *App) GroupChatGetJoinWay(req GroupChatGetJoinWayReq) (GroupChatGetJoinWay, error) {
	resp, err := c.execGroupChatGetJoinWay(req)
	if err != nil {
		return GroupChatGetJoinWay{}, err
	}
	return resp.JoinWay, err
}

// GroupChatUpdateJoinWay 更新客户群进群方式配置
// 文档：https://developer.work.weixin.qq.com/document/path/92229#%E6%9B%B4%E6%96%B0%E5%AE%A2%E6%88%B7%E7%BE%A4%E8%BF%9B%E7%BE%A4%E6%96%B9%E5%BC%8F%E9%85%8D%E7%BD%AE
func (c *App) GroupChatUpdateJoinWay(req GroupChatUpdateJoinWayReq) error {
	_, err := c.execGroupChatUpdateJoinWay(req)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// GroupChatDelJoinWay 删除客户群进群方式配置
// 文档：https://developer.work.weixin.qq.com/document/path/92229#%E5%88%A0%E9%99%A4%E5%AE%A2%E6%88%B7%E7%BE%A4%E8%BF%9B%E7%BE%A4%E6%96%B9%E5%BC%8F%E9%85%8D%E7%BD%AE
func (c *App) GroupChatDelJoinWay(req GroupChatDelJoinWayReq) error {
	_, err := c.execGroupChatDelJoinWay(req)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
