package nsxt

type ControlConnStatus struct {
	// Status of the control Node for e.g  UP, DOWN.
	Status string `json:"status,omitempty"`
	// IP address of the control Node.
	ControlNodeIp string `json:"control_node_ip,omitempty"`
	// Failure status of the control Node for e.g CONNECTION_REFUSED,INCOMPLETE_HOST_CERT.
	FailureStatus string `json:"failure_status,omitempty"`
}
