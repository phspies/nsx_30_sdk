# MigrationFeedbackRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolved** | **bool** | Indicates if a valid response already exist for this feedback request. | [optional] [default to null]
**AcceptedActions** | **[]string** | List of acceptable actions for this feedback request. | [optional] [default to null]
**Hash** | **string** | Identify a feedback request type across objects. This can be used to group together objects with similar feedback request and resolve them in one go. | [optional] [default to null]
**Vertical** | **string** | Functional area that this query falls into. | [optional] [default to null]
**VObjectId** | **string** | Identifier for this object in the source NSX endpoint. | [optional] [default to null]
**SuggestedValue** | **string** | The suggested value to resolve this feedback request. | [optional] [default to null]
**Message** | **string** | Detailed feedback request with options. | [optional] [default to null]
**ObjectId** | **string** | Identifier of the object for which feedback is requested. | [optional] [default to null]
**AcceptedValueType** | **string** | Data type of the items listed in acceptable values list. | [optional] [default to null]
**VObjectName** | **string** | Name of this object in the source NSX endpoint. | [optional] [default to null]
**MultiValue** | **bool** | Indicates if multiple values can be selected as response from the list of acceptable value. | [optional] [default to null]
**AcceptedValues** | **[]string** | List of acceptable values for this feedback request. | [optional] [default to null]
**Id** | **string** | Identifier of the feedback request. | [optional] [default to null]
**SuggestedAction** | **string** | The suggested action to resolve this feedback request. | [optional] [default to null]
**SubVertical** | **string** | Functional sub-area that this query falls into. | [optional] [default to null]
**Resolution** | **string** | If the feedback request was resolved earlier, provides details about the previous resolution. | [optional] [default to null]
**Rejected** | **bool** | Indicates if previous response was invalid. Please provide a valid response. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

