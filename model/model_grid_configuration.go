package nsxt

// Represents configuration of a Grid or Table widget.
type GridConfiguration struct {
	// Default filter values to be passed to datasources. This will be used when the report is requested without filter values.
	DefaultFilterValue []DefaultFilterValue `json:"default_filter_value,omitempty"`
	// Title of the widget. If display_name is omitted, the widget will be shown without a title.
	DisplayName string `json:"display_name,omitempty"`
	// The 'datasources' represent the sources from which data will be fetched. Currently, only NSX-API is supported as a 'default' datasource. An example of specifying 'default' datasource along with the urls to fetch data from is given at 'example_request' section of 'CreateWidgetConfiguration' API.
	Datasources []Datasource `json:"datasources,omitempty"`
	// Specify relavite weight in WidgetItem for placement in a view. Please see WidgetItem for details.
	Weight int32 `json:"weight,omitempty"`
	Footer *Footer `json:"footer,omitempty"`
	// Flag to indicate that widget will continue to work without filter value. If this flag is set to false then default_filter_value is manadatory.
	FilterValueRequired bool `json:"filter_value_required,omitempty"`
	// Represents the horizontal span of the widget / container.
	Span int32 `json:"span,omitempty"`
	// Icons to be applied at dashboard for widgets and UI elements.
	Icons []Icon `json:"icons,omitempty"`
	// Set to true if this widget should be used as a drilldown.
	IsDrilldown bool `json:"is_drilldown,omitempty"`
	// Id of filter widget for subscription, if any. Id should be a valid id of an existing filter widget. Filter widget should be from the same view. Datasource URLs should have placeholder values equal to filter alias to accept the filter value on filter change.
	Filter string `json:"filter,omitempty"`
	// Id of drilldown widget, if any. Id should be a valid id of an existing widget. A widget is considered as drilldown widget when it is associated with any other widget and provides more detailed information about any data item from the parent widget.
	DrilldownId string `json:"drilldown_id,omitempty"`
	// Please use the property 'shared' of View instead of this. The widgets of a shared view are visible to other users.
	Shared bool `json:"shared,omitempty"`
	Legend *Legend `json:"legend,omitempty"`
	// Supported visualization types are LabelValueConfiguration, DonutConfiguration, GridConfiguration, StatsConfiguration, MultiWidgetConfiguration, GraphConfiguration, ContainerConfiguration, CustomWidgetConfiguration and DropdownFilterWidgetConfiguration.
	ResourceType string `json:"resource_type"`
	// Number of records per page. page_size will be effective only when the urls provided in the datasource support paging.
	PageSize int32 `json:"page_size,omitempty"`
	// Rows of grid or table are formed from the list of objects returned by a row list field.
	RowListFields []RowListField `json:"row_list_fields"`
	// Array of columns of a Grid widget
	Columns []ColumnItem `json:"columns"`
}
