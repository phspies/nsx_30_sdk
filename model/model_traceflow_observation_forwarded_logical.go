package nsxt

type TraceflowObservationForwardedLogical struct {
	// Timestamp when the observation was created by the transport node (microseconds epoch)
	TimestampMicro int64 `json:"timestamp_micro,omitempty"`
	// The sub type of the component that issued the observation.
	ComponentSubType string `json:"component_sub_type,omitempty"`
	ResourceType string `json:"resource_type"`
	// The type of the component that issued the observation.
	ComponentType string `json:"component_type,omitempty"`
	// name of the transport node that observed a traceflow packet
	TransportNodeName string `json:"transport_node_name,omitempty"`
	// Timestamp when the observation was created by the transport node (milliseconds epoch)
	Timestamp int64 `json:"timestamp,omitempty"`
	// id of the transport node that observed a traceflow packet
	TransportNodeId string `json:"transport_node_id,omitempty"`
	// the hop count for observations on the transport node that a traceflow packet is injected in will be 0. The hop count is incremented each time a subsequent transport node receives the traceflow packet. The sequence number of 999 indicates that the hop count could not be determined for the containing observation.
	SequenceNo int64 `json:"sequence_no,omitempty"`
	// type of the transport node that observed a traceflow packet
	TransportNodeType string `json:"transport_node_type,omitempty"`
	// The name of the component that issued the observation.
	ComponentName string `json:"component_name,omitempty"`
	// The path index of the service insertion component
	ServicePathIndex int64 `json:"service_path_index,omitempty"`
	// The id of the component that forwarded the traceflow packet.
	ComponentId string `json:"component_id,omitempty"`
	// This field specified the VLAN id a traceflow packet matched in the whitelist in spoofguard.
	SpoofguardVlanId int64 `json:"spoofguard_vlan_id,omitempty"`
	// ARP_UNKNOWN_FROM_CP - Unknown ARP query result emitted by control plane ND_NS_UNKNOWN_FROM_CP - Unknown neighbor solicitation query result emitted by control plane UNKNOWN - Unknown resend type
	ResendType string `json:"resend_type,omitempty"`
	// The name of the logical port through which the traceflow packet was forwarded.
	LportName string `json:"lport_name,omitempty"`
	// The id of the acl rule that was applied to forward the traceflow packet
	AclRuleId int64 `json:"acl_rule_id,omitempty"`
	// The index of the service insertion component
	ServiceIndex int64 `json:"service_index,omitempty"`
	// VNI for the logical network on which the traceflow packet was forwarded.
	Vni int32 `json:"vni,omitempty"`
	// The name of the destination component to which the traceflow packet was forwarded.
	DstComponentName string `json:"dst_component_name,omitempty"`
	// The ID of the NAT rule that was applied to forward the traceflow packet
	NatRuleId int64 `json:"nat_rule_id,omitempty"`
	// The translated source IP address of VPN/NAT
	TranslatedSrcIp string `json:"translated_src_ip,omitempty"`
	// The translated destination IP address of VNP/NAT
	TranslatedDstIp string `json:"translated_dst_ip,omitempty"`
	// The source MAC address of form: \"^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$\". For example: 00:00:00:00:00:00. 
	SpoofguardMac string `json:"spoofguard_mac,omitempty"`
	// The type of the destination component to which the traceflow packet was forwarded.
	DstComponentType string `json:"dst_component_type,omitempty"`
	// The id of the logical port through which the traceflow packet was forwarded.
	LportId string `json:"lport_id,omitempty"`
	// The id of the destination component to which the traceflow packet was forwarded.
	DstComponentId string `json:"dst_component_id,omitempty"`
	// This field specified the prefix IP address a traceflow packet matched in the whitelist in spoofguard.
	SpoofguardIp string `json:"spoofguard_ip,omitempty"`
	// The ttl of the service insertion component
	ServiceTtl int64 `json:"service_ttl,omitempty"`
	// MAC address of nexthop for service insertion(SI) in service VM(SVM) where the traceflow packet was received. 
	SvcNhMac string `json:"svc_nh_mac,omitempty"`
}
