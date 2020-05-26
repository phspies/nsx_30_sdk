package nsxt

// List of all intrusions that are detected grouped by signature, it contains minimal details about the intrusions. 
type IdsEventsBySignatureResult struct {
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`
	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`
	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`
	// List of all intrusions detected, grouped by signature. The details include signature id, name, severity, timestamp, and total number of attempts per signature.
	Results []IdsEventsBySignature `json:"results,omitempty"`
}
