package nsxt

type BgpRouteAggregation struct {
	// Flag to send only summarized route
	SummaryOnly bool `json:"summary_only,omitempty"`
	// cidr of the aggregate address
	Prefix string `json:"prefix"`
}
