package workwx

import (
	"net/url"
	"strconv"
)

// DeptInfo 部门信息
type DeptInfo struct {
	// ID 部门 ID
	ID int64 `json:"id"`
	// Name 部门名称
	Name string `json:"name"`
	// ParentID 父亲部门id。根部门为1
	ParentID int64 `json:"parentid"`
	// Order 在父部门中的次序值。order值大的排序靠前。值范围是[0, 2^32)
	Order uint32 `json:"order"`
}

type GetDeptInfoReq struct {
	DeptId int64
}

var _ urlValuer = GetDeptInfoReq{}

func (x GetDeptInfoReq) intoURLValues() url.Values {
	return url.Values{
		"id": {strconv.Itoa(int(x.DeptId))},
	}
}
