# TaskProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Status** | **string** | Current status of the task | [optional] [default to null]
**AsyncResponseAvailable** | **bool** | True if response for asynchronous request is available | [optional] [default to null]
**Description** | **string** | Description of the task | [optional] [default to null]
**StartTime** | **int64** | The start time of the task in epoch milliseconds | [optional] [default to null]
**Cancelable** | **bool** | True if this task can be canceled | [optional] [default to null]
**RequestMethod** | **string** | HTTP request method | [optional] [default to null]
**User** | **string** | Name of the user who created this task | [optional] [default to null]
**Progress** | **int64** | Task progress if known, from 0 to 100 | [optional] [default to null]
**Message** | **string** | A message describing the disposition of the task | [optional] [default to null]
**RequestUri** | **string** | URI of the method invocation that spawned this task | [optional] [default to null]
**Id** | **string** | Identifier for this task | [optional] [default to null]
**EndTime** | **int64** | The end time of the task in epoch milliseconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

