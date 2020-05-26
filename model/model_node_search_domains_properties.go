package nsxt

// Node network search domains properties
type NodeSearchDomainsProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Search domains
	SearchDomains []string `json:"search_domains"`
}
