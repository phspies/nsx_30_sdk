package nsxt

// A set of IPv4 or IPv6 addresses defined by a start and end address.
type IpPoolRange struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// The start IP Address of the IP Range.
	Start string `json:"start"`
	// The end IP Address of the IP Range.
	End string `json:"end"`
}
