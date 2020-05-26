package nsxt

// Health Status of a third party partner VM.
type ServiceInstanceHealthStatus struct {
	// Protocol version might be different in both Mux and SVA.
	IsSvaMuxIncompatible bool `json:"is_sva_mux_incompatible,omitempty"`
	// Latest timestamp when mux was connected to SVA.
	ConnectTimestamp string `json:"connect_timestamp,omitempty"`
	// Mux version when Mux and SVA are incompatible
	MuxIncompatibleVersion string `json:"mux_incompatible_version,omitempty"`
	// Version of third party partner solution application.
	SolutionVersion string `json:"solution_version,omitempty"`
	// Latest timestamp when health status is received.
	SyncTime string `json:"sync_time,omitempty"`
	// Status of third party partner solution application.
	SolutionStatus string `json:"solution_status,omitempty"`
	// The parameter is set if the last received health status is older than the predefined interval. 
	IsStale bool `json:"is_stale,omitempty"`
	// Status of multiplexer which forwards the events from guest virtual machines to the partner appliance.
	MuxConnectedStatus string `json:"mux_connected_status,omitempty"`
}