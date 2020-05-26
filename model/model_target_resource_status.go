package nsxt

// Holds status of target resource in firewall context.
type TargetResourceStatus struct {
	// Firewall status on a target logical resource.
	TargetStatus string `json:"target_status"`
	// Identifier of the NSX resource.
	TargetId string `json:"target_id,omitempty"`
}
