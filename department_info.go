package workwx

// ListAllDepartments 获取全量组织架构。
func (c *App) ListAllDepartments() ([]*DeptInfo, error) {
	resp, err := c.execDeptList(deptListReq{
		HaveID: false,
		ID:     0,
	})
	if err != nil {
		return nil, err
	}

	return resp.Department, nil
}

// ListDepartments 获取指定部门及其下的子部门。
func (c *App) ListDepartments(id int64) ([]*DeptInfo, error) {
	resp, err := c.execDeptList(deptListReq{
		HaveID: true,
		ID:     id,
	})
	if err != nil {
		return nil, err
	}

	return resp.Department, nil
}

func (c *App) GetDepartment(id int64) (*DeptInfo, error) {
	type Response struct {
		CommonResp
		Department DeptInfo `json:"department"`
	}

	var resp Response
	err := c.executeWXApiGet("/cgi-bin/department/get", GetDeptInfoReq{DeptId: id}, &resp, true)
	if err != nil {
		return &resp.Department, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return &resp.Department, bizErr
	}
	return &resp.Department, nil
}
