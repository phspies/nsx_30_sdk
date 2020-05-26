# DnsForwarderStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueriesForwarded** | **int64** | The total number of forwarded dns queries | [optional] [default to null]
**ConditionalForwarderStatistics** | [**[]PerForwarderStatistics**](PerForwarderStatistics.md) | The statistics of conditional forwarders | [optional] [default to null]
**DefaultForwarderStatistics** | [***PerForwarderStatistics**](PerForwarderStatistics.md) |  | [optional] [default to null]
**QueriesAnsweredLocally** | **int64** | The totocal number of queries answered from local cache | [optional] [default to null]
**UsedCacheStatistics** | [**[]PerNodeUsedCacheStatistics**](PerNodeUsedCacheStatistics.md) | The statistics of used cache | [optional] [default to null]
**ConfiguredCacheSize** | **int64** | The configured cache size, in kb | [optional] [default to null]
**Timestamp** | **int64** | Time stamp of the current statistics, in ms | [optional] [default to null]
**ErrorMessage** | **string** | Error message, if available | [optional] [default to null]
**TotalQueries** | **int64** | The total number of received dns queries | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

