package nsxt

type IpfixCollector struct {
	// IP address for the IPFIX collector
	CollectorIpAddress string `json:"collector_ip_address"`
	// Port for the IPFIX collector
	CollectorPort int32 `json:"collector_port,omitempty"`
}
