package nsxt

// NSX global configs for Routing
type RoutingGlobalConfig struct {
	// Valid Global configuration types
	ResourceType string `json:"resource_type"`
	// This is the global default MTU for all the logical uplinks in a NSX domain. Currently logical uplink MTU can only be set globally and applies to the entire NSX domain. There is no option to override this value at transport zone level or transport node level. If this value is not set, the default value of 1500 will be used.
	LogicalUplinkMtu int32 `json:"logical_uplink_mtu,omitempty"`
	// This is the global default MAC address for all VDRs in all transport nodes in a NSX system nested in another NSX system. It can be changed only when there is no transport node in the NSX system. All transport zones in such a nested NSX system will have the \"nested_nsx\" property being true so that all transport nodes will use this MAC for the VDR ports to avoid conflict with the VDR MAC in the outer NSX system.
	VdrMacNested string `json:"vdr_mac_nested,omitempty"`
	// This setting does not restrict configuration as per other modes. But the forwarding will only work as per the mode set here.
	L3ForwardingMode string `json:"l3_forwarding_mode"`
	// This is the global default MAC address for all VDRs in all transport nodes in a NSX system. It can be changed only when there is no transport node in the NSX system.
	VdrMac string `json:"vdr_mac,omitempty"`
}
