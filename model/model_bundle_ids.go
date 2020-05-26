package nsxt

// Contains a list of bundle-ids
type BundleIds struct {
	// Id of a bundle whose upload is successful
	Successful string `json:"successful,omitempty"`
	// Id of a bundle whose upload was failed
	Failed string `json:"failed,omitempty"`
	// Id of a bundle whose upload is in-progress
	InProgress string `json:"in_progress,omitempty"`
}
