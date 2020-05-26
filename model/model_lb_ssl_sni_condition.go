package nsxt

// This condition is used to match SSL SNI in client hello. This condition is only supported in TRANSPORT phase. Only if virtual server is associated with client SSL profile, matching SNI condition in transport phase is available to be configured. 
type LbSslSniCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// If true, case is significant when comparing SNI value. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
	// Determine how a specified string value is used to match SNI. 
	MatchType string `json:"match_type,omitempty"`
	// The SNI(Server Name indication) in client hello message. 
	Sni string `json:"sni"`
}
