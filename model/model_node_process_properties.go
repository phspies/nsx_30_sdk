package nsxt

// Node process properties
type NodeProcessProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Virtual memory used by process in bytes
	MemUsed int64 `json:"mem_used,omitempty"`
	// CPU time (user and system) consumed by process in milliseconds
	CpuTime int64 `json:"cpu_time,omitempty"`
	// Parent process id
	Ppid int64 `json:"ppid,omitempty"`
	// Process start time expressed in milliseconds since epoch
	StartTime int64 `json:"start_time,omitempty"`
	// Process name
	ProcessName string `json:"process_name,omitempty"`
	// Process id
	Pid int64 `json:"pid,omitempty"`
	// Milliseconds since process started
	Uptime int64 `json:"uptime,omitempty"`
	// Resident set size of process in bytes
	MemResident int64 `json:"mem_resident,omitempty"`
}
