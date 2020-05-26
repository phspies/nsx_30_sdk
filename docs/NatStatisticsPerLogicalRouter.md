# NatStatisticsPerLogicalRouter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated; unset if data source has never updated the data. | [optional] [default to null]
**PerTransportNodeStatistics** | [**[]NatStatisticsPerTransportNode**](NatStatisticsPerTransportNode.md) | Detailed per node statistics | [optional] [default to null]
**StatisticsAcrossAllNodes** | [***NatCounters**](NatCounters.md) |  | [optional] [default to null]
**LogicalRouterId** | **string** | Id for the logical router | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

