# CurrentBackupOperationStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentStep** | **string** | Current step of operation | [optional] [default to null]
**BackupId** | **string** | Unique identifier of current backup | [optional] [default to null]
**CurrentStepMessage** | **string** | Additional human-readable status information about current step | [optional] [default to null]
**EndTime** | **int64** | Time when operation is expected to end | [optional] [default to null]
**OperationType** | **string** | Type of operation that is in progress. Returns none if no operation is in progress, in which case none of the other fields will be set.  | [default to null]
**StartTime** | **int64** | Time when operation was started | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

