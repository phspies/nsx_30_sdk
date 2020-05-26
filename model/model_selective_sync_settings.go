package nsxt

// Directory domain selective sync settings
type SelectiveSyncSettings struct {
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`
	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`
	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`
	// Enable or disable SelectiveSync
	Enabled bool `json:"enabled"`
	// If SelectiveSync is enabled, this contains 1 or more OrgUnits, which NSX will synchronize with in LDAP server. The full distiguished name (DN) should be used for OrgUnit. If SelectiveSync is disabled, do not define this or specify an empty list. 
	SelectedOrgUnits []string `json:"selected_org_units,omitempty"`
}
