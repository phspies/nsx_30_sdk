# PerNodeDnsFailedQueries

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Opaque cursor to be used for getting next page of records (supplied by current result page) | [optional] [default to null]
**SortAscending** | **bool** | If true, results are sorted in ascending order | [optional] [default to null]
**SortBy** | **string** | Field by which records are sorted | [optional] [default to null]
**ResultCount** | **int64** | Count of results found (across all pages), set only on first page | [optional] [default to null]
**Timestamp** | **string** | Timestamp of the request, in YYYY-MM-DD HH:MM:SS.zzz format. | [optional] [default to null]
**NodeId** | **string** | The Uuid of active/standby transport node. | [optional] [default to null]
**Results** | [**[]DnsFailedQuery**](DnsFailedQuery.md) | The list of failed DNS queries. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

