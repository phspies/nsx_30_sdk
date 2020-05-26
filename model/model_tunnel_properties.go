package nsxt

type TunnelProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Status of tunnel
	Status string `json:"status,omitempty"`
	// Corresponds to the interface where local_ip_address is routed.
	EgressInterface string `json:"egress_interface,omitempty"`
	// UUID of the remote transport node
	RemoteNodeId string `json:"remote_node_id,omitempty"`
	Bfd *BfdProperties `json:"bfd,omitempty"`
	// Local IP address of tunnel
	LocalIp string `json:"local_ip,omitempty"`
	// Time at which the Tunnel status has been fetched last time.
	LastUpdatedTime int64 `json:"last_updated_time,omitempty"`
	// Name of tunnel
	Name string `json:"name,omitempty"`
	// Represents the display name of the remote transport node at the other end of the tunnel.
	RemoteNodeDisplayName string `json:"remote_node_display_name,omitempty"`
	// Tunnel encap
	Encap string `json:"encap,omitempty"`
	// Latency type.
	LatencyType string `json:"latency_type,omitempty"`
	// The latency value is set only when latency_type is VALID.
	LatencyValue int64 `json:"latency_value,omitempty"`
	// Remote IP address of tunnel
	RemoteIp string `json:"remote_ip,omitempty"`
}
