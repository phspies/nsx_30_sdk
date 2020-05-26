package nsxt

type LacpGroupConfigInfo struct {
	// The key represents the identifier for the group that is unique across VC. 
	Key string `json:"key,omitempty"`
	// The display name of the LACP group.
	Name string `json:"name,omitempty"`
	// Keys for the uplink ports in the group. Each uplink port is assigned a key that is unique across VC. 
	UplinkPortKeys []string `json:"uplink_port_keys,omitempty"`
	// Load balance algorithm used in LACP group. The possible values are dictated by the values available in VC. Please refer VMwareDvsLacpLoadBalanceAlgorithm documentation for a full list of values. A few examples are srcDestIp where source and destination IP are considered, srcIp where only source IP is considered. 
	LoadBalanceAlgorithm string `json:"load_balance_algorithm,omitempty"`
	// The number of uplink ports
	UplinkNum int64 `json:"uplink_num,omitempty"`
	// Names for the uplink ports in the group.
	UplinkNames []string `json:"uplink_names,omitempty"`
	// The mode of LACP can be ACTIVE or PASSIVE. If the mode is ACTIVE, LACP is enabled unconditionally. If the mode is PASSIVE, LACP is enabled only if LACP device is detected. 
	Mode string `json:"mode,omitempty"`
}
