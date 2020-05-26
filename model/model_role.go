package nsxt

// Role
type Role struct {
	// Short identifier for the role. Must be all lower case with no spaces.
	Role string `json:"role"`
	// A short, human-friendly display name of the role.
	RoleDisplayName string `json:"role_display_name,omitempty"`
}
