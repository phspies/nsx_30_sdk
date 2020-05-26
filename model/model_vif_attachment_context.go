package nsxt

type VifAttachmentContext struct {
	// A flag to indicate whether to allocate addresses from allocation     pools bound to the parent logical switch. 
	AllocateAddresses string `json:"allocate_addresses,omitempty"`
	// Used to identify which concrete class it is
	ResourceType string `json:"resource_type"`
	// Type of the VIF attached to logical port
	VifType string `json:"vif_type"`
	// VIF ID of the parent VIF if vif_type is CHILD
	ParentVifId string `json:"parent_vif_id,omitempty"`
	// An application ID used to identify / look up a child VIF behind a parent VIF. Only effective when vif_type is CHILD. 
	AppId string `json:"app_id,omitempty"`
	// Current we use VLAN id as the traffic tag. Only effective when vif_type is CHILD. Each logical port inside a container must have a unique traffic tag. If the traffic_tag is not unique, no error is generated, but traffic will not be delivered to any port with a non-unique tag. 
	TrafficTag int32 `json:"traffic_tag,omitempty"`
	// Only effective when vif_type is INDEPENDENT. Each logical port inside a bare metal server or container must have a transport node UUID. We use transport node ID as transport node UUID. 
	TransportNodeUuid string `json:"transport_node_uuid,omitempty"`
}
