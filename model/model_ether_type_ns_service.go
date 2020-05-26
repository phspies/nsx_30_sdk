package nsxt

// An NSService element that represents an ethertype protocol
type EtherTypeNsService struct {
	// The specific type of NSServiceElement
	ResourceType string `json:"resource_type"`
	// Type of the encapsulated protocol
	EtherType int64 `json:"ether_type"`
}
