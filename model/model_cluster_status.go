package nsxt

type ClusterStatus struct {
	MgmtClusterStatus *ManagementClusterStatus `json:"mgmt_cluster_status,omitempty"`
	ControlClusterStatus *ControllerClusterStatus `json:"control_cluster_status,omitempty"`
	// Unique identifier of this cluster
	ClusterId string `json:"cluster_id,omitempty"`
	DetailedClusterStatus *AllClusterGroupStatus `json:"detailed_cluster_status,omitempty"`
}
