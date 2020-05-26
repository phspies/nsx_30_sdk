package nsxt

// Appliance registration access token
type RegistrationToken struct {
	// Access token
	Token string `json:"token,omitempty"`
	// User delegated by token
	User string `json:"user,omitempty"`
	// List results
	Roles []string `json:"roles"`
}
