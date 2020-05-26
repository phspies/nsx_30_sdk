package nsxt

// Upgrade Coordinator state properties
type UcStateProperties struct {
	// Flag for updating upgrade-coodinator state properties to database
	UpdateUcStateProperties bool `json:"update_uc_state_properties,omitempty"`
}
