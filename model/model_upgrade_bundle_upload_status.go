package nsxt

// Upload status of upgrade bundle uploaded from url
type UpgradeBundleUploadStatus struct {
	// URL for uploading upgrade bundle
	Url string `json:"url,omitempty"`
	// Current status of upgrade bundle upload
	Status string `json:"status,omitempty"`
	// Detailed status of upgrade bundle upload
	DetailedStatus string `json:"detailed_status,omitempty"`
	// Percent of bundle uploaded from URL
	Percent float64 `json:"percent,omitempty"`
}
