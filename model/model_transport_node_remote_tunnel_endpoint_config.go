package nsxt

// Remote tunnel endpoint configuration
type TransportNodeRemoteTunnelEndpointConfig struct {
	// Specifying this field will override the default teaming policy of the host switch and will be used by remote tunnel endpoint traffic.
	NamedTeamingPolicy string `json:"named_teaming_policy,omitempty"`
	// The host switch name should reference an existing host switch specified in the transport node configuration. The name will be used to identify the host switch responsible for processing remote tunnel endpoint traffic.
	HostSwitchName string `json:"host_switch_name"`
	// The transport VLAN id used for tagging intersite overlay traffic between remote tunnel endpoints.
	RtepVlan int64 `json:"rtep_vlan"`
	IpAssignmentSpec *IpAssignmentSpec `json:"ip_assignment_spec"`
}
