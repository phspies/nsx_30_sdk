# DnsFailedQueries

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | Timestamp of the request, in YYYY-MM-DD HH:MM:SS.zzz format.  | [optional] [default to null]
**PerNodeFailedQueries** | [**[]PerNodeDnsFailedQueries**](PerNodeDnsFailedQueries.md) | The array of failed DNS queries on active and standby transport node. If there is no standby node, the failed queries on standby node will not be present.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

