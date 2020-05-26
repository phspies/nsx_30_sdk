package nsxt

// Information about mandatory access control
type MandatoryAccessControlProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// current status of Mandatory Access Control
	Status string `json:"status,omitempty"`
	// Enabled can be True/False
	Enabled bool `json:"enabled,omitempty"`
}
