package nsxt

type PortMirroringFilter struct {
	// If set to MIRROR, packets will be mirrored. If set to DO_NOT_MIRROR, packets will not be mirrored.
	FilterAction string `json:"filter_action,omitempty"`
	// The transport protocols of TCP or UDP, used to match the transport protocol of a packet. If not provided, no filtering by IP protocols is performed.
	IpProtocol string `json:"ip_protocol,omitempty"`
	SrcIps *IpAddresses `json:"src_ips,omitempty"`
	DstIps *IpAddresses `json:"dst_ips,omitempty"`
	// Destination port in the form of a port or port range, used to match the destination port of a packet. If not provided, no filtering by destination port is performed.
	DstPorts string `json:"dst_ports,omitempty"`
	// Source port in the form of a port or port range, used to match the source port of a packet. If not provided, no filtering by source port is performed.
	SrcPorts string `json:"src_ports,omitempty"`
}
