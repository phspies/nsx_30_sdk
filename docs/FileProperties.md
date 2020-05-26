# FileProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**CreatedEpochMs** | **int64** | File creation time in epoch milliseconds | [default to null]
**ModifiedEpochMs** | **int64** | File modification time in epoch milliseconds | [default to null]
**Name** | **string** | File name | [default to null]
**Size** | **int64** | Size of the file in bytes | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

