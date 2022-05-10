package dingding

import (
	"context"
	"net/url"
	"time"

	"github.com/lack-io/dingding-sdk/api"
)

// GetAuthInfos 获取企业认证信息
func (c *Client) GetAuthInfos(ctx context.Context, req *api.GetAuthInfosRequest) (*api.GetAuthInfosResponse, error) {

	uri, err := url.Parse(api.URLGETAuthInfos)
	if err != nil {
		return nil, err
	}

	query := uri.Query()
	query.Set("targetCorpId", req.CorpId)
	uri.RawQuery = query.Encode()

	result := api.GetAuthInfosResponse{}

	tok, err := c.getToken(ctx, false)
	if err != nil {
		return nil, err
	}

	if time.Now().After(tok.Expired) {
		tok, err = c.getToken(ctx, true)
		if err != nil {
			return nil, err
		}
	}

	header := map[string]string{
		"x-acs-dingtalk-access-token": tok.AccessToken,
	}

	err = c.fetchV1(ctx, "GET", uri, header, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
