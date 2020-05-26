package nsxt

// Runtime status information of the compute manager
type ComputeManagerStatus struct {
	// Version of the compute manager
	Version string `json:"version,omitempty"`
	// Status of connection with the compute manager
	ConnectionStatus string `json:"connection_status,omitempty"`
	// Errors when connecting with compute manager
	ConnectionErrors []ErrorInfo `json:"connection_errors,omitempty"`
	// If Compute manager is trusted as authorization server, then this Id will be Id of corresponding oidc end point. 
	OidcEndPointId string `json:"oidc_end_point_id,omitempty"`
	// Timestamp of the last successful update of Inventory, in epoch milliseconds.
	LastSyncTime int64 `json:"last_sync_time,omitempty"`
	// Details about connection status
	ConnectionStatusDetails string `json:"connection_status_details,omitempty"`
	// Errors when registering with compute manager
	RegistrationErrors []ErrorInfo `json:"registration_errors,omitempty"`
	// Registration status of compute manager
	RegistrationStatus string `json:"registration_status,omitempty"`
}
