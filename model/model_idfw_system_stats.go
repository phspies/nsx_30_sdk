package nsxt

// Identity Firewall statistics data.
type IdfwSystemStats struct {
	// Number of concurrent logged on users (across VDI & RDSH).  Multiple logins by the same user is counted as 1. 
	NumConcurrentUsers int32 `json:"num_concurrent_users"`
	// Number of active user sessions/logins in IDFW enabled compute collections (including both UP and DOWN hosts).  N sessions/logins by the same user is counted as n. 
	NumUserSessions int32 `json:"num_user_sessions"`
}
