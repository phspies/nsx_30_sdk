package nsxt

// The summary of the failed DNS query. The query result represents a full query chain from client VM to dns forwarder, and upstream server if no forwarder cache was hit. 
type DnsFailedQuery struct {
	// The time the query took before it got a failed answer, in ms.
	TimeSpent int64 `json:"time_spent,omitempty"`
	// The record type be queried, e.g. A, CNAME, SOA, etc.
	RecordType string `json:"record_type,omitempty"`
	// The client host ip address from which the query was issued. 
	ClientIp string `json:"client_ip,omitempty"`
	// The upstream server ip address to which the query was forwarded. If the query could not be serviced from the DNS forwarder cache, this property will contain the IP address of the DNS server that serviced the request. If the request was serviced from the cache, this property will be absent. 
	UpstreamServerIp string `json:"upstream_server_ip,omitempty"`
	// The detailed error message of the failed query, if any.
	ErrorMessage string `json:"error_message,omitempty"`
	// The address be queried, can be a FQDN or an ip address.
	Address string `json:"address,omitempty"`
	// Timestamp of the request, in YYYY-MM-DD HH:MM:SS.zzz format.
	Timestamp string `json:"timestamp"`
	// The type of the query failure, e.g. NXDOMAIN, etc.
	ErrorType string `json:"error_type,omitempty"`
	// The source ip address that is used to forward a query to an upstream server. 
	SourceIp string `json:"source_ip,omitempty"`
	// The DNS forwarder ip address to which the query was first received. 
	ForwarderIp string `json:"forwarder_ip,omitempty"`
}