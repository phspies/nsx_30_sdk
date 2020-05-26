package nsxt

// Audit log collection results
type AuditLogListResult struct {
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`
	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`
	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`
	// Timestamp of the last full audit log collection
	LastFullSyncTimestamp string `json:"last_full_sync_timestamp"`
	// Audit log results
	Results []AuditLog `json:"results"`
}
