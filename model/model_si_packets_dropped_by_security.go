package nsxt

type SiPacketsDroppedBySecurity struct {
	// The packets dropped by \"Spoof Guard\"; supported packet types are IPv4, IPv6, ARP, ND, non-IP.
	SpoofGuardDropped []SiPacketTypeAndCounter `json:"spoof_guard_dropped,omitempty"`
	// The number of IPv4 packets dropped by \"DHCP server block\".
	DhcpServerDroppedIpv4 int64 `json:"dhcp_server_dropped_ipv4,omitempty"`
	// The number of IPv6 packets dropped by \"DHCP server block\".
	DhcpServerDroppedIpv6 int64 `json:"dhcp_server_dropped_ipv6,omitempty"`
	// The number of IPv4 packets dropped by \"DHCP client block\".
	DhcpClientDroppedIpv4 int64 `json:"dhcp_client_dropped_ipv4,omitempty"`
	// The number of packets dropped by \"BPDU filter\".
	BpduFilterDropped int64 `json:"bpdu_filter_dropped,omitempty"`
	// The number of IPv6 packets dropped by \"DHCP client block\".
	DhcpClientDroppedIpv6 int64 `json:"dhcp_client_dropped_ipv6,omitempty"`
}
