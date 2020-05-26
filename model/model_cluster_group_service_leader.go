package nsxt

// Each cluster node entity provides multiple services. When working in a group, each service can elect a cluster node entity to be the leader of the service. Leader election helps in coordination of the service. The leader holds a renewable lease on the leadership for a fixed period of time. The lease version is incremented every time the leadership lease is renewed. This type contains the attributes of a leader.
type ClusterGroupServiceLeader struct {
	// Name of the service
	ServiceName string `json:"service_name,omitempty"`
	// Number of times the lease has been renewed
	LeaseVersion int64 `json:"lease_version,omitempty"`
	// Member UUID of the leader
	LeaderUuid string `json:"leader_uuid,omitempty"`
}
