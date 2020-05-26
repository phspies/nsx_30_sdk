package nsxt

type AgentStatusCount struct {
	// Roll-up agent status
	Status string `json:"status,omitempty"`
	// Down count
	DownCount int32 `json:"down_count,omitempty"`
	// List of agent statuses belonging to the transport node
	Agents []AgentStatus `json:"agents,omitempty"`
	// Up count
	UpCount int32 `json:"up_count,omitempty"`
}
