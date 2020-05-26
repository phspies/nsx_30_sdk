package nsxt

// DHCP IPv6 static binding to define a static ip allocation.
type DhcpV6StaticBinding struct {
	// Lease time, in seconds.
	LeaseTime int64 `json:"lease_time,omitempty"`
	// SNTP server ips.
	SntpServers []string `json:"sntp_servers,omitempty"`
	// Preferred time, in seconds. If this value is not provided, the value of lease_time*0.8 will be used. 
	PreferredTime int64 `json:"preferred_time,omitempty"`
	// Primary and secondary DNS server address to assign host. They can be overridden by ip-pool or static-binding level property. 
	DnsNameservers []string `json:"dns_nameservers,omitempty"`
	// Host name or prefix to be assigned to host. It can be overridden by ip-pool or static-binding level property. 
	DomainNames []string `json:"domain_names,omitempty"`
	// When not specified, no ip address will be assigned to client host.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// The MAC address of the host. Either client-duid or mac-address, but not both. 
	MacAddress string `json:"mac_address,omitempty"`
}
