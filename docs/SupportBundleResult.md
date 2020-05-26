# SupportBundleResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestProperties** | [***SupportBundleRequest**](SupportBundleRequest.md) |  | [optional] [default to null]
**FailedNodes** | [**[]FailedNodeSupportBundleResult**](FailedNodeSupportBundleResult.md) | Nodes where bundles were not generated or not copied to remote server | [optional] [default to null]
**SuccessNodes** | [**[]SuccessNodeSupportBundleResult**](SuccessNodeSupportBundleResult.md) | Nodes whose bundles were successfully copied to remote file server | [optional] [default to null]
**RemainingNodes** | [**[]RemainingSupportBundleNode**](RemainingSupportBundleNode.md) | Nodes where bundle generation is pending or in progress | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

