package nsxt

// Node user properties
type NodeUserProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// User login name (must be \"root\" if userid is 0)
	Username string `json:"username,omitempty"`
	// Status of the user. This value can be ACTIVE indicating authentication attempts will be successful if the correct credentials are specified. The value can also be PASSWORD_EXPIRED indicating authentication attempts will fail because the user's password has expired and must be changed. Or, this value can be NOT_ACTIVATED indicating the user's password has not yet been set and must be set before the user can authenticate.
	Status string `json:"status,omitempty"`
	// Number of days since password was last changed
	LastPasswordChange int64 `json:"last_password_change,omitempty"`
	// Full name for the user
	FullName string `json:"full_name,omitempty"`
	// Number of days password is valid before it must be changed. This can be set to 0 to indicate no password change is required or a positive integer up to 9999. By default local user passwords must be changed every 90 days.
	PasswordChangeFrequency int64 `json:"password_change_frequency,omitempty"`
	// Password for the user (optionally specified on PUT, unspecified on GET)
	Password string `json:"password,omitempty"`
	// Numeric id for the user
	Userid int64 `json:"userid,omitempty"`
	// Old password for the user (required on PUT if password specified)
	OldPassword string `json:"old_password,omitempty"`
}