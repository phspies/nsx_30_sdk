package nsxt

type SupportBundleRequest struct {
	RemoteFileServer *SupportBundleRemoteFileServer `json:"remote_file_server,omitempty"`
	// List of cluster/fabric node UUIDs processed in specified order
	Nodes []string `json:"nodes"`
	// Bundle should include content of specified type
	ContentFilters []string `json:"content_filters,omitempty"`
	// Include log files with modified times not past the age limit in days
	LogAgeLimit int64 `json:"log_age_limit,omitempty"`
}
