package nsxt

// The list of the failed DNS queries with entry count and timestamp. The entry count is for per active/standby transport node. 
type PerNodeDnsFailedQueries struct {
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`
	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`
	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`
	// Timestamp of the request, in YYYY-MM-DD HH:MM:SS.zzz format.
	Timestamp string `json:"timestamp,omitempty"`
	// The Uuid of active/standby transport node.
	NodeId string `json:"node_id,omitempty"`
	// The list of failed DNS queries.
	Results []DnsFailedQuery `json:"results,omitempty"`
}
