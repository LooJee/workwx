package workwx

import (
	"context"
	"github.com/pkg/errors"
)

type SaveCustomerAcquisitionReq struct {
	LinkID     string                   `json:"link_id,omitempty"`
	LinkName   string                   `json:"link_name" binding:"required"`
	Range      CustomerAcquisitionRange `json:"range"`
	SkipVerify bool                     `json:"skip_verify"`
}
type CustomerAcquisitionRange struct {
	UserList       []string `json:"user_list" binding:"required"`
	DepartmentList []int64  `json:"department_list" binding:"required"`
}
type SaveCustomerAcquisitionResp struct {
	LinkID     string `json:"link_id"`
	LinkName   string `json:"link_name"`
	Url        string `json:"url"`
	CreateTime int64  `json:"create_time"`
}

func (c *App) SaveCustomerAcquisition(req *SaveCustomerAcquisitionReq) (link SaveCustomerAcquisitionResp, err error) {
	var path = "/cgi-bin/externalcontact/customer_acquisition/create_link"
	if req.LinkID != "" {
		path = "/cgi-bin/externalcontact/customer_acquisition/update_link"
	}
	var resp = struct {
		CommonResp
		Link SaveCustomerAcquisitionResp `json:"link"`
	}{}
	err = c.executeWXApiJSONPost(path, newIntoBodyer(req), &resp, true)
	if err != nil {
		return SaveCustomerAcquisitionResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return SaveCustomerAcquisitionResp{}, bizErr
	}

	return resp.Link, nil
}

type DelCustomerAcquisitionReq struct {
	LinkID string `json:"link_id"`
}

func (c *App) DelCustomerAcquisition(req *DelCustomerAcquisitionReq) (err error) {
	var path = "/cgi-bin/externalcontact/customer_acquisition/delete_link"
	var resp = struct {
		CommonResp
	}{}
	err = c.executeWXApiJSONPost(path, newIntoBodyer(req), &resp, true)
	if err != nil {
		return err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return bizErr
	}

	return nil
}

type ReqGetCustomerAcquisitionList struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type RspGetCustomerAcquisitionList struct {
	NextCursor string   `json:"next_cursor"`
	LinkIdList []string `json:"link_id_list"`
}

// GetCustomerAcquisitionList 获取获客链接列表
func (c *App) GetCustomerAcquisitionList(ctx context.Context, req *ReqGetCustomerAcquisitionList) (data RspGetCustomerAcquisitionList, err error) {
	var path = "/cgi-bin/externalcontact/customer_acquisition/list_link"
	var resp = struct {
		CommonResp
		RspGetCustomerAcquisitionList
	}{}

	err = c.executeWXApiJSONPost(path, newIntoBodyer(req), &resp, true)
	if err != nil {
		return RspGetCustomerAcquisitionList{}, errors.WithStack(err)
	}

	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RspGetCustomerAcquisitionList{}, errors.WithStack(bizErr)
	}

	return resp.RspGetCustomerAcquisitionList, nil
}

type ReqGetCustomerAcquisitionDetail struct {
	LinkId string `json:"link_id"`
}

type CALink struct {
	LinkName   string `json:"link_name"`
	Url        string `json:"url"`
	CreateTime int64  `json:"create_time"`
	SkipVerify bool   `json:"skip_verify"`
}

type CARange struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int64  `json:"department_list"`
}

type RspGetCustomerAcquisitionDetail struct {
	Link  CALink  `json:"link"`
	Range CARange `json:"range"`
}

func (c *App) GetCustomerAcquisitionDetail(ctx context.Context, req *ReqGetCustomerAcquisitionDetail) (data RspGetCustomerAcquisitionDetail, err error) {
	var path = "/cgi-bin/externalcontact/customer_acquisition/get"
	var resp = struct {
		CommonResp
		RspGetCustomerAcquisitionDetail
	}{}

	err = c.executeWXApiJSONPost(path, newIntoBodyer(req), &resp, true)
	if err != nil {
		return RspGetCustomerAcquisitionDetail{}, errors.WithStack(err)
	}

	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RspGetCustomerAcquisitionDetail{}, errors.WithStack(bizErr)
	}

	return resp.RspGetCustomerAcquisitionDetail, nil
}
