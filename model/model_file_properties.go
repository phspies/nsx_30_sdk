package nsxt

// File properties
type FileProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// File creation time in epoch milliseconds
	CreatedEpochMs int64 `json:"created_epoch_ms"`
	// File modification time in epoch milliseconds
	ModifiedEpochMs int64 `json:"modified_epoch_ms"`
	// File name
	Name string `json:"name"`
	// Size of the file in bytes
	Size int64 `json:"size"`
}
