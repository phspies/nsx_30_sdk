package nsxt

// This port can be configured only on a TIER1 LogicalRouter. Use the id of the LogicalRouterLinkPortOnTIER0 from TIER0 logical router to set the linked_logical_router_port_id on the port on TIER1 logical router. 
type LogicalRouterLinkPortOnTier1 struct {
	// Identifier for logical router on which this port is created
	LogicalRouterId string `json:"logical_router_id"`
	// Service Bindings
	ServiceBindings []ServiceBinding `json:"service_bindings,omitempty"`
	// LogicalRouterUpLinkPort is allowed only on TIER0 logical router.   It is the north facing port of the logical router. LogicalRouterLinkPortOnTIER0 is allowed only on TIER0 logical router.   This is the port where the LogicalRouterLinkPortOnTIER1 of TIER1 logical router connects to. LogicalRouterLinkPortOnTIER1 is allowed only on TIER1 logical router.   This is the port using which the user connected to TIER1 logical router for upwards connectivity via TIER0 logical router.   Connect this port to the LogicalRouterLinkPortOnTIER0 of the TIER0 logical router. LogicalRouterDownLinkPort is for the connected subnets on the logical router. LogicalRouterLoopbackPort is a loopback port for logical router component   which is placed on chosen edge cluster member. LogicalRouterIPTunnelPort is a IPSec VPN tunnel port created on   logical router when route based VPN session configured. LogicalRouterCentralizedServicePort is allowed only on Active/Standby TIER0 and TIER1   logical router. Port can be connected to VLAN or overlay logical switch.   Unlike downlink port it does not participate in distributed routing and hosted   on all edge cluster members associated with logical router.   Stateful services can be applied on this port. 
	ResourceType string `json:"resource_type"`
	// Logical router port subnets
	Subnets []IpSubnet `json:"subnets,omitempty"`
	// MAC address
	MacAddress string `json:"mac_address,omitempty"`
	LinkedLogicalRouterPortId *ResourceReference `json:"linked_logical_router_port_id,omitempty"`
	// Please use logical router API to pass edge cluster members indexes manually. 
	EdgeClusterMemberIndex []int64 `json:"edge_cluster_member_index,omitempty"`
}
