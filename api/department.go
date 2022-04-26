package api

import "encoding/json"

type ListDepartmentRequest struct {
	// DeptId : 部门ID，根部门ID为1
	//	企业内部应用，可调用获取部门列表获取
	// 	钉钉三方企业应用，可调用获取部门列表获取
	DeptId int32 `json:"dept_id"`

	// Language : 通讯录语言：
	//	zh_CN（默认）：中文
	//	en_US：英文
	Language string `json:"language"`
}

// ListDepartmentResponse 公司部门信息
// https://open.dingtalk.com/document/orgapp-server/obtain-the-department-list-v2
type ListDepartmentResponse struct {
	// DeptId 部门 id
	DeptId int32 `json:"dept_id,omitempty"`

	// Name 部门名称
	Name string `json:"name,omitempty"`

	// ParentId 父部门
	ParentId int32 `json:"parent_id,omitempty"`

	// CreateDeptGroup 是否同步创建一个关联此部门的企业群
	CreateDeptGroup bool `json:"create_dept_group,omitempty"`

	// AutoAddUser 部门群已经创建后，有新人加入部门是否会自动加入该群
	AutoAddUser bool `json:"auto_add_user,omitempty"`

	// Ext 额外信息
	Ext string `json:"ext,omitempty"`
}

func (d ListDepartmentResponse) String() string {
	data, _ := json.Marshal(d)
	return string(data)
}

type GetDepartmentRequest struct {
	// DeptId : 部门ID，根部门ID为1
	//	企业内部应用，可调用获取部门列表获取
	// 	钉钉三方企业应用，可调用获取部门列表获取
	DeptId int32 `json:"dept_id"`

	// Language : 通讯录语言：
	//	zh_CN（默认）：中文
	//	en_US：英文
	Language string `json:"language"`
}

// DeptGetResponse 部门详细信息
// https://open.dingtalk.com/document/orgapp-server/query-department-details0-v2
type DeptGetResponse struct {
	// DeptId 部门 id
	DeptId int32 `json:"dept_id,omitempty"`

	// Name 部门名称
	Name string `json:"name,omitempty"`

	// ParentId 父部门
	ParentId int32 `json:"parent_id,omitempty"`

	// SourceIdentifier 部门标识
	SourceIdentifier string `json:"source_identifier"`

	// CreateDeptGroup 是否同步创建一个关联此部门的企业群
	CreateDeptGroup bool `json:"create_dept_group,omitempty"`

	// AutoAddUser 部门群已经创建后，有新人加入部门是否会自动加入该群
	AutoAddUser bool `json:"auto_add_user,omitempty"`

	// Extention 额外信息
	Extention string `json:"extention,omitempty"`

	// AutoApproveApply 是否默认同意加入该部门的申请
	AutoApproveApply bool `json:"auto_approve_apply,omitempty"`

	// FromUnionOrg 部门是否来自关联组织
	FromUnionOrg bool `json:"from_union_org,omitempty"`

	// Tags 教育部门标签
	Tags string `json:"tags,omitempty"`

	// Order 在父部门中的次序值
	Order int32 `json:"order,omitempty"`

	// 部门群 ID
	DeptGroupChatId string `json:"dept_group_chat_id,omitempty"`

	// GroupContainSubDept 部门群是否包含子部门
	GroupContainSubDept bool `json:"group_contain_sub_dept,omitempty"`

	// OrgDeptOwner 企业群群主 userId
	OrgDeptOwner string `json:"org_dept_owner,omitempty"`

	// DeptManagerUserIdList 部门的主管 userid 列表
	DeptManagerUserIdList []string `json:"dept_manager_userid_list"`

	// OuterDept 是否限制本部门成员查看通讯录
	OuterDept bool `json:"outer_dept,omitempty"`

	// OuterPermitDepts 配置的部门员工可见部门 id 列表
	OuterPermitDepts []int32 `json:"outer_permit_depts,omitempty"`

	// OuterPermitUsers 配置的部门员工可见与昂工 userId 列表
	OuterPermitUsers []int32 `json:"outer_permit_users,omitempty"`

	// HideDept 是否开启隐藏本部门
	HideDept bool `json:"hide_dept,omitempty"`

	// UserPermits 隐藏部门的员工 userId 列表
	UserPermits []string `json:"user_permits,omitempty"`

	// DeptPermits 隐藏部门的部门 id 列表
	DeptPermits []int32 `json:"dept_permits,omitempty"`
}

func (c DeptGetResponse) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}

type DeptListSubIdRequest struct {
	// DeptId : 部门ID，根部门ID为1
	//	企业内部应用，可调用获取部门列表获取
	// 	钉钉三方企业应用，可调用获取部门列表获取
	DeptId int32 `json:"dept_id"`
}

// DeptListSubIdResponse 子部门ID列表
// https://open.dingtalk.com/document/orgapp-server/obtain-a-sub-department-id-list-v2
type DeptListSubIdResponse struct {
	// DeptIdList 子部门列表
	DeptIdList []int32 `json:"dept_id_list,omitempty"`
}

func (c DeptListSubIdResponse) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}

type DeptListParentByDeptIdRequest struct {
	// DeptId : 部门ID，根部门ID为1
	//	企业内部应用，可调用获取部门列表获取
	// 	钉钉三方企业应用，可调用获取部门列表获取
	DeptId int32 `json:"dept_id"`
}

// DeptListParentByDeptIdResponse 父部门ID列表
// https://open.dingtalk.com/document/orgapp-server/query-the-list-of-all-parent-departments-of-a-department
type DeptListParentByDeptIdResponse struct {
	// ParentIdList 该部门的所有父部门ID列表
	ParentIdList []int32 `json:"parent_id_list,omitempty"`
}

func (c DeptListParentByDeptIdResponse) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}

type DeptParentResponse struct {
	// ParentDeptIdList 父部门列表
	ParentDeptIdList []int32 `json:"parent_dept_id_list,omitempty"`
}

type DeptListParentByUserRequest struct {
	// UserId 用户ID
	UserId string `json:"userid,omitempty"`
}

// DeptListParentByUserResponse 父部门ID列表
// https://open.dingtalk.com/document/orgapp-server/queries-the-list-of-all-parent-departments-of-a-user
type DeptListParentByUserResponse struct {
	// ParentList 父部门列表集合
	ParentList []DeptParentResponse `json:"parent_list,omitempty"`
}

func (c DeptListParentByUserResponse) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}

type ListExContactRequest struct {
	Pager
}

// ListExContactResponse 外部联系人
type ListExContactResponse struct {
	// Title 职位
	Title string `json:"title,omitempty"`

	// ShareDeptIds 共享部门 ID 列表
	ShareDeptIds []string `json:"share_dept_ids,omitempty"`

	// LabelIds 外部联系人标签列表
	LabelIds []int32 `json:"label_ids,omitempty"`

	// Remark 备注内容
	Remark string `json:"remark,omitempty"`

	// Address 地址
	Address string `json:"address,omitempty"`

	// Name 姓名
	Name string `json:"name,omitempty"`

	// FollowerUserId 负责人的userid。
	FollowerUserId string `json:"followerUserId,omitempty"`

	// StateCode 国家码
	StateCode string `json:"state_code,omitempty"`

	// CompanyName 公司名
	CompanyName string `json:"company_name,omitempty"`

	// ShareUserIds 共享员工userid列表
	ShareUserId []string `json:"share_user_ids,omitempty"`

	// Mobile 手机号
	Mobile string `json:"mobile,omitempty"`

	// UserId 外部联系人的userid
	UserId string `json:"userid,omitempty"`

	// Email 邮箱
	Email string `json:"email,omitempty"`
}

func (c ListExContactResponse) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}
