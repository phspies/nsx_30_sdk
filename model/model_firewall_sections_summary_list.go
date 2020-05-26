package nsxt

type FirewallSectionsSummaryList struct {
	// Timestamp of the last computation, in epoch milliseconds.
	LastComputeTime int64 `json:"last_compute_time,omitempty"`
	// List of firewall sections summary.
	SectionsSummary []FirewallSectionsSummary `json:"sections_summary,omitempty"`
}
