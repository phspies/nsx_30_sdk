# LbEdgeNodeUsage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** | The property identifies the node UUID for load balancer node usage.  | [default to null]
**Type_** | **string** | The property identifies the load balancer node usage type.  | [default to null]
**CurrentCreditNumber** | **int64** | The current credit number reflects the current usage on the node. For example, configuring a medium load balancer on a node consumes 10 credits. If there are 2 medium instances configured on a node, the current credit number is 2 * 10 &#x3D; 20.  | [optional] [default to null]
**FormFactor** | **string** | The form factor of the given edge node.  | [optional] [default to null]
**CurrentVirtualServers** | **int64** | The number of virtual servers configured on the node.  | [optional] [default to null]
**CurrentSmallLoadBalancerServices** | **int64** | The number of small load balancer services configured on the node.  | [optional] [default to null]
**CurrentPoolMembers** | **int64** | The number of pool members configured on the node.  | [optional] [default to null]
**Severity** | **string** | The severity calculation is based on current credit usage percentage of load balancer for one node.  | [optional] [default to null]
**CurrentPools** | **int64** | The number of pools configured on the node.  | [optional] [default to null]
**RemainingPoolMembers** | **int64** | The remaining number of pool members which could be configured on the given edge node.  | [optional] [default to null]
**EdgeClusterId** | **string** | The ID of edge cluster which contains the edge node.  | [optional] [default to null]
**RemainingXlargeLoadBalancerServices** | **int64** | The remaining number of xlarge load balancer services which could be configured on the given edge node.  | [optional] [default to null]
**RemainingSmallLoadBalancerServices** | **int64** | The remaining number of small load balancer services which could be configured on the given edge node.  | [optional] [default to null]
**CurrentXlargeLoadBalancerServices** | **int64** | The number of xlarge load balancer services configured on the node.  | [optional] [default to null]
**UsagePercentage** | **float64** | The usage percentage of the edge node for load balancer. The value is the larger value between load balancer credit usage percentage and pool member usage percentage for the edge node.  | [optional] [default to null]
**CurrentLargeLoadBalancerServices** | **int64** | The number of large load balancer services configured on the node.  | [optional] [default to null]
**RemainingCreditNumber** | **int64** | The remaining credit number is the remaining credits that can be used for load balancer service configuration. For example, an edge node with form factor LARGE_VIRTUAL_MACHINE has 40 credits, and a medium load balancer instance costs 10 credits. If there are currently 3 medium instances configured, the remaining credit number is 40 - (3 * 10) &#x3D; 10.  | [optional] [default to null]
**RemainingLargeLoadBalancerServices** | **int64** | The remaining number of large load balancer services which could be configured on the given edge node.  | [optional] [default to null]
**RemainingMediumLoadBalancerServices** | **int64** | The remaining number of medium load balancer services which could be configured on the given edge node.  | [optional] [default to null]
**CurrentMediumLoadBalancerServices** | **int64** | The number of medium load balancer services configured on the node.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

