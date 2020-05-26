package nsxt

// Authenticated User Info
type UserInfo struct {
	// User Name
	UserName string `json:"user_name,omitempty"`
	// Permissions
	Roles []NsxRole `json:"roles,omitempty"`
}
