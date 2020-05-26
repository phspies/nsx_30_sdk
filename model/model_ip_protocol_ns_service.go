package nsxt

// An NSService element that represents an IP protocol
type IpProtocolNsService struct {
	// The specific type of NSServiceElement
	ResourceType string `json:"resource_type"`
	// The IP protocol number
	ProtocolNumber int64 `json:"protocol_number"`
}
