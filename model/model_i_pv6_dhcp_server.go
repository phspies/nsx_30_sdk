package nsxt

// DHCP server to support IPv6 DHCP service. Properties defined at DHCP server level can be overridden by ip-pool or static-binding level properties. 
type IPv6DhcpServer struct {
	// DHCP server ip in CIDR format.
	DhcpServerIp string `json:"dhcp_server_ip,omitempty"`
	// DHCP server id.
	ServerId string `json:"server_id,omitempty"`
	// Primary and secondary DNS server address to assign host. They can be overridden by ip-pool or static-binding level property. 
	DnsNameservers []string `json:"dns_nameservers,omitempty"`
	// SNTP server ips.
	SntpServers []string `json:"sntp_servers,omitempty"`
	// Host name or prefix to be assigned to host. It can be overridden by ip-pool or static-binding level property. 
	DomainNames []string `json:"domain_names,omitempty"`
}
