package nsxt

type AcceptableComponentVersionList struct {
	// Acceptable version whitelist for different components
	Results []AcceptableComponentVersion `json:"results"`
}
