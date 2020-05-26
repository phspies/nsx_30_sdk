package nsxt

type RaDnsConfig struct {
	// Lifetime of Domain names in milliseconds
	DomainNameLifetime int64 `json:"domain_name_lifetime,omitempty"`
	// DNS server. 
	DnsServer []string `json:"dns_server,omitempty"`
	// Domain name in RA message. 
	DomainName []string `json:"domain_name,omitempty"`
	// Lifetime of DNS server in milliseconds
	DnsServerLifetime int64 `json:"dns_server_lifetime,omitempty"`
}
