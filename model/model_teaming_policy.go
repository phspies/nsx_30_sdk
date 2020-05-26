package nsxt

// Uplink Teaming Policy
type TeamingPolicy struct {
	// Teaming policy
	Policy string `json:"policy"`
	// List of Uplinks used in standby list
	StandbyList []Uplink `json:"standby_list,omitempty"`
	// List of Uplinks used in active list
	ActiveList []Uplink `json:"active_list"`
}
