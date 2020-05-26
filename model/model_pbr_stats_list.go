package nsxt

type PbrStatsList struct {
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`
	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`
	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`
	// PBR section identifier.
	SectionId string `json:"section_id,omitempty"`
	// List of rule statistics.
	Results []PbrStats `json:"results,omitempty"`
}
