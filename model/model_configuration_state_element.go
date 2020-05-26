package nsxt

// Describes status of configuration of an entity
type ConfigurationStateElement struct {
	// URI of backing resource on sub system
	SubSystemAddress string `json:"sub_system_address,omitempty"`
	// State of configuration on this sub system
	State string `json:"state,omitempty"`
	// Error code
	FailureCode int64 `json:"failure_code,omitempty"`
	// Name of backing resource on sub system
	SubSystemName string `json:"sub_system_name,omitempty"`
	// Error message in case of failure
	FailureMessage string `json:"failure_message,omitempty"`
	// Type of backing resource on sub system
	SubSystemType string `json:"sub_system_type,omitempty"`
	// Identifier of backing resource on sub system
	SubSystemId string `json:"sub_system_id,omitempty"`
}
