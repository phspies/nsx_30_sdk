package nsxt

// Node log properties
type NodeLogProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Last modified time expressed in milliseconds since epoch
	LastModifiedTime int64 `json:"last_modified_time,omitempty"`
	// Size of log file in bytes
	LogSize int64 `json:"log_size,omitempty"`
	// Name of log file
	LogName string `json:"log_name,omitempty"`
}
