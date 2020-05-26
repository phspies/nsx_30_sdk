package nsxt

type FireWallServiceAssociationListResult struct {
	ServiceType string `json:"service_type"`
	// Firewall rule list result with pagination support.
	Results []FirewallRule `json:"results,omitempty"`
}
