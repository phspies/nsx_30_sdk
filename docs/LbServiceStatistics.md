# LbServiceStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | [**[]LbPoolStatistics**](LbPoolStatistics.md) | Statistics of load balancer pools | [optional] [default to null]
**ServiceId** | **string** | load balancer service identifier | [default to null]
**VirtualServers** | [**[]LbVirtualServerStatistics**](LbVirtualServerStatistics.md) | Statistics of load balancer virtual servers | [optional] [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated | [optional] [default to null]
**Statistics** | [***LbServiceStatisticsCounter**](LbServiceStatisticsCounter.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

