package nsxt

type DiscoveredNode struct {
	// Timestamp of last modification
	LastSyncTime int64 `json:"_last_sync_time,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// The stateless property describes whether host persists its state across reboot or not. If state persists, value is set as false otherwise true.
	Stateless bool `json:"stateless,omitempty"`
	// External id of the compute collection to which this node belongs
	ParentComputeCollection string `json:"parent_compute_collection,omitempty"`
	// Certificate of the discovered node
	Certificate string `json:"certificate,omitempty"`
	// Id of the compute manager from where this node was discovered
	OriginId string `json:"origin_id,omitempty"`
	// IP Addresses of the the discovered node.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Hardware Id is generated using system hardware info. It is used to retrieve fabric node of the esx.
	HardwareId string `json:"hardware_id,omitempty"`
	// OS version of the discovered node
	OsVersion string `json:"os_version,omitempty"`
	// Discovered Node type like Host
	NodeType string `json:"node_type,omitempty"`
	// OS type of the discovered node
	OsType string `json:"os_type,omitempty"`
	// Key-Value map of additional specific properties of discovered node in the Compute Manager 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// External id of the discovered node, ex. a mo-ref from VC
	ExternalId string `json:"external_id,omitempty"`
	// Local Id of the discovered node in the Compute Manager
	CmLocalId string `json:"cm_local_id,omitempty"`
}
