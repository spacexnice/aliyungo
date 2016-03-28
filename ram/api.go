package ram

/*
	ringtail 2016/1/19
	All RAM apis provided
*/

type RamClientInterface interface {
	//ram user
	CreateUser(user UserRequest) (UserResponse, error)
	GetUser(userQuery UserQueryRequest) (UserResponse, error)
	UpdateUser(newUser UpdateUserRequest) (UserResponse, error)
	DeleteUser(userQuery UserQueryRequest) (RamCommonResponse, error)
	ListUsers(listParams ListUserRequest) (ListUserResponse, error)

	//TODO login ram console
	CreateLoginProfile()
	GetLoginProfile()
	DeleteLoginProfile()
	UpdateLoginProfile()

	//ram ak
	CreateAccessKey(userQuery UserQueryRequest) (AccessKeyResponse, error)
	UpdateAccessKey(accessKeyRequest UpdateAccessKeyRequest) (RamCommonResponse, error)
	DeleteAccessKey(accessKeyRequest UpdateAccessKeyRequest) (RamCommonResponse, error)
	ListAccessKeys(userQuery UserQueryRequest) (AccessKeyListResponse, error)

	//TODO MFA
	CreateVirtualMFADevices()
	ListVirtualMFADevices()
	DeleteVirtualMFADevices()
	BindMFADevice()
	GetUserMFAInfo()

	//TODO group
	CreateGroup()
	GetGroup()
	UpdateGroup()
	ListGroup()
	DeleteGroup()
	AddUserToGroup()
	RemoveUserFromGroup()
	ListGroupsForUser()
	ListUsersForGroup()

	//TODO role
	CreateRole()
	GetRole()
	UpdateRole()
	ListRoles()
	DeleteRole()

	//DONE policy
	CreatePolicy(policyReq PolicyRequest) (PolicyResponse, error)
	GetPolicy(policyReq PolicyRequest) (PolicyResponse, error)
	DeletePolicy(policyReq PolicyRequest) (RamCommonResponse, error)
	ListPolicies(policyQuery PolicyQueryRequest) (PolicyQueryResponse, error)
	ListPoliciesForUser(userQuery UserQueryRequest) (PolicyListResponse, error)

	//TODO policy
	CreatePolicyVersion(policyReq PolicyRequest) (PolicyVersionResponse, error)
	GetPolicyVersion(policyReq PolicyRequest) (PolicyVersionResponse, error)
	DeletePolicyVersion(policyReq PolicyRequest) (RamCommonResponse, error)
	ListPolicyVersions(policyReq PolicyRequest) (PolicyVersionsResponse, error)
	AttachPolicyToUser(attachPolicyRequest AttachPolicyRequest) (RamCommonResponse, error)
	DetachPolicyFromUser(attachPolicyRequest AttachPolicyRequest) (RamCommonResponse, error)
	ListEnitiesForPolicy()
	SetDefaultPolicyVersion()
	ListPoliciesForGroup()
	ListPoliciesForRole()

	//TODO security apis
	SetAccountAlias(accountAlias AccountAlias) (RamCommonResponse, error)
	GetAccountAlias() (AccountAliasResponse, error)
	ClearAccountAlias() (RamCommonResponse, error)
	SetPasswordPolicy(passwordPolicy PasswordPolicyRequest) (PasswordPolicyResponse, error)
	GetPasswordPolicy(accountAlias AccountAlias) (PasswordPolicyResponse, error)
}
