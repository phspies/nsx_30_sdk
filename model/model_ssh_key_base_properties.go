package nsxt

type SshKeyBaseProperties struct {
	// Current password for user (required for users root and admin)
	Password string `json:"password,omitempty"`
	// SSH key label (used to identify the key)
	Label string `json:"label"`
}
