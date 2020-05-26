package nsxt

// Node status properties
type NodeStatusProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// One, five, and fifteen minute load averages for the system
	LoadAverage []float64 `json:"load_average,omitempty"`
	// Amount of swap disk in use, in kilobytes
	SwapUsed int64 `json:"swap_used,omitempty"`
	CpuUsage *CpuUsage `json:"cpu_usage,omitempty"`
	// Number of non-DPDK cores on Edge Node.
	NonDpdkCpuCores int64 `json:"non_dpdk_cpu_cores,omitempty"`
	// Amount of disk available for swap, in kilobytes
	SwapTotal int64 `json:"swap_total,omitempty"`
	// Current time expressed in milliseconds since epoch
	SystemTime int64 `json:"system_time,omitempty"`
	// Number of CPU cores on the system
	CpuCores int64 `json:"cpu_cores,omitempty"`
	// Milliseconds since system start
	Uptime int64 `json:"uptime,omitempty"`
	// Amount of RAM on the system that can be flushed out to disk, in kilobytes
	MemCache int64 `json:"mem_cache,omitempty"`
	// Amount of RAM allocated to the system, in kilobytes
	MemTotal int64 `json:"mem_total,omitempty"`
	// Amount of RAM in use on the system, in kilobytes
	MemUsed int64 `json:"mem_used,omitempty"`
	// Number of DPDK cores on Edge Node which are used for packet IO processing.
	DpdkCpuCores int64 `json:"dpdk_cpu_cores,omitempty"`
	// File systems configured on the system
	FileSystems []NodeFileSystemProperties `json:"file_systems,omitempty"`
	// Source of status data.
	Source string `json:"source,omitempty"`
}
