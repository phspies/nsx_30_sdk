package nsxt

// The current global barrier number of the realized state
type CurrentRealizationStateBarrier struct {
	// Gives the current global barrier number for NSX
	CurrentBarrierNumber int64 `json:"current_barrier_number,omitempty"`
}
