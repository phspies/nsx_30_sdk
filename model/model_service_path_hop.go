package nsxt

// Represents a service VM implementing a particular service in a service chain
type ServicePathHop struct {
	// Indicating whether the corresponding service VM is active or not per DP.
	IsActiveFromDp bool `json:"is_active_from_dp,omitempty"`
	// Indicating whether the corresponding service VM is active or not per MP.
	IsActiveFromMp bool `json:"is_active_from_mp,omitempty"`
	// ID of the virtual network interface.
	Vif string `json:"vif,omitempty"`
	// MAC address of the virtual network interface.
	MacAddress string `json:"mac_address,omitempty"`
	// Action that will be taken by the corresponding service VM of the hop.
	Action string `json:"action,omitempty"`
	// Indicating whether the corresponding service VM is active or not per CCP.
	IsActiveFromCcp bool `json:"is_active_from_ccp,omitempty"`
	// Indicating the maintenance mode of the corresponding service VM.
	InMaintenanceMode bool `json:"in_maintenance_mode,omitempty"`
	// Indicating whether NSH liveness is supported or not by the corresponding service VM.
	NshLivenessSupport bool `json:"nsh_liveness_support,omitempty"`
	// Indicating whether service is configured to decrement SI field in NSH metadata.
	CanDecrementSi bool `json:"can_decrement_si,omitempty"`
}
