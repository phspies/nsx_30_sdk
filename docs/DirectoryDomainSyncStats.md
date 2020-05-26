# DirectoryDomainSyncStats

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrevSyncStatus** | **string** | Directory domain previous sync status. It could be one of the following two states. | [optional] [default to null]
**AvgFullSyncTime** | **int64** | All the historical full sync are counted in calculating the average full sync time in milliseconds. | [optional] [default to null]
**PrevSyncType** | **string** | Directory domain previous sync type. It could be one of the following five states. Right after the directory domain is configured, this field is set to IDLE. | [optional] [default to null]
**NumFullSync** | **int64** | number of successful historical full sync initiated either by system or by API request. | [optional] [default to null]
**CurrentStateBeginTime** | **int64** | Since what time the current state has begun. The time is expressed in millisecond epoch time. | [optional] [default to null]
**AvgDeltaSyncTime** | **int64** | All the historical delta sync are counted in calculating the average delta sync time in milliseconds. | [optional] [default to null]
**PrevSyncError** | **string** | Directory domain previous sync status error if last status was failure. | [optional] [default to null]
**CurrentState** | **string** | Current running state of the directory domain in synchronization life cycle. It could be one of the following five states. SELECTIVE_FULL_SYNC and SELECTIVE_DELTA_SYNC are sync states for selective sync. | [optional] [default to null]
**NumDeltaSync** | **int64** | number of successful historical delta sync initiated either by system or by API request. | [optional] [default to null]
**PrevSyncEndTime** | **int64** | Directory domain previous sync ending time expressed in millisecond epoch time. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

