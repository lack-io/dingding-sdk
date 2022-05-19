package api

import "time"

const (
	// URLGetAccessToken (v0 - post) 获取 token 信息
	URLGetAccessToken = "https://oapi.dingtalk.com/gettoken"

	// URLListExtContact (v0 - post) 获取外部联系人列表
	URLListExtContact = "https://oapi.dingtalk.com/topapi/extcontact/list"

	// URLListDepartment (v2 - post) 获取部门列表
	URLListDepartment = "https://oapi.dingtalk.com/topapi/v2/department/listsub"

	// URLGetDepartment (v2 - post) 获取部门详情
	URLGetDepartment = "https://oapi.dingtalk.com/topapi/v2/department/get"

	// URLListDeptSubIds (v2 - post) 获取子部门 id 列表
	URLListDeptSubIds = "https://oapi.dingtalk.com/topapi/v2/department/listsubid"

	// URLListParentByDept (v2 - post) 获取指定部门的所有父部门列表
	URLListParentByDept = "https://oapi.dingtalk.com/topapi/v2/department/listparentbydept"

	// URLListParentByUser (v2 - post) 获取指定用户的所有父部门列表
	URLListParentByUser = "https://oapi.dingtalk.com/topapi/v2/department/listparentbyuser"

	// URLListUserSimple (v1 - post) 获取部门用户基础信息
	URLListUserSimple = "https://oapi.dingtalk.com/topapi/user/listsimple"

	// URLGetUser (v2 - post) 获取用户详细信息
	URLGetUser = "https://oapi.dingtalk.com/topapi/v2/user/get"

	// URLListUserId (v1 - post) 获取部门用户userid列表
	URLListUserId = "https://oapi.dingtalk.com/topapi/user/listid"

	// URLListUser (v2 - post) 获取部门用户详情
	URLListUser = "https://oapi.dingtalk.com/topapi/v2/user/list"

	// URLGetUserCount (v1 - post) 获取员工人数
	URLGetUserCount = "https://oapi.dingtalk.com/topapi/user/count"

	// URLGETUserAccessToken (v1 - post) 获取用户 token
	URLGETUserAccessToken = "https://api.dingtalk.com/v1.0/oauth2/userAccessToken"

	// URLGETContactUser (v1 - get) 获取用户通讯录个人信息
	URLGETContactUser = "https://api.dingtalk.com/v1.0/contact/users/"

	// URLGETAuthInfos (v1 - get) 获取企业认证信息
	URLGETAuthInfos = "https://api.dingtalk.com/v1.0/contact/organizations/authInfos"

	// URLGETUserByUnionId (v1 - post) 根据用户unionId 获取用户 id
	URLGETUserByUnionId = "https://oapi.dingtalk.com/topapi/user/getbyunionid"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpireIn    int32  `json:"expire_in"`
	Expired     time.Time
}

type ResultV0 struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestId string `json:"request_id"`
	Results   any    `json:"results"`
	Result    any    `json:"result"`

	AccessToken
}

type Pager struct {
	Size   int32 `json:"size"`
	Offset int32 `json:"offset"`
}

type PageResult struct {
	HasMore    bool        `json:"has_more,omitempty"`
	NextCursor int32       `json:"next_cursor,omitempty"`
	List       interface{} `json:"list,omitempty"`
}
