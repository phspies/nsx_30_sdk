package nsxt

// Capability of a fabric node
type NodeCapability struct {
	// Value of this capability
	Value string `json:"value,omitempty"`
	// Version of the capability
	Version int32 `json:"version,omitempty"`
	// Description of this capability that can be displayed in UI
	Description string `json:"description,omitempty"`
	// String that identifies the base capability for all nodes
	Key string `json:"key,omitempty"`
	// Provider of this capability for the node
	Provider string `json:"provider,omitempty"`
}
