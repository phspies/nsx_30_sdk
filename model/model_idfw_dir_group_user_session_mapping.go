package nsxt

// Identity Firewall NSGorup to user mapping to link DirGroup to user session data. 
type IdfwDirGroupUserSessionMapping struct {
	// User ID.
	UserId string `json:"user_id,omitempty"`
	// Directory Group ID.
	DirGroupId string `json:"dir_group_id,omitempty"`
}
