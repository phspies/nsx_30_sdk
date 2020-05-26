package nsxt

type BgpNeighborRouteDetailsInCsvFormat struct {
	// File name set by HTTP server if API  returns CSV result as a file.
	FileName string `json:"file_name,omitempty"`
	Results []BgpNeighborRouteDetailsCsvRecord `json:"results,omitempty"`
}
