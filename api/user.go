package api

import "encoding/json"

// ListUserSimpleRequest 获取部门用户基础信息
// https://open.dingtalk.com/document/orgapp-server/queries-the-simple-information-of-a-department-user
type ListUserSimpleRequest struct {
	// DeptId 部门ID，如果是根部门，该参数传1。
	//	企业内部应用，可调用获取部门列表获取
	//	钉钉三方企业应用，可调用获取部门列表获取
	// +required
	DeptId int32 `json:"dept_id"`

	// Cursor 分页查询的游标，最开始传0，后续传返回参数中的next_cursor值
	// +required
	Cursor int32 `json:"cursor"`

	// Size 分页长度，最大值100
	// +required
	Size int32 `json:"size"`

	// OrderField 部门成员的排序规则：
	//	entry_asc：代表按照进入部门的时间升序。
	//	entry_desc：代表按照进入部门的时间降序。
	//	modify_asc：代表按照部门信息修改时间升序。
	//	modify_desc：代表按照部门信息修改时间降序。
	//	custom：代表用户定义(未定义时按照拼音)排序。
	// 默认值：custom。
	OrderField string `json:"order_field,omitempty"`

	// ContainAccessLimit 是否返回访问受限的员工
	ContainAccessLimit bool `json:"contain_access_limit,omitempty"`

	// Language 通讯录语言，取值。
	//	zh_CN：中文（默认值）。
	//	en_US：英文。
	Language string `json:"language,omitempty"`
}

// ListUserSimpleResponse 获取部门用户基础信息
type ListUserSimpleResponse struct {
	PageResult
}

type UserSimple struct {
	UserId string `json:"userid,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (u UserSimple) String() string {
	data, _ := json.Marshal(u)
	return string(data)
}

// GetUserRequest 获取用户详细信息
// https://open.dingtalk.com/document/orgapp-server/query-user-details
type GetUserRequest struct {
	// +required
	UserId   string `json:"userid"`
	Language string `json:"language,omitempty"`
}

// GetUserResponse 获取用户详细信息
type GetUserResponse struct {
	// UserId 员工的userId
	UserId string `json:"userid,omitempty"`

	// UnionId 员工在当前开发者企业账号范围内的唯一标识
	UnionId string `json:"unionid,omitempty"`

	// Name 员工姓名
	Name string `json:"name,omitempty"`

	// Avatar 员工头像
	Avatar string `json:"avatar,omitempty"`

	StateCode string `json:"state_code,omitempty"`

	ManagerUserId string `json:"manager_userid,omitempty"`

	Mobile string `json:"mobile,omitempty"`

	HideMobile bool `json:"hideMobile,omitempty"`

	Telephone string `json:"telephone,omitempty"`

	JobNumber string `json:"job_number,omitempty"`

	Title string `json:"title,omitempty"`

	Email string `json:"email,omitempty"`

	WorkPlace string `json:"work_place,omitempty"`

	Remark string `json:"remark,omitempty"`

	ExclusiveAccount bool `json:"exclusive_account,omitempty"`

	OrgEmail string `json:"org_email,omitempty"`

	DeptIdList []int32 `json:"dept_id_list,omitempty"`

	DeptOrderList []struct {
		DeptId int32 `json:"dept_id,omitempty"`
		Order  int64 `json:"order,omitempty"`
	} `json:"dept_order_list,omitempty"`

	Extension string `json:"extension,omitempty"`

	HiredDate int64 `json:"hired_date,omitempty"`

	Active bool `json:"active,omitempty"`

	RealAuthed bool `json:"real_authed,omitempty"`

	OrgEmailType string `json:"org_email_type,omitempty"`

	Senior bool `json:"senior,omitempty"`

	Admin bool `json:"admin,omitempty"`

	Boss bool `json:"boss,omitempty"`

	LeaderInDept []struct {
		DeptId int32 `json:"dept_id,omitempty"`
		Leader bool  `json:"leader,omitempty"`
	} `json:"leader_in_dept,omitempty"`

	RoleList []struct {
		Id        int32  `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		GroupName string `json:"group_name,omitempty"`
	} `json:"role_list,omitempty"`

	UnionEmpExt *UnionEmpExt `json:"union_emp_ext,omitempty"`
}

func (u GetUserResponse) String() string {
	data, _ := json.Marshal(u)
	return string(data)
}

type UnionEmpExt struct {
	UserId string `json:"userid,omitempty"`

	UnionEmpMapList []struct {
		UserId string `json:"userid,omitempty"`
		CorpId string `json:"corp_id,omitempty"`
	} `json:"union_emp_map_list,omitempty"`

	CorpId string `json:"corp_id,omitempty"`
}

func (u UnionEmpExt) String() string {
	data, _ := json.Marshal(u)
	return string(data)
}

// ListUserIdRequest 获取部门用户ID
// https://open.dingtalk.com/document/orgapp-server/query-the-list-of-department-userids
type ListUserIdRequest struct {
	// +required
	DeptId int32 `json:"dept_id"`
}

// ListUserIdResponse 获取部门用户ID
type ListUserIdResponse struct {
	UserIdList []string `json:"userid_list,omitempty"`
}

// ListUserRequest 获取部门用户详细信息
// https://open.dingtalk.com/document/orgapp-server/queries-the-complete-information-of-a-department-user
type ListUserRequest struct {
	// DeptId 部门ID，如果是根部门，该参数传1。
	//	企业内部应用，可调用获取部门列表获取
	//	钉钉三方企业应用，可调用获取部门列表获取
	// +required
	DeptId int32 `json:"dept_id"`

	// Cursor 分页查询的游标，最开始传0，后续传返回参数中的next_cursor值
	// +required
	Cursor int32 `json:"cursor"`

	// Size 分页长度，最大值100
	// +required
	Size int32 `json:"size"`

	// OrderField 部门成员的排序规则：
	//	entry_asc：代表按照进入部门的时间升序。
	//	entry_desc：代表按照进入部门的时间降序。
	//	modify_asc：代表按照部门信息修改时间升序。
	//	modify_desc：代表按照部门信息修改时间降序。
	//	custom：代表用户定义(未定义时按照拼音)排序。
	// 默认值：custom。
	OrderField string `json:"order_field,omitempty"`

	// ContainAccessLimit 是否返回访问受限的员工
	ContainAccessLimit bool `json:"contain_access_limit,omitempty"`

	// Language 通讯录语言，取值。
	//	zh_CN：中文（默认值）。
	//	en_US：英文。
	Language string `json:"language,omitempty"`
}

// ListUserResponse 获取部门用户详细信息
type ListUserResponse struct {
	PageResult
}

type User struct {
	// UserId 员工的userId
	UserId string `json:"userid,omitempty"`

	// UnionId 员工在当前开发者企业账号范围内的唯一标识
	UnionId string `json:"unionid,omitempty"`

	// Name 员工姓名
	Name string `json:"name,omitempty"`

	// Avatar 员工头像
	Avatar string `json:"avatar,omitempty"`

	StateCode string `json:"state_code,omitempty"`

	ManagerUserId string `json:"manager_userid,omitempty"`

	Mobile string `json:"mobile,omitempty"`

	HideMobile bool `json:"hideMobile,omitempty"`

	Telephone string `json:"telephone,omitempty"`

	JobNumber string `json:"job_number,omitempty"`

	Title string `json:"title,omitempty"`

	Email string `json:"email,omitempty"`

	WorkPlace string `json:"work_place,omitempty"`

	Remark string `json:"remark,omitempty"`

	OrgEmail string `json:"org_email,omitempty"`

	DeptIdList []int32 `json:"dept_id_list,omitempty"`

	DeptOrder int64 `json:"dept_order,omitempty"`

	Extension string `json:"extension,omitempty"`

	HiredDate int64 `json:"hired_date,omitempty"`

	Active bool `json:"active,omitempty"`

	RealAuthed bool `json:"real_authed,omitempty"`

	OrgEmailType string `json:"org_email_type,omitempty"`

	Senior bool `json:"senior,omitempty"`

	Admin bool `json:"admin,omitempty"`

	Boss bool `json:"boss,omitempty"`

	ExclusiveAccount bool `json:"exclusive_account,omitempty"`

	LoginId string `json:"login_id,omitempty"`

	// ExclusiveAccountType dingtalk|sso
	ExclusiveAccountType string `json:"exclusive_account_type,omitempty"`

	Nickname string `json:"nickname,omitempty"`

	ExclusiveAccountCorpName string `json:"exclusive_account_corp_name,omitempty"`

	ExclusiveAccountCorpId string `json:"exclusive_account_corp_id,omitempty"`
}

func (u User) String() string {
	data, _ := json.Marshal(u)
	return string(data)
}

// GetUserCountRequest 获取员工人数
// https://open.dingtalk.com/document/orgapp-server/obtain-the-number-of-employees-v2
type GetUserCountRequest struct {
	OnlyActive bool `json:"only_active"`
}

// GetUserCountResponse 获取员工人数
type GetUserCountResponse struct {
	Count int32 `json:"count,omitempty"`
}

// GetUserAccessTokenRequest 获取用户 token
// https://open.dingtalk.com/document/orgapp-server/obtain-user-token
type GetUserAccessTokenRequest struct {
	// +required
	ClientId string `json:"clientId"`

	// +required
	ClientSecret string `json:"clientSecret"`

	// +required
	Code string `json:"code"`

	RefreshToken string `json:"refreshToken"`

	// 如果使用授权码换token，传authorization_code。
	//
	// 如果使用刷新token换用户token，传refresh_token。
	// +required
	GrantType string `json:"grantType"`
}

// GetUserAccessTokenResponse 获取用户 token
type GetUserAccessTokenResponse struct {
	// 生成的accessToken。
	AccessToken string `json:"accessToken,omitempty"`
	// 生成的refresh_token。可以使用此刷新token，定期的获取用户的accessToken
	RefreshToken string `json:"refreshToken,omitempty"`
	// 超时时间，单位秒。
	ExpireIn int64 `json:"expireIn,omitempty"`
	// 所选企业corpId。
	CorpId string `json:"corpId,omitempty"`
}

func (r GetUserAccessTokenResponse) String() string {
	data, _ := json.Marshal(r)
	return string(data)
}

// GetContactUserRequest 获取用户通讯录个人信息
// https://open.dingtalk.com/document/orgapp-server/dingtalk-retrieve-user-information
type GetContactUserRequest struct {
	// 调用服务端接口的授权凭证。使用个人用户的accessToken
	// +required
	Token string

	// 用户的unionId, 如需获取当前授权人的信息，unionId参数可以传me。
	// +required
	UnionId string
}

// GetContactUserResponse 获取用户通讯录个人信息
type GetContactUserResponse struct {
	Nick string `json:"nick,omitempty"`

	AvatarUrl string `json:"avatarUrl,omitempty"`

	Mobile string `json:"mobile,omitempty"`

	OpenId string `json:"openId,omitempty"`

	UnionId string `json:"unionId,omitempty"`

	Email string `json:"email,omitempty"`

	StateCode string `json:"stateCode,omitempty"`
}

func (r GetContactUserResponse) String() string {
	data, _ := json.Marshal(r)
	return string(data)
}
