# ApiError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleName** | **string** | The module name where the error occurred | [optional] [default to null]
**ErrorMessage** | **string** | A description of the error | [optional] [default to null]
**ErrorCode** | **int64** | A numeric error code | [optional] [default to null]
**Details** | **string** | Further details about the error | [optional] [default to null]
**ErrorData** | [***interface{}**](interface{}.md) | Additional data about the error | [optional] [default to null]
**RelatedErrors** | [**[]RelatedApiError**](RelatedApiError.md) | Other errors related to this error | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

