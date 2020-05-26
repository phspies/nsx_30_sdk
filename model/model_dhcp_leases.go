package nsxt

type DhcpLeases struct {
	// timestamp of the lease info
	Timestamp int64 `json:"timestamp,omitempty"`
	// The lease info list of the server
	Leases []DhcpLeasePerIp `json:"leases,omitempty"`
	// dhcp server uuid
	DhcpServerId string `json:"dhcp_server_id,omitempty"`
}
