package nsxt

// A NSService that represents IPv4 or IPv6 ICMP protocol
type IcmpTypeNsService struct {
	// The specific type of NSServiceElement
	ResourceType string `json:"resource_type"`
	// ICMP message code
	IcmpCode int64 `json:"icmp_code,omitempty"`
	// ICMP message type
	IcmpType int64 `json:"icmp_type,omitempty"`
	// ICMP protocol type
	Protocol string `json:"protocol"`
}
