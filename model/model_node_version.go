package nsxt

type NodeVersion struct {
	// Product version
	ProductVersion string `json:"product_version,omitempty"`
	// Node version
	NodeVersion string `json:"node_version,omitempty"`
}
