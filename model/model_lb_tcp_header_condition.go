package nsxt

// This condition is used to match TCP header fields of HTTP messages. Currently, only the TCP source port is supported. Ports can be expressed as a single port number like 80, or a port range like 1024-1030. 
type LbTcpHeaderCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// TCP source port of HTTP message
	SourcePort string `json:"source_port"`
}
