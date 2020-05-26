package nsxt

// Role
type NsxRole struct {
	// This field represents the identifier of the role. With the introduction of custom roles, this field is no longer an enum.
	Role string `json:"role"`
	// Please use the /user-info/permissions api to get the permission that the user has on each feature.
	Permissions []string `json:"permissions,omitempty"`
}
