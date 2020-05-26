package nsxt

// This port can be configured on both TIER0 and TIER1 logical router. Port can be connected to VLAN or overlay logical switch. Unlike downlink port it does not participate in distributed routing and hosted on all edge cluster members associated with logical router. Stateful services can be applied on this port. 
type LogicalRouterCentralizedServicePort struct {
	// Identifier for logical router on which this port is created
	LogicalRouterId string `json:"logical_router_id"`
	// Service Bindings
	ServiceBindings []ServiceBinding `json:"service_bindings,omitempty"`
	// LogicalRouterUpLinkPort is allowed only on TIER0 logical router.   It is the north facing port of the logical router. LogicalRouterLinkPortOnTIER0 is allowed only on TIER0 logical router.   This is the port where the LogicalRouterLinkPortOnTIER1 of TIER1 logical router connects to. LogicalRouterLinkPortOnTIER1 is allowed only on TIER1 logical router.   This is the port using which the user connected to TIER1 logical router for upwards connectivity via TIER0 logical router.   Connect this port to the LogicalRouterLinkPortOnTIER0 of the TIER0 logical router. LogicalRouterDownLinkPort is for the connected subnets on the logical router. LogicalRouterLoopbackPort is a loopback port for logical router component   which is placed on chosen edge cluster member. LogicalRouterIPTunnelPort is a IPSec VPN tunnel port created on   logical router when route based VPN session configured. LogicalRouterCentralizedServicePort is allowed only on Active/Standby TIER0 and TIER1   logical router. Port can be connected to VLAN or overlay logical switch.   Unlike downlink port it does not participate in distributed routing and hosted   on all edge cluster members associated with logical router.   Stateful services can be applied on this port. 
	ResourceType string `json:"resource_type"`
	LinkedLogicalSwitchPortId *ResourceReference `json:"linked_logical_switch_port_id,omitempty"`
	// Logical router port subnets
	Subnets []IpSubnet `json:"subnets,omitempty"`
	// Port is exclusively used for N-S service insertion
	EnableNetx bool `json:"enable_netx,omitempty"`
	// Unicast Reverse Path Forwarding mode
	UrpfMode string `json:"urpf_mode,omitempty"`
	// Identifier of Neighbor Discovery Router Advertisement profile associated with port. When NDRA profile id is associated at both the port level and logical router level, the profile id specified at port level takes the precedence. 
	NdraProfileId string `json:"ndra_profile_id,omitempty"`
	// Maximum transmission unit specifies the size of the largest packet that a network protocol can transmit. If not specified, the global logical MTU set in the /api/v1/global-configs/RoutingGlobalConfig API will be used. 
	Mtu int64 `json:"mtu,omitempty"`
}
