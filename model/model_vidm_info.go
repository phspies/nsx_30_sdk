package nsxt

// Vidm Info
type VidmInfo struct {
	// User's Full Name Or User Group's Display Name
	DisplayName string `json:"display_name,omitempty"`
	// Type
	Type_ string `json:"type,omitempty"`
	// Username Or Groupname
	Name string `json:"name,omitempty"`
}
