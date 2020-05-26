package nsxt

type AcceptableComponentVersion struct {
	// List of component versions
	AcceptableVersions []string `json:"acceptable_versions"`
	// Node type
	ComponentType string `json:"component_type"`
}
