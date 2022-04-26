package dingding

import (
	"context"
	"net/url"

	"github.com/lack-io/dingding-sdk/api"
)

// 公司部门相关接口

// ListDepartment 获取部门列表
func (c *Client) ListDepartment(ctx context.Context, req *api.ListDepartmentRequest) ([]*api.ListDepartmentResponse, error) {

	uri, err := url.Parse(api.URLListDepartment)
	if err != nil {
		return nil, err
	}

	if req.Language == "" {
		req.Language = "zh_CN"
	}

	depts := make([]*api.ListDepartmentResponse, 0)
	result := api.ResultV0{Result: &depts}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return depts, nil
}

// GetDepartment 获取部门详情
// 	deptId : 部门ID，根部门ID为1
// 		企业内部应用，可调用获取部门列表获取
// 		钉钉三方企业应用，可调用获取部门列表获取
//
//  language : 通讯录语言：
// 		zh_CN（默认）：中文
// 		en_US：英文
func (c *Client) GetDepartment(ctx context.Context, req *api.GetDepartmentRequest) (*api.DeptGetResponse, error) {

	uri, err := url.Parse(api.URLGetDepartment)
	if err != nil {
		return nil, err
	}

	if req.Language == "" {
		req.Language = "zh_CN"
	}

	dept := api.DeptGetResponse{}
	result := api.ResultV0{Result: &dept}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return &dept, nil
}

// ListDeptSubIds 获取子部门ID列表
// 	deptId : 部门ID，根部门ID为1
// 		企业内部应用，可调用获取部门列表获取
// 		钉钉三方企业应用，可调用获取部门列表获取
func (c *Client) ListDeptSubIds(ctx context.Context, req *api.DeptListSubIdRequest) (*api.DeptListSubIdResponse, error) {

	uri, err := url.Parse(api.URLListDeptSubIds)
	if err != nil {
		return nil, err
	}
	dept := api.DeptListSubIdResponse{}
	result := api.ResultV0{Result: &dept}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return &dept, nil
}

// ListParentByDept 获取父部门ID列表
func (c *Client) ListParentByDept(ctx context.Context, req *api.DeptListParentByDeptIdRequest) (*api.DeptListParentByDeptIdResponse, error) {

	uri, err := url.Parse(api.URLListParentByDept)
	if err != nil {
		return nil, err
	}

	dept := api.DeptListParentByDeptIdResponse{}
	result := api.ResultV0{Result: &dept}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return &dept, nil
}

// ListParentByUser 获取父部门ID列表
func (c *Client) ListParentByUser(ctx context.Context, req *api.DeptListParentByUserRequest) (*api.DeptListParentByUserResponse, error) {

	uri, err := url.Parse(api.URLListParentByUser)
	if err != nil {
		return nil, err
	}

	dept := api.DeptListParentByUserResponse{}
	result := api.ResultV0{Result: &dept}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return &dept, nil
}
