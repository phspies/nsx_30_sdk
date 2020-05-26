package nsxt

// The Instance Deployment Config contains settings that is applied during install time.
type InstanceDeploymentConfig struct {
	// Context Id or VCenter Id.
	ContextId string `json:"context_id"`
	// List of NIC information for VMs
	VmNicInfos []VmNicInfo `json:"vm_nic_infos"`
	// Storage Id.
	StorageId string `json:"storage_id"`
	// The service VM will be deployed on the specified host in the specified server within the cluster if host_id is specified. Note: You must ensure that storage and specified networks are accessible by this host. 
	HostId string `json:"host_id,omitempty"`
	// Resource Pool or Compute Id.
	ComputeId string `json:"compute_id"`
}
