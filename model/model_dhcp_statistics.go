package nsxt

type DhcpStatistics struct {
	// The total number of DHCP errors
	Errors int64 `json:"errors"`
	// The total number of DHCP RELEASE packets
	Releases int64 `json:"releases"`
	// The total number of DHCP INFORM packets
	Informs int64 `json:"informs"`
	// The total number of DHCP DECLINE packets
	Declines int64 `json:"declines"`
	// dhcp server uuid
	DhcpServerId string `json:"dhcp_server_id"`
	// The total number of DHCP NACK packets
	Nacks int64 `json:"nacks"`
	// The total number of DHCP OFFER packets
	Offers int64 `json:"offers"`
	// The total number of DHCP DISCOVER packets
	Discovers int64 `json:"discovers"`
	// The total number of DHCP ACK packets
	Acks int64 `json:"acks"`
	// timestamp of the statistics
	Timestamp int64 `json:"timestamp"`
	// The total number of DHCP REQUEST packets
	Requests int64 `json:"requests"`
	// The DHCP ip pool usage statistics
	IpPoolStats []DhcpIpPoolUsage `json:"ip_pool_stats,omitempty"`
}
