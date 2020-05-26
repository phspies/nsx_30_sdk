package nsxt

// DHCP server to support IPv4 DHCP service. Properties defined at DHCP server level can be overridden by ip-pool or static-binding level properties. 
type IPv4DhcpServer struct {
	Options *DhcpOptions `json:"options,omitempty"`
	// Enable or disable monitoring of DHCP ip-pools usage. When enabled, system events are generated when pool usage exceeds the configured thresholds. System events can be viewed in REST API /api/v2/hpm/alarms 
	MonitorIppoolUsage bool `json:"monitor_ippool_usage,omitempty"`
	// DHCP server ip in CIDR format.
	DhcpServerIp string `json:"dhcp_server_ip"`
	// Primary and secondary DNS server address to assign host. They can be overridden by ip-pool or static-binding level property. 
	DnsNameservers []string `json:"dns_nameservers,omitempty"`
	// Host name or prefix to be assigned to host. It can be overridden by ip-pool or static-binding level property. 
	DomainName string `json:"domain_name,omitempty"`
	// Gateway ip to be assigned to host. It can be overridden by ip-pool or static-binding level property. 
	GatewayIp string `json:"gateway_ip,omitempty"`
}