package nsxt

// A profile holding protection configuration for SYN flood, UDP flood, ICMP flood and other flood attack.
type FirewallFloodProtectionProfile struct {
	// Resource type to use as profile type
	ResourceType string `json:"resource_type"`
	// The maximum limit of active icmp connections. If this property is omitted, or set to null, then there is no limit on active icmp connections for those components if it's applied to ESX components (such as segment, segment port, virtual machine, etc); on the other side, if it's applied to EDGE components (such as, gateway), it will be set to default limit (10,000) on the specific components.
	IcmpActiveFlowLimit int64 `json:"icmp_active_flow_limit,omitempty"`
	// The maximum limit of other active connections besides udp, icmp and half open tcp connections. If this property is omitted, or set to null, then there is no limit on other active connections besides udp, icmp and tcp half open connections for those components if it's applied to ESX components (such as segment, segment port, virtual machine, etc); on the other side, if it's applied to EDGE components (such as, gateway), it will be set to default limit (10,000) on the specific components.
	OtherActiveConnLimit int64 `json:"other_active_conn_limit,omitempty"`
	// The flag to indicate syncache is enabled or not. This option does not apply to EDGE components.
	EnableSyncache bool `json:"enable_syncache,omitempty"`
	// The maximum limit of active udp connections. If this property is omitted, or set to null, then there is no limit on active udp connections for those components if it's applied to ESX components (such as segment, segment port, virtual machine, etc); on the other side, if it's applied to EDGE components (such as, gateway), it will be set to default limit (100,000) on the specific component.
	UdpActiveFlowLimit int64 `json:"udp_active_flow_limit,omitempty"`
	// The maximum limit of tcp half open connections. If this property is omitted, or set to null, then there is no limit on active tcp half open connections for those components if it's applied to ESX components (such as segment, segment port, virtual machine, etc); on the other side, if it's applied to EDGE components (such as, gateway), it will be set to default limit (1,000,000) on the specific components.
	TcpHalfOpenConnLimit int64 `json:"tcp_half_open_conn_limit,omitempty"`
	// The flag to indicate RST spoofing is enabled or not. This option does not apply to EDGE components. This can be enabled only if syncache is enabled.
	EnableRstSpoofing bool `json:"enable_rst_spoofing,omitempty"`
}
