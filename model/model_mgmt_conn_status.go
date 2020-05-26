package nsxt

type MgmtConnStatus struct {
	// Indicates the controller node's MP channel connectivity status
	ConnectivityStatus string `json:"connectivity_status,omitempty"`
}
