package nsxt

type AgentStatus struct {
	// Agent status
	Status string `json:"status,omitempty"`
	// Agent name
	Name string `json:"name,omitempty"`
}
