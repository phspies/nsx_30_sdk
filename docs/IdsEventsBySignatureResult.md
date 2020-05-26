# IdsEventsBySignatureResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Opaque cursor to be used for getting next page of records (supplied by current result page) | [optional] [default to null]
**SortAscending** | **bool** | If true, results are sorted in ascending order | [optional] [default to null]
**SortBy** | **string** | Field by which records are sorted | [optional] [default to null]
**ResultCount** | **int64** | Count of results found (across all pages), set only on first page | [optional] [default to null]
**Results** | [**[]IdsEventsBySignature**](IDSEventsBySignature.md) | List of all intrusions detected, grouped by signature. The details include signature id, name, severity, timestamp, and total number of attempts per signature. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

