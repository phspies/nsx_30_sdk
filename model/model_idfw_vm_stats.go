package nsxt

// Identity Firewall user login/session data for a single VM.
type IdfwVmStats struct {
	// Virtual machine (external ID or BIOS UUID) where login/logout event occurred.
	VmExtId string `json:"vm_ext_id"`
	// List of active (still logged in) user login/sessions data (no limit)
	ActiveSessions []IdfwUserSessionData `json:"active_sessions"`
	// Optional list of up to 5 most recent archived (previously logged in) user login/session data.
	ArchivedSessions []IdfwUserSessionData `json:"archived_sessions,omitempty"`
}
