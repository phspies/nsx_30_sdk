package nsxt

type AttachmentContext struct {
	// A flag to indicate whether to allocate addresses from allocation     pools bound to the parent logical switch. 
	AllocateAddresses string `json:"allocate_addresses,omitempty"`
	// Used to identify which concrete class it is
	ResourceType string `json:"resource_type"`
}
