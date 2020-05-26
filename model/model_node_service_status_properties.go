package nsxt

// Node service status properties
type NodeServiceStatusProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Reason for service degradation
	Reason string `json:"reason,omitempty"`
	// Service health in addition to runtime_state
	Health string `json:"health,omitempty"`
	// Service monitor process id
	MonitorPid int64 `json:"monitor_pid,omitempty"`
	// Service process ids
	Pids []int64 `json:"pids,omitempty"`
	// Service runtime state
	RuntimeState string `json:"runtime_state,omitempty"`
	// Service monitor runtime state
	MonitorRuntimeState string `json:"monitor_runtime_state,omitempty"`
}
