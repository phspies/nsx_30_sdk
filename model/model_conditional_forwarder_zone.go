package nsxt

type ConditionalForwarderZone struct {
	// Ip address of the upstream DNS servers the DNS forwarder accesses. 
	UpstreamServers []string `json:"upstream_servers"`
	// The source ip used by the fowarder of the zone. If no source ip specified, the ip address of listener of the DNS forwarder will be used. 
	SourceIp string `json:"source_ip,omitempty"`
	// A forwarder domain name should be a valid FQDN. If reverse lookup is needed for this zone, reverse lookup domain name like X.in-addr.arpa can be defined. Here the X represents a subnet. 
	DomainNames []string `json:"domain_names"`
}
