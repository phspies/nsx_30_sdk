package nsxt

// Describes wether bundle upload is allowed or not
type BundleUploadPermission struct {
	// Reason for not allowing upload
	Reason string `json:"reason,omitempty"`
	// Flag indecation whether upload is allowed or not
	UploadAllowed bool `json:"upload_allowed,omitempty"`
}
