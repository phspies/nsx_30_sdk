package nsxt

// Preconfigured host switch is used for manually configured transport node.
type PreconfiguredHostSwitch struct {
	// External Id of the preconfigured host switch.
	HostSwitchId string `json:"host_switch_id"`
	// List of virtual tunnel endpoints which are preconfigured on this host switch
	Endpoints []PreconfiguredEndpoint `json:"endpoints,omitempty"`
	// List of TransportZones that are to be associated with specified host switch.
	TransportZoneEndpoints []TransportZoneEndPoint `json:"transport_zone_endpoints,omitempty"`
}
