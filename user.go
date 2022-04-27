package dingding

import (
	"context"
	"net/url"

	"github.com/lack-io/dingding-sdk/api"
)

// ListUserSimple 获取用户简单信息
func (c *Client) ListUserSimple(ctx context.Context, req *api.ListUserSimpleRequest) ([]*api.UserSimple, error) {

	uri, err := url.Parse(api.URLListUserSimple)
	if err != nil {
		return nil, err
	}

	if req.Language == "" {
		req.Language = "zh_CN"
	}

	users := make([]*api.UserSimple, 0)
	rsp := api.ListUserSimpleResponse{}
	rsp.List = &users
	result := api.ResultV0{Result: &rsp}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser 获取用户完整信息
func (c *Client) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {

	uri, err := url.Parse(api.URLGetUser)
	if err != nil {
		return nil, err
	}

	if req.Language == "" {
		req.Language = "zh_CN"
	}

	user := api.GetUserResponse{}
	result := api.ResultV0{Result: &user}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ListUserId 获取部门用户ID
func (c *Client) ListUserId(ctx context.Context, req *api.ListUserIdRequest) ([]string, error) {

	uri, err := url.Parse(api.URLListUserId)
	if err != nil {
		return nil, err
	}

	rsp := api.ListUserIdResponse{}
	result := api.ResultV0{Result: &rsp}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return rsp.UserIdList, nil
}

// ListUser 获取部门用户详细信息
func (c *Client) ListUser(ctx context.Context, req *api.ListUserRequest) ([]*api.User, error) {

	uri, err := url.Parse(api.URLListUser)
	if err != nil {
		return nil, err
	}

	users := make([]*api.User, 0)
	rsp := api.ListUserResponse{}
	rsp.List = &users
	result := api.ResultV0{Result: &rsp}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserCount 获取员工人数
func (c *Client) GetUserCount(ctx context.Context, req *api.GetUserCountRequest) (int32, error) {

	uri, err := url.Parse(api.URLGetUserCount)
	if err != nil {
		return 0, err
	}

	rsp := api.GetUserCountResponse{}
	result := api.ResultV0{Result: &rsp}

	err = c.retryV0(ctx, "POST", uri, req, &result, DefaultRetryFn)
	if err != nil {
		return 0, err
	}

	return rsp.Count, nil
}

// GetUserAccessToken 获取用户 Token
func (c *Client) GetUserAccessToken(ctx context.Context, req *api.GetUserAccessTokenRequest) (*api.GetUserAccessTokenResponse, error) {
	uri, err := url.Parse(api.URLGETUserAccessToken)
	if err != nil {
		return nil, err
	}

	if req.ClientId == "" {
		req.ClientId = c.options.appKey
	}
	if req.ClientSecret == "" {
		req.ClientSecret = c.options.appSecret
	}

	rsp := api.GetUserAccessTokenResponse{}

	err = c.fetchV1(ctx, "POST", uri, map[string]string{}, req, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}

// GetContactUser 获取用户通讯录个人信息
func (c *Client) GetContactUser(ctx context.Context, req *api.GetContactUserRequest) (*api.GetContactUserResponse, error) {

	uri, err := url.Parse(api.URLGETContactUser + req.UnionId)
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"x-acs-dingtalk-access-token": req.Token,
	}
	rsp := api.GetContactUserResponse{}

	err = c.fetchV1(ctx, "GET", uri, header, req, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
