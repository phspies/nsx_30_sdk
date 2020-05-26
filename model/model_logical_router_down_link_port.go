package nsxt

type LogicalRouterDownLinkPort struct {
	// Identifier for logical router on which this port is created
	LogicalRouterId string `json:"logical_router_id"`
	// Service Bindings
	ServiceBindings []ServiceBinding `json:"service_bindings,omitempty"`
	// LogicalRouterUpLinkPort is allowed only on TIER0 logical router.   It is the north facing port of the logical router. LogicalRouterLinkPortOnTIER0 is allowed only on TIER0 logical router.   This is the port where the LogicalRouterLinkPortOnTIER1 of TIER1 logical router connects to. LogicalRouterLinkPortOnTIER1 is allowed only on TIER1 logical router.   This is the port using which the user connected to TIER1 logical router for upwards connectivity via TIER0 logical router.   Connect this port to the LogicalRouterLinkPortOnTIER0 of the TIER0 logical router. LogicalRouterDownLinkPort is for the connected subnets on the logical router. LogicalRouterLoopbackPort is a loopback port for logical router component   which is placed on chosen edge cluster member. LogicalRouterIPTunnelPort is a IPSec VPN tunnel port created on   logical router when route based VPN session configured. LogicalRouterCentralizedServicePort is allowed only on Active/Standby TIER0 and TIER1   logical router. Port can be connected to VLAN or overlay logical switch.   Unlike downlink port it does not participate in distributed routing and hosted   on all edge cluster members associated with logical router.   Stateful services can be applied on this port. 
	ResourceType string `json:"resource_type"`
	// Logical router port subnets
	Subnets []IpSubnet `json:"subnets"`
	LinkedLogicalSwitchPortId *ResourceReference `json:"linked_logical_switch_port_id,omitempty"`
	// MAC address
	MacAddress string `json:"mac_address,omitempty"`
	// Identifier of Neighbor Discovery Router Advertisement profile associated with port. When NDRA profile id is associated at both the port level and logical router level, the profile id specified at port level takes the precedence. 
	NdraProfileId string `json:"ndra_profile_id,omitempty"`
	// Routing policies used to specify how the traffic, which matches the policy routes, will be processed. 
	RoutingPolicies []RoutingPolicy `json:"routing_policies,omitempty"`
	// If this flag is set to true - it will enable multicast on the downlink interface. If this flag is set to false - it will disable multicast on the downlink interface. This is supported only on Tier0 downlinks. Default value for Tier0 downlink will be true. 
	EnableMulticast bool `json:"enable_multicast,omitempty"`
	// Unicast Reverse Path Forwarding mode
	UrpfMode string `json:"urpf_mode,omitempty"`
}
