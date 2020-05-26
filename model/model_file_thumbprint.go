package nsxt

// File thumbprint
type FileThumbprint struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// File's SHA256 thumbprint
	Sha256 string `json:"sha256"`
	// File name
	Name string `json:"name"`
	// File's SHA1 thumbprint
	Sha1 string `json:"sha1"`
}
