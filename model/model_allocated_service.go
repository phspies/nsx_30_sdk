package nsxt

type AllocatedService struct {
	ServiceReference *ResourceReference `json:"service_reference,omitempty"`
	// Represents the active or the standby state of the service.
	HighAvailabilityStatus string `json:"high_availability_status,omitempty"`
	// Additional properties of a service, say the sub_pool_size and sub_pool_type for a LoadBalancer. 
	AllocationDetails []KeyValuePair `json:"allocation_details,omitempty"`
}
