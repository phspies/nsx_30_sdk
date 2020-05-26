package nsxt

// An instance of a datasource configuration.
type Datasource struct {
	// Name of a datasource instance.
	DisplayName string `json:"display_name"`
	// Array of urls relative to the datasource configuration. For example, api/v1/fabric/nodes is a relative url of nsx-manager instance.
	Urls []UrlAlias `json:"urls"`
}
