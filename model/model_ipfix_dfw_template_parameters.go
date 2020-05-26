package nsxt

// This entity represents the flow parameters which are exported. 
type IpfixDfwTemplateParameters struct {
	// Type of the IPv4 ICMP message. 
	SourceIcmpType bool `json:"source_icmp_type,omitempty"`
	// Code of the IPv4 ICMP message. 
	IcmpCode bool `json:"icmp_code,omitempty"`
	// The destination transport port of a monitored network flow. 
	DestinationTransportPort bool `json:"destination_transport_port,omitempty"`
	// The number of octets since the previous report (if any) in incoming packets for this flow at the observation point. The number of octets include IP header(s) and payload. 
	OctetDeltaCount bool `json:"octet_delta_count,omitempty"`
	// VIF UUID - enterprise specific Information Element that uniquely identifies VIF. 
	VifUuid bool `json:"vif_uuid,omitempty"`
	// The value of the protocol number in the IP packet header. 
	ProtocolIdentifier bool `json:"protocol_identifier,omitempty"`
	// Five valid values are allowed: 1. Flow Created. 2. Flow Deleted. 3. Flow Denied. 4. Flow Alert (not used in DropKick implementation). 5. Flow Update. 
	FirewallEvent bool `json:"firewall_event,omitempty"`
	// Two valid values are allowed: 1. 0x00: igress flow to VM. 2. 0x01: egress flow from VM. 
	FlowDirection bool `json:"flow_direction,omitempty"`
	// The absolute timestamp (seconds) of the last packet of this flow. 
	FlowEnd bool `json:"flow_end,omitempty"`
	// The source transport port of a monitored network flow. 
	SourceTransportPort bool `json:"source_transport_port,omitempty"`
	// The number of incoming packets since the previous report (if any) for this flow at the observation point. 
	PacketDeltaCount bool `json:"packet_delta_count,omitempty"`
	// The destination IP address of a monitored network flow. 
	DestinationAddress bool `json:"destination_address,omitempty"`
	// The source IP address of a monitored network flow. 
	SourceAddress bool `json:"source_address,omitempty"`
	// Firewall rule Id - enterprise specific Information Element that uniquely identifies firewall rule. 
	RuleId bool `json:"rule_id,omitempty"`
	// The absolute timestamp (seconds) of the first packet of this flow. 
	FlowStart bool `json:"flow_start,omitempty"`
}
