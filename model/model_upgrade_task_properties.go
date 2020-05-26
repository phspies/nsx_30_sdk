package nsxt

// Task properties
type UpgradeTaskProperties struct {
	// Name of Bundle
	BundleName string `json:"bundle_name"`
	// Step name
	Step string `json:"step,omitempty"`
	// Bundle arguments
	Parameters *interface{} `json:"parameters,omitempty"`
}
