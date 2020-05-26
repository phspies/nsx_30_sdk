package nsxt

// Size of Directory Domain
type DirectoryDomainSize struct {
	// Number of groups
	GroupCount int32 `json:"group_count,omitempty"`
	// Number of users
	UserCount int32 `json:"user_count,omitempty"`
	// Number of group members
	GroupMemberCount int32 `json:"group_member_count,omitempty"`
}
