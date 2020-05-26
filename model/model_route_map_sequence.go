package nsxt

type RouteMapSequence struct {
	SetCriteria *RouteMapSequenceSet `json:"set_criteria,omitempty"`
	// Action for the Sequence
	Action string `json:"action"`
	MatchCriteria *RouteMapSequenceMatch `json:"match_criteria"`
}
