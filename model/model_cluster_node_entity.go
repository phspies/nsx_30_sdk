package nsxt

// NSX Cluster is made up of multiple cluster nodes. Each node can perform multiple functions, commonly referred to as roles. Cluster node entities are processes running in a cluster node that assist in the performance of a role. Cluster Boot Manager is a daemon that securely bootstraps and configures the entities. This type contains attributes of a cluster node entity that are relevant to the Cluster Boot Manager.
type ClusterNodeEntity struct {
	// Public certificate of the entity in PEM format
	Certificate string `json:"certificate"`
	// UUID of the entity
	EntityUuid string `json:"entity_uuid"`
	// Type of the entity
	EntityType string `json:"entity_type"`
	// Subnet mask prefix length of the entity binds to
	SubnetPrefixLength int64 `json:"subnet_prefix_length,omitempty"`
	// IP address the entity binds to
	IpAddress string `json:"ip_address,omitempty"`
	// Domain name the entity binds to
	Fqdn string `json:"fqdn,omitempty"`
	// Port the entity binds to
	Port int64 `json:"port,omitempty"`
}