package dingding

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/lack-io/dingding-sdk/api"
)

var DefaultRetryFn = func(err error) bool {
	if e, ok := err.(*DDErrorV0); ok {
		// 过期的 access_token
		if e.Code == "40014" {
			return true
		}
	}
	return false
}

type Client struct {
	options Options

	conn *http.Client

	token *atomic.Value
}

func NewClient(opts ...Option) *Client {

	options := NewOptions(opts...)

	cli := &Client{options: options}

	transport := &http.Transport{
		WriteBufferSize: options.writeSize,
		ReadBufferSize:  options.readSize,
	}

	cli.conn = &http.Client{
		Transport: transport,
		Timeout:   options.timeout,
	}

	cli.token = &atomic.Value{}
	cli.token.Store(&api.AccessToken{
		AccessToken: options.token,
		ExpireIn:    7200,
		Expired:     time.Now().Add(time.Second * 7200),
	})

	return cli
}

// GetAccessToken 获取钉钉 access_token
func (c *Client) GetAccessToken(ctx context.Context) (*api.AccessToken, error) {

	u, err := url.Parse(api.URLGetAccessToken)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("appkey", c.options.appKey)
	query.Set("appsecret", c.options.appSecret)
	u.RawQuery = query.Encode()

	result := api.ResultV0{}
	err = c.fetchV0(ctx, "GET", u, map[string]interface{}{}, &result)
	if err != nil {
		return nil, err
	}
	token := &result.AccessToken
	token.Expired = time.Now().Add(time.Second * time.Duration(token.ExpireIn))

	return &result.AccessToken, nil
}

// ListExtContact 获取外部联系人
func (c *Client) ListExtContact(ctx context.Context, req *api.ListExContactRequest) ([]*api.ListExContactResponse, error) {

	u, err := url.Parse(api.URLListExtContact)
	if err != nil {
		return nil, err
	}

	if req.Offset < 0 {
		req.Offset = 0
	}

	contact := make([]*api.ListExContactResponse, 0)
	result := api.ResultV0{Results: &contact}

	err = c.retryV0(ctx, "POST", u, req, &result, DefaultRetryFn)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

// UpdateToken 更新内部 token 并返回
func (c *Client) UpdateToken(ctx context.Context) (*api.AccessToken, error) {
	return c.getToken(ctx, true)
}

func (c *Client) getToken(ctx context.Context, force bool) (*api.AccessToken, error) {
	token := c.token.Load().(*api.AccessToken)
	if token == nil || token.AccessToken == "" || token.Expired.Sub(time.Now()) < 10*time.Second || force {
		tok, err := c.GetAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		c.setToken(tok)
		token = tok
	}
	return token, nil
}

func (c *Client) setToken(token *api.AccessToken) {
	c.token.Store(token)
}

func (c *Client) retryV0(
	ctx context.Context, method string, uri *url.URL,
	payload interface{}, result *api.ResultV0,
	retry func(err error) bool,
) error {
	token, err := c.getToken(ctx, false)
	if err != nil {
		return err
	}

	query := uri.Query()
	query.Set("access_token", token.AccessToken)
	uri.RawQuery = query.Encode()

	err = c.fetchV0(ctx, method, uri, payload, result)
	if err != nil {
		if !retry(err) {
			return err
		}
		token, err = c.getToken(ctx, true)
		if err != nil {
			return err
		}
		query := uri.Query()
		query.Set("access_token", token.AccessToken)
		uri.RawQuery = query.Encode()
	}
	return c.fetchV0(ctx, method, uri, payload, result)
}

// fetchV0 支持钉钉旧版接口
func (c *Client) fetchV0(
	ctx context.Context, method string, uri *url.URL,
	payload interface{}, result *api.ResultV0,
) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, uri.String(), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	rsp, err := c.conn.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("bad response body: %v", err)
	}
	if err = ParseResultV0(result); err != nil {
		return err
	}

	return nil
}
