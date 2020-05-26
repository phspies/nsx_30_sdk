package nsxt

// Software module details
type SoftwareModule struct {
	// Name of the module in the node
	ModuleName string `json:"module_name"`
	// Version of the module in the node
	ModuleVersion string `json:"module_version"`
}
