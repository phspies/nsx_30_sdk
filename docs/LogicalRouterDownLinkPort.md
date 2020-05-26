# LogicalRouterDownLinkPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalRouterId** | **string** | Identifier for logical router on which this port is created | [default to null]
**ServiceBindings** | [**[]ServiceBinding**](ServiceBinding.md) | Service Bindings | [optional] [default to null]
**ResourceType** | **string** | LogicalRouterUpLinkPort is allowed only on TIER0 logical router.   It is the north facing port of the logical router. LogicalRouterLinkPortOnTIER0 is allowed only on TIER0 logical router.   This is the port where the LogicalRouterLinkPortOnTIER1 of TIER1 logical router connects to. LogicalRouterLinkPortOnTIER1 is allowed only on TIER1 logical router.   This is the port using which the user connected to TIER1 logical router for upwards connectivity via TIER0 logical router.   Connect this port to the LogicalRouterLinkPortOnTIER0 of the TIER0 logical router. LogicalRouterDownLinkPort is for the connected subnets on the logical router. LogicalRouterLoopbackPort is a loopback port for logical router component   which is placed on chosen edge cluster member. LogicalRouterIPTunnelPort is a IPSec VPN tunnel port created on   logical router when route based VPN session configured. LogicalRouterCentralizedServicePort is allowed only on Active/Standby TIER0 and TIER1   logical router. Port can be connected to VLAN or overlay logical switch.   Unlike downlink port it does not participate in distributed routing and hosted   on all edge cluster members associated with logical router.   Stateful services can be applied on this port.  | [default to null]
**Subnets** | [**[]IpSubnet**](IPSubnet.md) | Logical router port subnets | [default to null]
**LinkedLogicalSwitchPortId** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**MacAddress** | **string** | MAC address | [optional] [default to null]
**NdraProfileId** | **string** | Identifier of Neighbor Discovery Router Advertisement profile associated with port. When NDRA profile id is associated at both the port level and logical router level, the profile id specified at port level takes the precedence.  | [optional] [default to null]
**RoutingPolicies** | [**[]RoutingPolicy**](RoutingPolicy.md) | Routing policies used to specify how the traffic, which matches the policy routes, will be processed.  | [optional] [default to null]
**EnableMulticast** | **bool** | If this flag is set to true - it will enable multicast on the downlink interface. If this flag is set to false - it will disable multicast on the downlink interface. This is supported only on Tier0 downlinks. Default value for Tier0 downlink will be true.  | [optional] [default to null]
**UrpfMode** | **string** | Unicast Reverse Path Forwarding mode | [optional] [default to URPF_MODE.STRICT]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

