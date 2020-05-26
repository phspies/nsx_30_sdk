package nsxt

type TraceflowObservationDroppedLogical struct {
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
	// The index of service path that is a chain of services represents the point where the traceflow packet was dropped. 
	ServicePathIndex int64 `json:"service_path_index,omitempty"`
	// The id of the component that dropped the traceflow packet.
	ComponentId string `json:"component_id,omitempty"`
}
