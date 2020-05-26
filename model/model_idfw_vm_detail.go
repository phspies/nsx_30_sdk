package nsxt

// Identity Firewall user login/session data for a single VM
type IdfwVmDetail struct {
	// Virtual machine (external ID or BIOS UUID) where login/logout event occurred.
	VmExtId string `json:"vm_ext_id"`
	// List of client machine IP addresses.
	VmIpAddresses []string `json:"vm_ip_addresses,omitempty"`
	LastLoginUserSession *ResourceReference `json:"last_login_user_session,omitempty"`
	// List of user session data.
	UserSessions []IdfwUserSessionData `json:"user_sessions"`
}
