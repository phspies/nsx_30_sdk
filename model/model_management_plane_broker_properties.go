package nsxt

// Information about a management plane node this controller is configured to communicate with
type ManagementPlaneBrokerProperties struct {
	// IP address or hostname of the message bus broker on the management plane node.
	Host string `json:"host"`
	// Port number of the message bus broker on the management plane node.
	Port int64 `json:"port,omitempty"`
	// Certificate thumbprint of the message bus broker on the management plane node.
	Thumbprint string `json:"thumbprint"`
}
