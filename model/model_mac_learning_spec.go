package nsxt

// MAC learning configuration
type MacLearningSpec struct {
	// This property specifies the limit on the maximum number of MAC addresses that can be learned on a port. It is consumed by vswitch kernel module on the hypervisor while learning MACs per port for VMs that are local to the host. 
	Limit int32 `json:"limit,omitempty"`
	// The policy after MAC Limit is exceeded
	LimitPolicy string `json:"limit_policy,omitempty"`
	// This property specifies the limit on the maximum number of MACs learned for a remote Virtual Machine's MAC to vtep binding per overlay logical switch. 
	RemoteOverlayMacLimit int32 `json:"remote_overlay_mac_limit,omitempty"`
	// Aging time in sec for learned MAC address
	AgingTime int32 `json:"aging_time,omitempty"`
	// Allowing source MAC address learning
	Enabled bool `json:"enabled"`
	// Allowing flooding for unlearned MAC for ingress traffic
	UnicastFloodingAllowed bool `json:"unicast_flooding_allowed,omitempty"`
}
