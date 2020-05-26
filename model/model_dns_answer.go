package nsxt

// The response for DNS nslookup. 
type DnsAnswer struct {
	// Dns server ip address and port, format is \"ip address#port\". 
	DnsServer string `json:"dns_server"`
	// It can be NXDOMAIN or error message which is not consisted of authoritative_answer or non_authoritative_answer. 
	RawAnswer string `json:"raw_answer,omitempty"`
	// Non-authotitative answers of the query. This is a deprecated property, please use 'answers' instead. 
	NonAuthoritativeAnswers []DnsQueryAnswer `json:"non_authoritative_answers,omitempty"`
	// The source ip used in this lookup. 
	SourceIp string `json:"source_ip"`
	// ID of the edge node that performed the query. 
	EdgeNodeId string `json:"edge_node_id"`
	// Authotitative answers of the query. This is a deprecated property, please use 'answers' instead. 
	AuthoritativeAnswers []DnsQueryAnswer `json:"authoritative_answers,omitempty"`
	// The answers of the query. 
	Answers []DnsQueryAnswer `json:"answers,omitempty"`
}
