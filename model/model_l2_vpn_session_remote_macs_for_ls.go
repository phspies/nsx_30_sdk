package nsxt

// Remote MAC addresses for logical switch.
type L2VpnSessionRemoteMacsForLs struct {
	// Mac addresses.
	RemoteMacAddresses []string `json:"remote_mac_addresses,omitempty"`
	LogicalSwitch *ResourceReference `json:"logical_switch,omitempty"`
	// Contains policy specific information like policy path.
	Tags []Tag `json:"tags,omitempty"`
}
