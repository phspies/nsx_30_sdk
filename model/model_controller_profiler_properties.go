package nsxt

// Controller profiler properties
type ControllerProfilerProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// True for enabling controller profiler, False for disabling controller profiler. 
	Enabled bool `json:"enabled,omitempty"`
}
