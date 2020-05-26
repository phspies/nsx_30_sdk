package nsxt

// The Deployment Config contains settings that are applied during install time.
type ServiceDeploymentConfig struct {
	// Moref of the datastore in VC. If it is to be taken from 'Agent VM Settings', then it should be empty.
	StorageId string `json:"storage_id,omitempty"`
	// The service VM will be deployed on the specified host in the specified server within the cluster if host_id is specified. Note: You must ensure that storage and specified networks are accessible       by this host. 
	HostId string `json:"host_id,omitempty"`
	// Resource Pool or cluster Id.
	ComputeCollectionId string `json:"compute_collection_id"`
	VmNicInfo *VmNicInfo `json:"vm_nic_info,omitempty"`
	// Context Id or VCenter Id.
	ComputeManagerId string `json:"compute_manager_id"`
}
