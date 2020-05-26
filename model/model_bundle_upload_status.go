package nsxt

// Upload status of bundle uploaded from local or remote location
type BundleUploadStatus struct {
	// URL for uploading bundle
	Url string `json:"url,omitempty"`
	// Name of the uploaded bundle.
	BundleName string `json:"bundle_name,omitempty"`
	// Detailed status of bundle upload
	DetailedStatus string `json:"detailed_status,omitempty"`
	// Percent of bundle uploaded from remote location
	Percent float64 `json:"percent,omitempty"`
	// Current status of bundle upload
	Status string `json:"status,omitempty"`
}
