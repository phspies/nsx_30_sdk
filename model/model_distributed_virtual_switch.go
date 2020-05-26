package nsxt

// DistributedVirtualSwitch on a VC
type DistributedVirtualSwitch struct {
	// ID of the virtual switch in compute manager
	CmLocalId string `json:"cm_local_id,omitempty"`
	// External id of the virtual switch
	ExternalId string `json:"external_id,omitempty"`
	// Switch type like VmwareDistributedVirtualSwitch
	OriginType string `json:"origin_type,omitempty"`
	// ID of the compute manager where this virtual switch is discovered. 
	OriginId string `json:"origin_id,omitempty"`
	// Array of discovered nodes connected to this switch.
	DiscoveredNodes []DiscoveredNode `json:"discovered_nodes,omitempty"`
	UplinkPortgroup *DistributedVirtualPortgroup `json:"uplink_portgroup,omitempty"`
	// UUID of the switch
	Uuid string `json:"uuid,omitempty"`
	// Key-Value map of additional properties of switch
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// It contains information about VMware specific multiple dynamic LACP groups. 
	LacpGroupConfigs []LacpGroupConfigInfo `json:"lacp_group_configs,omitempty"`
	// The uniform name of uplink ports on each host.
	UplinkPortNames []string `json:"uplink_port_names,omitempty"`
}
