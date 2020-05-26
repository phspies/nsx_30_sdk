package nsxt

type ForwarderZone struct {
	// Ip address of the upstream DNS servers the DNS forwarder accesses. 
	UpstreamServers []string `json:"upstream_servers"`
	// The source ip used by the fowarder of the zone. If no source ip specified, the ip address of listener of the DNS forwarder will be used. 
	SourceIp string `json:"source_ip,omitempty"`
}
