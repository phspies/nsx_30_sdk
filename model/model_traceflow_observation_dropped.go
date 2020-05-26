package nsxt

type TraceflowObservationDropped struct {
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
	// The ID of the NAT rule that was applied to forward the traceflow packet
	NatRuleId int64 `json:"nat_rule_id,omitempty"`
	// The reason traceflow packet was dropped
	Reason string `json:"reason,omitempty"`
	// The id of the logical port at which the traceflow packet was dropped
	LportId string `json:"lport_id,omitempty"`
	// The name of the logical port at which the traceflow packet was dropped
	LportName string `json:"lport_name,omitempty"`
	// The id of the acl rule that was applied to drop the traceflow packet
	AclRuleId int64 `json:"acl_rule_id,omitempty"`
	// This field specifies the ARP fails reason ARP_TIMEOUT - ARP failure due to query control plane timeout ARP_CPFAIL - ARP failure due post ARP query message to control plane failure ARP_FROMCP - ARP failure due to deleting ARP entry from control plane ARP_PORTDESTROY - ARP failure due to port destruction ARP_TABLEDESTROY - ARP failure due to ARP table destruction ARP_NETDESTROY - ARP failure due to overlay network destruction
	ArpFailReason string `json:"arp_fail_reason,omitempty"`
}
