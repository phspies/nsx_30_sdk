package nsxt

// A profile holding TCP, UDP and ICMP session timeout configuration.
type FirewallSessionTimerProfile struct {
	// Resource type to use as profile type
	ResourceType string `json:"resource_type"`
	// The timeout value of connection in seconds after one endpoint sends an RST.
	TcpClosed int64 `json:"tcp_closed"`
	// The timeout value of connection in seconds after a second packet has been transferred.
	TcpOpening int64 `json:"tcp_opening"`
	// The timeout value of connection in seconds if the source host sends more than one packet but the destination host has never sent one back.
	UdpSingle int64 `json:"udp_single"`
	// The timeout value of connection in seconds after both FINs have been exchanged and connection is closed.
	TcpFinwait int64 `json:"tcp_finwait"`
	// The timeout value of connection in seconds after the first packet has been sent.
	TcpFirstPacket int64 `json:"tcp_first_packet"`
	// The timeout value of connection in seconds after the first FIN has been sent.
	TcpClosing int64 `json:"tcp_closing"`
	// The timeout value of connection in seconds once the connection has become fully established.
	TcpEstablished int64 `json:"tcp_established"`
	// The timeout value of connection in seconds if both hosts have sent packets.
	UdpMultiple int64 `json:"udp_multiple"`
	// The timeout value for the connection after an ICMP error came back in response to an ICMP packet.
	IcmpErrorReply int64 `json:"icmp_error_reply"`
	// The timeout value of connection in seconds after the first packet. This will be the initial timeout for the new UDP flow.
	UdpFirstPacket int64 `json:"udp_first_packet"`
	// The timeout value of connection in seconds after the first packet. This will be the initial timeout for the new ICMP flow.
	IcmpFirstPacket int64 `json:"icmp_first_packet"`
}
