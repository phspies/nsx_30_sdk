package nsxt

type EffectiveMemberTypeListResult struct {
	// Collection of member types for the given NSGroup
	Results []string `json:"results"`
	// Count of the member types in the results array
	ResultCount int64 `json:"result_count,omitempty"`
}
