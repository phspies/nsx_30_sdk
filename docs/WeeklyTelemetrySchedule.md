# WeeklyTelemetrySchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrequencyType** | **string** | Specify one of DailyTelemetrySchedule, WeeklyTelemetrySchedule, or MonthlyTelemetrySchedule. | [default to null]
**Minutes** | **int64** | Minute at which data will be collected. Specify a value between 0 through 59.  | [optional] [default to 0]
**HourOfDay** | **int64** | Hour at which data will be collected. Specify a value between 0 through 23.  | [default to null]
**DayOfWeek** | **string** | Day of week on which data will be collected. Specify one of SUNDAY through SATURDAY.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

