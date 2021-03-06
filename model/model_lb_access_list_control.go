package nsxt

// LbAccessListControl is used to define how IP access list control can filter the connections from clients. 
type LbAccessListControl struct {
	// ALLOW means connections matching grouping object IP list are allowed and requests not matching grouping object IP list are dropped. DROP means connections matching grouping object IP list are dropped and requests not matching grouping object IP list are allowed. 
	Action string `json:"action"`
	// The identifier of grouping object which defines the IP addresses or ranges to match the client IP. 
	GroupId string `json:"group_id"`
	// The enabled flag indicates whether to enable access list control option. It is false by default. 
	Enabled bool `json:"enabled,omitempty"`
}
