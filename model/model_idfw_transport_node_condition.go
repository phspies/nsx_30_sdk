package nsxt

// Status of the Identity Firewall Compute Collection's transport node. 
type IdfwTransportNodeCondition struct {
	// Transport node status for IDFW compute collection.
	Status string `json:"status"`
	// IDFW Compute collection's transport node condition.
	StatusDetail string `json:"status_detail,omitempty"`
}
