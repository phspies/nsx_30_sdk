package nsxt

type IpfixServiceAssociationListResult struct {
	ServiceType string `json:"service_type"`
	// Ipfix config list result with pagination support.
	Results []IpfixConfig `json:"results,omitempty"`
}
