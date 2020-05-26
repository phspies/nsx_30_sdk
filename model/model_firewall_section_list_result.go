package nsxt

type FirewallSectionListResult struct {
	// List of the firewall sections. The list has to be homogenous.
	Results []FirewallSection `json:"results,omitempty"`
}
