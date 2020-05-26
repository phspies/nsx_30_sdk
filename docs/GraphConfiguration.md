# GraphConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFilterValue** | [**[]DefaultFilterValue**](DefaultFilterValue.md) | Default filter values to be passed to datasources. This will be used when the report is requested without filter values. | [optional] [default to null]
**DisplayName** | **string** | Title of the widget. If display_name is omitted, the widget will be shown without a title. | [optional] [default to null]
**Datasources** | [**[]Datasource**](Datasource.md) | The &#x27;datasources&#x27; represent the sources from which data will be fetched. Currently, only NSX-API is supported as a &#x27;default&#x27; datasource. An example of specifying &#x27;default&#x27; datasource along with the urls to fetch data from is given at &#x27;example_request&#x27; section of &#x27;CreateWidgetConfiguration&#x27; API. | [optional] [default to null]
**Weight** | **int32** | Specify relavite weight in WidgetItem for placement in a view. Please see WidgetItem for details. | [optional] [default to null]
**Footer** | [***Footer**](Footer.md) |  | [optional] [default to null]
**FilterValueRequired** | **bool** | Flag to indicate that widget will continue to work without filter value. If this flag is set to false then default_filter_value is manadatory. | [optional] [default to true]
**Span** | **int32** | Represents the horizontal span of the widget / container. | [optional] [default to null]
**Icons** | [**[]Icon**](Icon.md) | Icons to be applied at dashboard for widgets and UI elements. | [optional] [default to null]
**IsDrilldown** | **bool** | Set to true if this widget should be used as a drilldown. | [optional] [default to false]
**Filter** | **string** | Id of filter widget for subscription, if any. Id should be a valid id of an existing filter widget. Filter widget should be from the same view. Datasource URLs should have placeholder values equal to filter alias to accept the filter value on filter change. | [optional] [default to null]
**DrilldownId** | **string** | Id of drilldown widget, if any. Id should be a valid id of an existing widget. A widget is considered as drilldown widget when it is associated with any other widget and provides more detailed information about any data item from the parent widget. | [optional] [default to null]
**Shared** | **bool** | Please use the property &#x27;shared&#x27; of View instead of this. The widgets of a shared view are visible to other users. | [optional] [default to null]
**Legend** | [***Legend**](Legend.md) |  | [optional] [default to null]
**ResourceType** | **string** | Supported visualization types are LabelValueConfiguration, DonutConfiguration, GridConfiguration, StatsConfiguration, MultiWidgetConfiguration, GraphConfiguration, ContainerConfiguration, CustomWidgetConfiguration and DropdownFilterWidgetConfiguration. | [default to null]
**Graphs** | [**[]GraphDefinition**](GraphDefinition.md) | Graphs | [default to null]
**Axes** | [***Axes**](Axes.md) |  | [optional] [default to null]
**Navigation** | **string** | Hyperlink of the specified UI page that provides details. | [optional] [default to null]
**SubType** | **string** | Describes the the type of graph. LINE_GRAPH shows a line graph chart BAR_GRAPH shows a simple bar graph chart STACKED_BAR_GRAPH shows a stacked bar graph chart | [optional] [default to SUB_TYPE.BAR_GRAPH]
**DisplayXValue** | **bool** | If true, value of a point is shown as label on X axis. If false, value of point is not shown as label on X axis. false can be useful in situations where there are too many points and showing the X value as label can clutter the X axis. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

