package nsxt

type FeatureUsageCsvRecord struct {
	// count of number of concurrent users
	CcuUsageCount int64 `json:"ccu_usage_count,omitempty"`
	// count of number of vms used by this feature
	VmUsageCount int64 `json:"vm_usage_count,omitempty"`
	// count of number of vcpus of public cloud VMs
	VcpuUsageCount int64 `json:"vcpu_usage_count,omitempty"`
	// count of number of cpu sockets used by this feature
	CpuUsageCount int64 `json:"cpu_usage_count,omitempty"`
	// name of the feature
	Feature string `json:"feature,omitempty"`
	// Number of CPU cores used by this feature
	CoreUsageCount int64 `json:"core_usage_count,omitempty"`
}
