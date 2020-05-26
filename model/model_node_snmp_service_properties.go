package nsxt

// Node SNMP service properties
type NodeSnmpServiceProperties struct {
	// Service name
	ServiceName string `json:"service_name"`
	ServiceProperties *SnmpServiceProperties `json:"service_properties"`
}
