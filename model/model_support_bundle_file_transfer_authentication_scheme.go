package nsxt

type SupportBundleFileTransferAuthenticationScheme struct {
	// User name to authenticate with
	Username string `json:"username"`
	// Authentication scheme name
	SchemeName string `json:"scheme_name"`
	// Password to authenticate with
	Password string `json:"password,omitempty"`
}
