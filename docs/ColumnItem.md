# ColumnItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortKey** | **string** | Sorting on column is based on the sort_key. sort_key represents the field in the output data on which sort is requested. | [optional] [default to null]
**Type_** | **string** | Data type of the field. | [default to TYPE_.STRING_]
**Tooltip** | [**[]Tooltip**](Tooltip.md) | Multi-line text to be shown on tooltip while hovering over a cell in the grid. | [optional] [default to null]
**Label** | [***Label**](Label.md) |  | [default to null]
**Field** | **string** | Field from which values of the column will be derived. | [default to null]
**SortAscending** | **bool** | If true, the value of the column are sorted in ascending order. Otherwise, in descending order. | [optional] [default to true]
**DrilldownId** | **string** | Id of drilldown widget, if any. Id should be a valid id of an existing widget. | [optional] [default to null]
**Hidden** | **bool** | If set to true, hides the column | [optional] [default to false]
**Navigation** | **string** | Hyperlink of the specified UI page that provides details. If drilldown_id is provided, then navigation cannot be used. | [optional] [default to null]
**ColumnIdentifier** | **string** | Identifies the column and used for fetching content upon an user click or drilldown. If column identifier is not provided, the column&#x27;s data will not participate in searches and drilldowns. | [optional] [default to null]
**RenderConfiguration** | [**[]RenderConfiguration**](RenderConfiguration.md) | Render configuration to be applied, if any. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

