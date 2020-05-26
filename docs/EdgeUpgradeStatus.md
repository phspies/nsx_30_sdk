# EdgeUpgradeStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Upgrade status of component | [optional] [default to null]
**PreUpgradeStatus** | [***UpgradeChecksExecutionStatus**](UpgradeChecksExecutionStatus.md) |  | [optional] [default to null]
**Details** | **string** | Details about the upgrade status | [optional] [default to null]
**ComponentType** | **string** | Component type for the upgrade status | [optional] [default to null]
**NodeCountAtTargetVersion** | **int32** | Number of nodes of the type and at the component version | [optional] [default to null]
**TargetComponentVersion** | **string** | Target component version | [optional] [default to null]
**PercentComplete** | **float64** | Indicator of upgrade progress in percentage | [optional] [default to null]
**CanSkip** | **bool** | Can the upgrade of the remaining units in this component be skipped | [optional] [default to null]
**CurrentVersionNodeSummary** | [***NodeSummaryList**](NodeSummaryList.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

