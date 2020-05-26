package nsxt

// This condition is used to match IP header fields of HTTP messages. Either source_address or group_id should be specified. 
type LbIpHeaderCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// Source IP address of HTTP message should match IP addresses which are configured in Group in order to perform actions. 
	GroupId string `json:"group_id,omitempty"`
	// Source IP address of HTTP message. IP Address can be expressed as a single IP address like 10.1.1.1, or a range of IP addresses like 10.1.1.101-10.1.1.160. Both IPv4 and IPv6 addresses are supported. 
	SourceAddress string `json:"source_address,omitempty"`
}
