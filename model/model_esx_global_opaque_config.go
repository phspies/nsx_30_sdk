package nsxt

// ESX global opaque configuration
type EsxGlobalOpaqueConfig struct {
	// Valid Global configuration types
	ResourceType string `json:"resource_type"`
	// A list of global opaque configuration for ESX hosts.
	OpaqueConfig []KeyValuePair `json:"opaque_config"`
}
