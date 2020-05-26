# LogicalRouterPortArpTable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Opaque cursor to be used for getting next page of records (supplied by current result page) | [optional] [default to null]
**SortAscending** | **bool** | If true, results are sorted in ascending order | [optional] [default to null]
**SortBy** | **string** | Field by which records are sorted | [optional] [default to null]
**ResultCount** | **int64** | Count of results found (across all pages), set only on first page | [optional] [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated; unset if data source has never updated the data. | [optional] [default to null]
**Results** | [**[]LogicalRouterPortArpEntry**](LogicalRouterPortArpEntry.md) |  | [optional] [default to null]
**LogicalRouterPortId** | **string** | The ID of the logical router port | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

