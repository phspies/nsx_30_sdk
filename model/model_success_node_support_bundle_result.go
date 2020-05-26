package nsxt

type SuccessNodeSupportBundleResult struct {
	// Name of support bundle, e.g. nsx_NODETYPE_UUID_YYYYMMDD_HHMMSS.tgz
	BundleName string `json:"bundle_name,omitempty"`
	// Display name of node
	NodeDisplayName string `json:"node_display_name,omitempty"`
	// UUID of node
	NodeId string `json:"node_id,omitempty"`
	// File's SHA256 thumbprint
	Sha256Thumbprint string `json:"sha256_thumbprint,omitempty"`
	// Size of support bundle in bytes
	BundleSize int64 `json:"bundle_size,omitempty"`
}
