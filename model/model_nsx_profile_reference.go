package nsxt

// It is a reference to any NSX profile. It comprise of NSX profile type eg. DFWCPUProfile, CentralConfigProfile etc. and id of profile i.e. target_id 
type NsxProfileReference struct {
	// Display name of the NSX resource.
	TargetDisplayName string `json:"target_display_name,omitempty"`
	// Will be set to false if the referenced NSX resource has been deleted.
	IsValid bool `json:"is_valid,omitempty"`
	// Identifier of the NSX resource.
	TargetId string `json:"target_id,omitempty"`
	// Type of the NSX resource.
	TargetType string `json:"target_type,omitempty"`
	// Profile type of the ServiceConfig
	ProfileType string `json:"profile_type"`
}