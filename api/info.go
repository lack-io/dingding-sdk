package api

type GetAuthInfosRequest struct {
	// 需要获取的企业认证信息的企业corpId
	CorpId string `json:"corp_id"`
}

// GetAuthInfosResponse 获取企业认证信息
// https://open.dingtalk.com/document/isvapp-server/obtain-enterprise-authentication-information
type GetAuthInfosResponse struct {
	OrgName string `json:"orgName"`

	LicenseOrgName string `json:"licenseOrgName"`

	RegistrationNum string `json:"registrationNum"`

	UnifiedSocialCredit string `json:"unifiedSocialCredit"`

	OrganizationCode string `json:"organizationCode"`

	LegalPerson string `json:"legalPerson"`

	LicenseUrl string `json:"licenseUrl"`

	// 认证等级，取值：
	//   1：高级认证
	//   2：中级认证
	AuthLevel int64 `json:"authLevel"`
}
