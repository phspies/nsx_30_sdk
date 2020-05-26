# LogicalSwitchStatusSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated; unset if data source has never updated the data. | [optional] [default to null]
**TotalSwitches** | **int64** | The total number of logical switches. | [default to null]
**Filters** | [**[]Filter**](Filter.md) | The filters used to find the logical switches- TransportZone id, LogicalSwitchProfile id or TransportType | [optional] [default to null]
**FullyRealizedSwitches** | **int64** | The number of logical switches that are realized in all transport nodes. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

