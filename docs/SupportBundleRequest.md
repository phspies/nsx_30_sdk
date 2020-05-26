# SupportBundleRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteFileServer** | [***SupportBundleRemoteFileServer**](SupportBundleRemoteFileServer.md) |  | [optional] [default to null]
**Nodes** | **[]string** | List of cluster/fabric node UUIDs processed in specified order | [default to null]
**ContentFilters** | **[]string** | Bundle should include content of specified type | [optional] [default to null]
**LogAgeLimit** | **int64** | Include log files with modified times not past the age limit in days | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

