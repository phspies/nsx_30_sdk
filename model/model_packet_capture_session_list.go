package nsxt

type PacketCaptureSessionList struct {
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`
	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`
	// Total capture session count
	ResultCount int64 `json:"result_count"`
	// Packet capture list for all sessoins
	Results []PacketCaptureSession `json:"results,omitempty"`
}
