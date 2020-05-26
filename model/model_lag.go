package nsxt

// LACP group
type Lag struct {
	// uplink names
	Uplinks []Uplink `json:"uplinks,omitempty"`
	// Lag name
	Name string `json:"name"`
	// number of uplinks
	NumberOfUplinks int32 `json:"number_of_uplinks"`
	// LACP timeout type
	TimeoutType string `json:"timeout_type,omitempty"`
	// LACP load balance Algorithm
	LoadBalanceAlgorithm string `json:"load_balance_algorithm"`
	// unique id
	Id string `json:"id,omitempty"`
	// LACP group mode
	Mode string `json:"mode"`
}
