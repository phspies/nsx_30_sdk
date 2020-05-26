# PropertyItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Represents field value of the property. | [default to null]
**Separator** | **bool** | If true, separates this property in a widget. | [optional] [default to false]
**Navigation** | **string** | Hyperlink of the specified UI page that provides details. This will be linked with value of the property. | [optional] [default to null]
**RenderConfiguration** | [**[]RenderConfiguration**](RenderConfiguration.md) | Render configuration to be applied, if any. | [optional] [default to null]
**Type_** | **string** | Data type of the field. | [default to TYPE_.STRING_]
**Heading** | **bool** | Set to true if the field is a heading. Default is false. | [optional] [default to false]
**Condition** | **string** | If the condition is met then the property will be displayed. Examples of expression syntax are provided under &#x27;example_request&#x27; section of &#x27;CreateWidgetConfiguration&#x27; API. | [optional] [default to null]
**Label** | [***Label**](Label.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

