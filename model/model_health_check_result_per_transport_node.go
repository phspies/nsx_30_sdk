package nsxt

// Health check result on specific transport node 
type HealthCheckResultPerTransportNode struct {
	ResultOnHostSwitch *HealthCheckResultOnHostSwitch `json:"result_on_host_switch,omitempty"`
	// ID of the Transport Node.
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
