package nsxt

type DailyTelemetrySchedule struct {
	// Specify one of DailyTelemetrySchedule, WeeklyTelemetrySchedule, or MonthlyTelemetrySchedule.
	FrequencyType string `json:"frequency_type"`
	// Minute at which data will be collected. Specify a value between 0 through 59. 
	Minutes int64 `json:"minutes,omitempty"`
	// Hour at which data will be collected. Specify a value between 0 through 23. 
	HourOfDay int64 `json:"hour_of_day"`
}
