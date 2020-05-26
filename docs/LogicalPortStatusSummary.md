# LogicalPortStatusSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPorts** | **int64** | The total number of logical ports. | [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated; unset if data source has never updated the data. | [optional] [default to null]
**UpPorts** | **int64** | The number of logical ports whose Operational status is UP | [default to null]
**Filters** | [**[]Filter**](Filter.md) | The filters used to find the logical ports- TransportZone id, LogicalSwitch id or LogicalSwitchProfile id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

