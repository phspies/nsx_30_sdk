package nsxt

// Query statistics counters of a forwarder identified by domain names. 
type PerForwarderStatistics struct {
	// Statistics per upstream server.
	UpstreamStatistics []PerUpstreamServerStatistics `json:"upstream_statistics,omitempty"`
	// Domain names configured for the forwarder. Empty if this is the default forwarder. 
	DomainNames []string `json:"domain_names,omitempty"`
}
