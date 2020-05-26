package nsxt

// Short name or alias of a url. It is used to represent the url.
type UrlAlias struct {
	// Url to fetch data from.
	Url string `json:"url"`
	// Short name or alias of url, if any. If not specified, the url can be referenced by its index in the array of urls of the datasource instance as $<index> (for example, $0).
	Alias string `json:"alias,omitempty"`
	// Search query to be applied, if any. If query string is not provided, it will be ignored.
	Query string `json:"query,omitempty"`
}
