# IdsEventFlowData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationPort** | **int64** | Port on the destination VM where the traffic was sent to. | [optional] [default to null]
**ClientIp** | **string** | IP address of the VM that initiated the communication. | [optional] [default to null]
**Protocol** | **string** | Traffic protocol pertaining to the detected intrusion, could be TCP/UDP etc. | [optional] [default to null]
**SourcePort** | **int64** | Source port through which traffic was initiated that caused the intrusion to be detected. | [optional] [default to null]
**LocalVmIp** | **string** | IP address of VM on the host where IDS engine is running. | [optional] [default to null]
**SourceIp** | **string** | IP address of the source VM on the intrusion flow. | [optional] [default to null]
**DestinationIp** | **string** | IP address of the destination VM on the intrusion flow. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

