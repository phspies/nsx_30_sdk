package nsxt

// Details about the destination NSX manager for the migration
type DestinationNsxApiEndpoint struct {
	// IP address or host name of the destination NSX manager to which the config will be migrated.
	DestinationNsxIp string `json:"destination_nsx_ip"`
	// Valid password for connecting to the destination NSX manager.
	DestinationNsxPassword string `json:"destination_nsx_password,omitempty"`
	// Destination NSX manager port that will be used to apply details.
	DestinationNsxPort int32 `json:"destination_nsx_port,omitempty"`
	// Valid username for connecting to the destination NSX manager.
	DestinationNsxUsername string `json:"destination_nsx_username,omitempty"`
}
