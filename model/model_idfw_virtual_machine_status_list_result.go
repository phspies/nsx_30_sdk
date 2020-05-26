package nsxt

// VM ID and status of the Identity Firewall Compute collection. 
type IdfwVirtualMachineStatusListResult struct {
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`
	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`
	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`
	// List of VM's on Identity Firewall Compute collection.
	Results []IdfwVirtualMachineStatus `json:"results"`
}
