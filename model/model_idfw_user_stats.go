package nsxt

// Identity Firewall user login/session data for a single user.
type IdfwUserStats struct {
	// AD user Identifier (String ID)
	UserId string `json:"user_id,omitempty"`
	// List of active (still logged in) user login/sessions data (no limit)
	ActiveSessions []IdfwUserSessionData `json:"active_sessions"`
	// Optional list of up to 5 most recent archived (previously logged in) user login/session data. 
	ArchivedSessions []IdfwUserSessionData `json:"archived_sessions,omitempty"`
}
