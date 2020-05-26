package nsxt

// Details about source and destination NSX setup to be migrated
type MigrationSetupInfo struct {
	DestinationNsx *DestinationNsxApiEndpoint `json:"destination_nsx,omitempty"`
	// List of source NSX manager endpoints.
	SourceNsx []SourceNsxApiEndpoint `json:"source_nsx,omitempty"`
	// Migration mode can be VMC_V2T, ONPREMISE_V2T, ONPREMISE_VSPHERE2T
	MigrationMode string `json:"migration_mode,omitempty"`
}
