# LogicalRouterLoopbackPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalRouterId** | **string** | Identifier for logical router on which this port is created | [default to null]
**ServiceBindings** | [**[]ServiceBinding**](ServiceBinding.md) | Service Bindings | [optional] [default to null]
**ResourceType** | **string** | LogicalRouterUpLinkPort is allowed only on TIER0 logical router.   It is the north facing port of the logical router. LogicalRouterLinkPortOnTIER0 is allowed only on TIER0 logical router.   This is the port where the LogicalRouterLinkPortOnTIER1 of TIER1 logical router connects to. LogicalRouterLinkPortOnTIER1 is allowed only on TIER1 logical router.   This is the port using which the user connected to TIER1 logical router for upwards connectivity via TIER0 logical router.   Connect this port to the LogicalRouterLinkPortOnTIER0 of the TIER0 logical router. LogicalRouterDownLinkPort is for the connected subnets on the logical router. LogicalRouterLoopbackPort is a loopback port for logical router component   which is placed on chosen edge cluster member. LogicalRouterIPTunnelPort is a IPSec VPN tunnel port created on   logical router when route based VPN session configured. LogicalRouterCentralizedServicePort is allowed only on Active/Standby TIER0 and TIER1   logical router. Port can be connected to VLAN or overlay logical switch.   Unlike downlink port it does not participate in distributed routing and hosted   on all edge cluster members associated with logical router.   Stateful services can be applied on this port.  | [default to null]
**Subnets** | [**[]IpSubnet**](IPSubnet.md) | Loopback port subnets | [default to null]
**EdgeClusterMemberIndex** | **[]int64** | Member index of the edge node on the cluster | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

