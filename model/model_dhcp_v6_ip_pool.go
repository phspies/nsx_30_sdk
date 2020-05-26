package nsxt

// DHCP IPv6 ip pool to define dynamic ip allocation ranges. The DhcpV6IpPool would only provide stateless DHCP (domain search list, DNS servers, SNTP servers) to client if both the ranges and excluded_ranges are not specified. 
type DhcpV6IpPool struct {
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
	// Ip address ranges to define dynamic ip allocation ranges.
	Ranges []IpPoolRange `json:"ranges,omitempty"`
	// Excluded addresses to define dynamic ip allocation ranges.
	ExcludedRanges []IpPoolRange `json:"excluded_ranges,omitempty"`
}
