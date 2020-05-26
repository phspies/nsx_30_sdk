package nsxt

// usage of each capacity type ex. vm, cpu
type CapacityUsage struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// count of number of items of capacity_type
	UsageCount int64 `json:"usage_count,omitempty"`
	// type of the capacity field
	CapacityType string `json:"capacity_type,omitempty"`
}
