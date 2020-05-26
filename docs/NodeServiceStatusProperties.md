# NodeServiceStatusProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Reason** | **string** | Reason for service degradation | [optional] [default to null]
**Health** | **string** | Service health in addition to runtime_state | [optional] [default to null]
**MonitorPid** | **int64** | Service monitor process id | [optional] [default to null]
**Pids** | **[]int64** | Service process ids | [optional] [default to null]
**RuntimeState** | **string** | Service runtime state | [optional] [default to null]
**MonitorRuntimeState** | **string** | Service monitor runtime state | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

