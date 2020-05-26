package nsxt

// Identity Firewall user session data on a client machine (typically a VM). Multiple entries for the same user can be returned if the user logins to multiple sessions on the same VM. 
type IdfwUserSessionData struct {
	// AD user ID (may not exist).
	UserId string `json:"user_id,omitempty"`
	// User session ID.  This also indicates whether this is VDI / RDSH.
	UserSessionId int32 `json:"user_session_id"`
	// Virtual machine (external ID or BIOS UUID) where login/logout events occurred.
	VmExtId string `json:"vm_ext_id,omitempty"`
	// Identifier of user session data.
	Id string `json:"id,omitempty"`
	// Login time.
	LoginTime int64 `json:"login_time"`
	// AD user name.
	UserName string `json:"user_name"`
	// Logout time if applicable.  An active user session has no logout time. Non-active user session is stored (up to last 5 most recent entries) per VM and per user. 
	LogoutTime int64 `json:"logout_time,omitempty"`
	// AD Domain of user.
	DomainName string `json:"domain_name"`
}