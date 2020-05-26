# EdgeHighAvailabilityProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Supported cluster profiles. | [default to null]
**StandbyRelocationConfig** | [***StandbyRelocationConfig**](StandbyRelocationConfig.md) |  | [optional] [default to null]
**BfdDeclareDeadMultiple** | **int64** | Number of times a packet is missed before BFD declares the neighbor down. | [optional] [default to 3]
**BfdProbeInterval** | **int64** | the time interval (in millisec) between probe packets for heartbeat purpose | [optional] [default to 500]
**BfdAllowedHops** | **int64** | BFD allowed hops | [optional] [default to 255]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

