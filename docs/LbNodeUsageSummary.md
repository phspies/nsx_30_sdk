# LbNodeUsageSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCounts** | [**[]LbNodeCountPerSeverity**](LbNodeCountPerSeverity.md) | The property identifies array of node count for each severity (RED, ORANGE and GREEN).  | [optional] [default to null]
**CurrentCreditNumber** | **int64** | The current credit number reflects the overall credit usage for all nodes.  | [optional] [default to null]
**NodeUsages** | [**[]LbNodeUsage**](LbNodeUsage.md) | The property contains lb node usages for each node.  | [optional] [default to null]
**Severity** | **string** | The severity calculation is based on current credit usage percentage of load balancer for all nodes.  | [optional] [default to null]
**RemainingPoolMembers** | **int64** | The overall remaining number of pool members which could be configured on all nodes.  | [optional] [default to null]
**CurrentPoolMembers** | **int64** | The overall number of pool members configured on all nodes.  | [optional] [default to null]
**UsagePercentage** | **float64** | The overall usage percentage of all nodes for load balancer. The value is the larger value between overall pool member usage percentage and overall load balancer credit usage percentage.  | [optional] [default to null]
**RemainingCreditNumber** | **int64** | The remaining credit number is the overall remaining credits that can be used for load balancer service configuration for all nodes.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

