package nsxt

type TraceflowObservationReceivedLogical struct {
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
	// MAC address of SAN volume controller for service insertion(SI) in service VM(SVM) where the traceflow packet was received. 
	SvcMac string `json:"svc_mac,omitempty"`
	// The id of the source component from which the traceflow packet was received.
	SrcComponentId string `json:"src_component_id,omitempty"`
	// The id of the component that received the traceflow packet.
	ComponentId string `json:"component_id,omitempty"`
	// The id of the logical port at which the traceflow packet was received
	LportId string `json:"lport_id,omitempty"`
	// The type of the source component from which the traceflow packet was received.
	SrcComponentType string `json:"src_component_type,omitempty"`
	// The name of the logical port at which the traceflow packet was received
	LportName string `json:"lport_name,omitempty"`
	// The name of source component from which the traceflow packet was received.
	SrcComponentName string `json:"src_component_name,omitempty"`
	// VNI for the logical network on which the traceflow packet was received.
	Vni int32 `json:"vni,omitempty"`
}
