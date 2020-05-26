package nsxt

type TraceflowObservationRelayedLogical struct {
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
	// This field specified the message type of the relay service REQUEST - The relay service will relay a request message to the destination server REPLY - The relay service will relay a reply message to the client
	MessageType string `json:"message_type,omitempty"`
	// This field specified the IP address of the destination which the packet will be relayed.
	DstServerAddress string `json:"dst_server_address,omitempty"`
	// This field specified the logical component that relay service located.
	LogicalCompUuid string `json:"logical_comp_uuid,omitempty"`
	// This field specified the IP address of the relay service.
	RelayServerAddress string `json:"relay_server_address,omitempty"`
}