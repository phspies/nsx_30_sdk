package nsxt

type IpDiscoverySwitchingProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	ResourceType string `json:"resource_type"`
	// Indicates whether ARP snooping is enabled
	ArpSnoopingEnabled bool `json:"arp_snooping_enabled,omitempty"`
	// Indicates the number of arp snooped IP addresses to be remembered per LogicalPort. Decreasing this value, will retain the latest bindings from the existing list of address bindings. Increasing this value will retain existing bindings and also learn any new address bindings discovered on the port until the new limit is reached. This limit only applies to IPv4 addresses and is independent of the nd_bindings_limit used for IPv6 snooping.
	ArpBindingsLimit int32 `json:"arp_bindings_limit,omitempty"`
	// This option is the IPv6 equivalent of DHCP snooping.
	Dhcpv6SnoopingEnabled bool `json:"dhcpv6_snooping_enabled,omitempty"`
	// This option is the IPv6 equivalent of ARP snooping.
	NdSnoopingEnabled bool `json:"nd_snooping_enabled,omitempty"`
	// This option is only supported on ESX where vm-tools is installed.
	VmToolsV6Enabled bool `json:"vm_tools_v6_enabled,omitempty"`
	// Indicates whether DHCP snooping is enabled
	DhcpSnoopingEnabled bool `json:"dhcp_snooping_enabled,omitempty"`
	// This property controls the ARP and ND cache timeout period.It is recommended that this property be greater than the ARP/ND cache timeout on the VM. 
	ArpNdBindingTimeout int32 `json:"arp_nd_binding_timeout,omitempty"`
	// This option is only supported on ESX where vm-tools is installed.
	VmToolsEnabled bool `json:"vm_tools_enabled,omitempty"`
	// ARP snooping being inherently susceptible to ARP spoofing, uses a turst-on-fisrt-use (TOFU) paradigm where only the first IP address discovered via ARP snooping is trusted. The remaining are ignored. In order to allow for more flexibility, we allow the user to configure how many ARP snooped address bindings should be trusted for the lifetime of the logical port. This is controlled by the arp_bindings_limit property in the IP Discovery profile. We refer to this extension of TOFU as N-TOFU. However, if TOFU is disabled, then N ARP snooped IP addresses will be trusted until they are timed out, where N is configured by arp_bindings_limit. 
	TrustOnFirstUseEnabled bool `json:"trust_on_first_use_enabled,omitempty"`
	// Indicates the number of neighbor-discovery snooped IP addresses to be remembered per LogicalPort. Decreasing this value, will retain the latest bindings from the existing list of address bindings. Increasing this value will retain existing bindings and also learn any new address bindings discovered on the port until the new limit is reached. This limit only applies to IPv6 addresses and is independent of the arp_bindings_limit used for IPv4 snooping.
	NdBindingsLimit int32 `json:"nd_bindings_limit,omitempty"`
	DuplicateIpDetection *DuplicateIpDetection `json:"duplicate_ip_detection,omitempty"`
}
