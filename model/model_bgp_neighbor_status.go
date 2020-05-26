package nsxt

type BgpNeighborStatus struct {
	// Current state of the BGP session.
	ConnectionState string `json:"connection_state,omitempty"`
	// Count of messages received from the neighbor
	MessagesReceived int64 `json:"messages_received,omitempty"`
	// Time in ms to wait for HELLO packet from BGP peer
	KeepAliveInterval int64 `json:"keep_alive_interval,omitempty"`
	// Router ID of the BGP neighbor.
	NeighborRouterId string `json:"neighbor_router_id,omitempty"`
	// Sum of out prefixes counts across all address families.
	TotalOutPrefixCount int64 `json:"total_out_prefix_count,omitempty"`
	// Logical router component(Service Router/Distributed Router) id
	LrComponentId string `json:"lr_component_id,omitempty"`
	// Count of connections established
	EstablishedConnectionCount int64 `json:"established_connection_count,omitempty"`
	// Count of messages sent to the neighbor
	MessagesSent int64 `json:"messages_sent,omitempty"`
	// Time(in milliseconds) since connection was established.
	TimeSinceEstablished int64 `json:"time_since_established,omitempty"`
	// Time in ms to wait for HELLO from BGP peer. If a HELLO packet is not seen from BGP Peer withing hold_time then BGP neighbor will be marked as down.
	HoldTime int64 `json:"hold_time,omitempty"`
	// Current state of graceful restart where graceful_restart = true indicates graceful restart is enabled and graceful_restart = false indicates graceful restart is disabled. This is deprecated field, use graceful_restart_mode instead. 
	GracefulRestart bool `json:"graceful_restart,omitempty"`
	// Current state of graceful restart of BGP neighbor. Possible values are - 1. GR_AND_HELPER - Graceful restart with Helper 2. HELPER_ONLY - Helper only 3. DISABLE - Disabled 
	GracefulRestartMode string `json:"graceful_restart_mode,omitempty"`
	// Count of connection drop
	ConnectionDropCount int64 `json:"connection_drop_count,omitempty"`
	// TCP port number of remote BGP Connection
	RemotePort int64 `json:"remote_port,omitempty"`
	// Sum of in prefixes counts across all address families.
	TotalInPrefixCount int64 `json:"total_in_prefix_count,omitempty"`
	RemoteSite *ResourceReference `json:"remote_site,omitempty"`
	TransportNode *ResourceReference `json:"transport_node,omitempty"`
	// TCP port number of Local BGP connection
	LocalPort int64 `json:"local_port,omitempty"`
	// AS number of the BGP neighbor
	RemoteAsNumber string `json:"remote_as_number,omitempty"`
	// BGP capabilities sent to BGP neighbor.
	AnnouncedCapabilities []string `json:"announced_capabilities,omitempty"`
	// BGP capabilities negotiated with BGP neighbor.
	NegotiatedCapability []string `json:"negotiated_capability,omitempty"`
	// Address families of BGP neighbor
	AddressFamilies []BgpAddressFamily `json:"address_families,omitempty"`
	// The Ip address of logical port
	SourceAddress string `json:"source_address,omitempty"`
	// The IP of the BGP neighbor
	NeighborAddress string `json:"neighbor_address,omitempty"`
}
