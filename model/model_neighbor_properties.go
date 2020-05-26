package nsxt

// Neighbor properties
type NeighborProperties struct {
	// System name
	SystemName string `json:"system_name,omitempty"`
	// System description
	SystemDesc string `json:"system_desc,omitempty"`
	// System port number
	SystemPortNumber int64 `json:"system_port_number,omitempty"`
	// Interface name
	Name string `json:"name,omitempty"`
	// Object identifier
	Oid string `json:"oid,omitempty"`
	// Management address
	MgmtAddr string `json:"mgmt_addr,omitempty"`
	// Capabilities
	Capabilities string `json:"capabilities,omitempty"`
	// True if currently in aggregation
	LinkAggregationStatus bool `json:"link_aggregation_status,omitempty"`
	// Interface index
	Ifindex int64 `json:"ifindex,omitempty"`
	// Interface MAC address
	Mac string `json:"mac,omitempty"`
	// Aggregation Capability
	LinkAggregationCapable bool `json:"link_aggregation_capable,omitempty"`
	// Port description
	PortDesc string `json:"port_desc,omitempty"`
	// Enabled capabilities
	EnabledCapabilities string `json:"enabled_capabilities,omitempty"`
	// Aggregation port id
	LinkAggregationPortId string `json:"link_aggregation_port_id,omitempty"`
}
