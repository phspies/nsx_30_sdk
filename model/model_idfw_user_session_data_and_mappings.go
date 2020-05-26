package nsxt

// Identity Firewall user session data list and Directory Group to user mappings. 
type IdfwUserSessionDataAndMappings struct {
	// Archived user session data list
	ArchivedUserSessions []IdfwUserSessionData `json:"archived_user_sessions"`
	// Active user session data list
	ActiveUserSessions []IdfwUserSessionData `json:"active_user_sessions"`
	// Directory Group to user session data mappings
	DirGroupToUserSessionDataMappings []IdfwDirGroupUserSessionMapping `json:"dir_group_to_user_session_data_mappings"`
}
