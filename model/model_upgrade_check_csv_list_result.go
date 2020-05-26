package nsxt

type UpgradeCheckCsvListResult struct {
	// File name set by HTTP server if API  returns CSV result as a file.
	FileName string `json:"file_name,omitempty"`
	Results []UpgradeCheckCsvRecord `json:"results,omitempty"`
}
