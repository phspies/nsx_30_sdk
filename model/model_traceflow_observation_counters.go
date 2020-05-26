package nsxt

type TraceflowObservationCounters struct {
	// Total number of forwarded observations for this traceflow round.
	ForwardedCount int64 `json:"forwarded_count,omitempty"`
	// Total number of dropped observations for this round.
	DroppedCount int64 `json:"dropped_count,omitempty"`
	// Total number of delivered observations for this traceflow round.
	DeliveredCount int64 `json:"delivered_count,omitempty"`
	// Total number of received observations for this traceflow round.
	ReceivedCount int64 `json:"received_count,omitempty"`
}
