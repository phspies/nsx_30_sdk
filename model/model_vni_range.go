package nsxt

// A range of virtual network identifiers (VNIs)
type VniRange struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Start value for vni range to be used for virtual networks
	Start int64 `json:"start"`
	// End value for vni range to be used for virtual networks
	End int64 `json:"end"`
}
