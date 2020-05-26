package nsxt

// Information about the upgrade bundle
type UpgradeBundleInfo struct {
	// URL for uploading upgrade bundle
	Url string `json:"url,omitempty"`
	// size of upgrade bundle
	BundleSize string `json:"bundle_size,omitempty"`
}
