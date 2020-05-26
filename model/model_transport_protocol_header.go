package nsxt

type TransportProtocolHeader struct {
	UdpHeader *UdpHeader `json:"udp_header,omitempty"`
	DhcpHeader *DhcpHeader `json:"dhcp_header,omitempty"`
	TcpHeader *TcpHeader `json:"tcp_header,omitempty"`
	IcmpEchoRequestHeader *IcmpEchoRequestHeader `json:"icmp_echo_request_header,omitempty"`
	Dhcpv6Header *Dhcpv6Header `json:"dhcpv6_header,omitempty"`
	NdpHeader *NdpHeader `json:"ndp_header,omitempty"`
	DnsHeader *DnsHeader `json:"dns_header,omitempty"`
}
