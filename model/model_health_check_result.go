package nsxt

// Result of health check.
type HealthCheckResult struct {
	// Overall status of VLAN-MTU health check result. 
	VlanMtuStatus string `json:"vlan_mtu_status,omitempty"`
	// List of health check results on specific transport node. 
	ResultsPerTransportNode []HealthCheckResultPerTransportNode `json:"results_per_transport_node,omitempty"`
	// Timestamp of check result updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`
}
