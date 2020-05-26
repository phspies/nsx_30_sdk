package nsxt

// Health check performed by system automatically on a specific transport zone.  For overlay based zone, health check is performed on corresponding N-VDS of each transport node with the VLAN and MTU specified by uplink profile of N-VDS for the node.  For VLAN based zone, health check is performed on corresponding N-VDS of each  transport node with MTU specified by uplink profile of N-VDS for the node  and VLAN specified by all logical switches in this zone. 
type AutomaticHealthCheck struct {
	// ID of the transport zone where this automatic health check is performed. 
	TransportZoneId string `json:"transport_zone_id,omitempty"`
	Result *HealthCheckResult `json:"result,omitempty"`
}
