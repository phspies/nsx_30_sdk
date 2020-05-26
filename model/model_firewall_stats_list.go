package nsxt

type FirewallStatsList struct {
	// Total count for firewall rule statistics in results set
	ResultCount int64 `json:"result_count,omitempty"`
	// Corresponding firewall section identifier for list of rule statistics
	SectionId string `json:"section_id,omitempty"`
	// List of rule statistics
	Results []FirewallStats `json:"results,omitempty"`
}
