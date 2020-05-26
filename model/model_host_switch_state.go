package nsxt

// Host Switch State
type HostSwitchState struct {
	// VDS represents VMware vSphere Distributed Switch from vSphere that is used as HostSwitch through TransportNode or TransportNodeProfile configuration. When VDS is used as a HostSwitch, Hosts have to be added to VDS from vSphere and VDS instance is created on Hosts. To configure NSX on such hosts, you can use this VDS as a HostSwitch from NSX manager. vCenter has the ownership of MTU, LAG, NIOC and LLDP configuration of such VDS backed HostSwitch. Remaining configuration (e.g. UplinkHostswitchProfile) will be managed by NSX. NVDS represents NSX Virtual Switch which is NSX native HostSwitch. All configurations of NVDS will be managed by NSX.
	HostSwitchType string `json:"host_switch_type,omitempty"`
	// External ID of the HostSwitch
	HostSwitchId string `json:"host_switch_id,omitempty"`
	// List of virtual tunnel endpoints which are configured on this switch
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// List of Ids of TransportZones this HostSwitch belongs to
	TransportZoneIds []string `json:"transport_zone_ids,omitempty"`
	// The name must be unique among all host switches specified in a given Transport Node.
	HostSwitchName string `json:"host_switch_name,omitempty"`
}
