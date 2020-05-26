package nsxt

// CPU usage of DPDK and non-DPDK cores
type CpuUsage struct {
	// Indicates the average usage of all DPDK cores in percentage.
	AvgCpuCoreUsageDpdk float64 `json:"avg_cpu_core_usage_dpdk,omitempty"`
	// Indicates the highest cpu utilization value among non_dpdk cores in percentage.
	HighestCpuCoreUsageNonDpdk float64 `json:"highest_cpu_core_usage_non_dpdk,omitempty"`
	// Indicates the average usage of all non-DPDK cores in percentage.
	AvgCpuCoreUsageNonDpdk float64 `json:"avg_cpu_core_usage_non_dpdk,omitempty"`
	// Indicates the highest CPU utilization value among DPDK cores in percentage.
	HighestCpuCoreUsageDpdk float64 `json:"highest_cpu_core_usage_dpdk,omitempty"`
}
