package nsxt

type BackupUiFramesInfo struct {
	// prefix to be used for api call
	ApiEndpoint string `json:"api_endpoint,omitempty"`
	// Type of service, for which backup is handled
	FrameType string `json:"frame_type,omitempty"`
	// Version of the site
	SiteVersion string `json:"site_version,omitempty"`
	// Id of the site
	SiteId string `json:"site_id,omitempty"`
}
