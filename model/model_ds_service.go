package nsxt

type DsService struct {
	// Display name of the NSX resource.
	TargetDisplayName string `json:"target_display_name,omitempty"`
	// Will be set to false if the referenced NSX resource has been deleted.
	IsValid bool `json:"is_valid,omitempty"`
	// Identifier of the NSX resource.
	TargetId string `json:"target_id,omitempty"`
	// Type of the NSX resource.
	TargetType string `json:"target_type,omitempty"`
	Service *NsServiceElement `json:"service,omitempty"`
}
