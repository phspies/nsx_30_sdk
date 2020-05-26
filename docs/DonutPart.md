# DonutPart

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | A numerical value that represents the portion or entity of the donut or stats chart. | [default to null]
**HideEmptyLegend** | **bool** | If true, legend will be shown only if the data for the part is available. This is applicable only if legends are specified in widget configuration. | [optional] [default to false]
**Condition** | **string** | If the condition is met then the part will be displayed. Examples of expression syntax are provided under &#x27;example_request&#x27; section of &#x27;CreateWidgetConfiguration&#x27; API. | [optional] [default to null]
**DrilldownId** | **string** | Id of drilldown widget, if any. Id should be a valid id of an existing widget. A widget is considered as drilldown widget when it is associated with any other widget and provides more detailed information about any data item from the parent widget. | [optional] [default to null]
**Label** | [***Label**](Label.md) |  | [optional] [default to null]
**Navigation** | **string** | Hyperlink of the specified UI page that provides details. If drilldown_id is provided, then navigation cannot be used. | [optional] [default to null]
**Tooltip** | [**[]Tooltip**](Tooltip.md) | Multi-line text to be shown on tooltip while hovering over the portion. | [optional] [default to null]
**RenderConfiguration** | [**[]RenderConfiguration**](RenderConfiguration.md) | Additional rendering or conditional evaluation of the field values to be performed, if any. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

