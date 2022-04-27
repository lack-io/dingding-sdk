package dingding

import (
	"context"
	"os"
	"testing"

	"github.com/lack-io/dingding-sdk/api"
)

func newClient() *Client {
	appKey, _ := os.LookupEnv("ddAppKey")
	appSecret, _ := os.LookupEnv("ddAppSecret")
	token, _ := os.LookupEnv("ddToken")

	return NewClient(
		WithAppKey(appKey),
		WithAppSecret(appSecret),
		WithToken(token),
	)
}

func TestClient_GetAccessToken(t *testing.T) {

	cli := newClient()
	token, err := cli.GetAccessToken(context.TODO())
	if err != nil {
		t.Errorf("get token: %v", err)
	}
	t.Log(token)
}

func TestClient_ListExtContact(t *testing.T) {
	cli := newClient()
	contacts, err := cli.ListExtContact(context.TODO(), &api.ListExContactRequest{Pager: api.Pager{Size: 20}})
	if err != nil {
		t.Errorf("list ext contact: %v", err)
	}
	t.Log(contacts)
}

func TestClient_ListDepartment(t *testing.T) {
	cli := newClient()
	out, err := cli.ListDepartment(context.TODO(), &api.ListDepartmentRequest{
		DeptId:   1,
		Language: "zh_CN",
	})
	if err != nil {
		t.Errorf("list department: %v", err)
	}
	t.Log(out)
}

func TestClient_GetDepartment(t *testing.T) {
	cli := newClient()
	out, err := cli.GetDepartment(context.TODO(), &api.GetDepartmentRequest{DeptId: 7957684})
	if err != nil {
		t.Errorf("get department: %v", err)
	}
	t.Log(out)
}

func TestClient_ListDeptSubIds(t *testing.T) {
	cli := newClient()
	out, err := cli.ListDeptSubIds(context.TODO(), &api.DeptListSubIdRequest{DeptId: 1})
	if err != nil {
		t.Errorf("list department suids: %v", err)
	}
	t.Log(out)
}

func TestClient_ListParentByDept(t *testing.T) {
	cli := newClient()
	out, err := cli.ListParentByDept(context.TODO(), &api.DeptListParentByDeptIdRequest{DeptId: 7957684})
	if err != nil {
		t.Errorf("list parent by department: %v", err)
	}
	t.Log(out)
}

func TestClient_ListParentByUser(t *testing.T) {
	cli := newClient()
	out, err := cli.ListParentByUser(context.TODO(), &api.DeptListParentByUserRequest{UserId: "16337403375506420"})
	if err != nil {
		t.Errorf("list parent by user: %v", err)
	}
	t.Log(out)
}

func TestClient_ListUserSimple(t *testing.T) {
	cli := newClient()
	out, err := cli.ListUserSimple(context.TODO(), &api.ListUserSimpleRequest{
		DeptId: 7957684,
		Cursor: 0,
		Size:   10,
	})
	if err != nil {
		t.Errorf("list user simple: %v", err)
	}
	t.Log(out)
}

func TestClient_GetUser(t *testing.T) {
	cli := newClient()
	out, err := cli.GetUser(context.TODO(), &api.GetUserRequest{
		UserId: "16337403375506420",
	})
	if err != nil {
		t.Errorf("get user: %v", err)
	}
	t.Log(out)
}

func TestClient_ListUserId(t *testing.T) {
	cli := newClient()
	out, err := cli.ListUserId(context.TODO(), &api.ListUserIdRequest{
		DeptId: 7957684,
	})
	if err != nil {
		t.Errorf("list userid: %v", err)
	}
	t.Log(out)
}

func TestClient_ListUser(t *testing.T) {
	cli := newClient()
	out, err := cli.ListUser(context.TODO(), &api.ListUserRequest{
		DeptId: 7957684,
		Cursor: 0,
		Size:   10,
	})
	if err != nil {
		t.Errorf("list user: %v", err)
	}
	t.Log(out)
}

func TestClient_GetUserCount(t *testing.T) {
	cli := newClient()
	out, err := cli.GetUserCount(context.TODO(), &api.GetUserCountRequest{})
	if err != nil {
		t.Errorf("get user count: %v", err)
	}
	t.Log(out)
}

func TestClient_GetUserAccessToken(t *testing.T) {
	cli := newClient()
	out, err := cli.GetUserAccessToken(context.TODO(), &api.GetUserAccessTokenRequest{
		Code:         "580dfd7103ad3f839d033d74efbb63b6",
		RefreshToken: "580dfd7103ad3f839d033d74efbb63b6",
		GrantType:    "authorization_code",
	})
	if err != nil {
		t.Errorf("get user token: %v", err)
	}
	t.Log(out)
}

func TestClient_GetContactUser(t *testing.T) {
	cli := newClient()
	out, err := cli.GetContactUser(context.TODO(), &api.GetContactUserRequest{
		Token:   "0269d878c58e36f28dc27e002a83258a",
		UnionId: "me",
	})
	if err != nil {
		t.Errorf("get contact user: %v", err)
	}
	t.Log(out)
}
