package nsxt

type MonthlyTelemetrySchedule struct {
	// Specify one of DailyTelemetrySchedule, WeeklyTelemetrySchedule, or MonthlyTelemetrySchedule.
	FrequencyType string `json:"frequency_type"`
	// Minute at which data will be collected. Specify a value between 0 through 59. 
	Minutes int64 `json:"minutes,omitempty"`
	// Day of month on which data will be collected. Specify a value between 1 through 31. 
	DayOfMonth int64 `json:"day_of_month"`
	// Hour at which data will be collected. Specify a value between 0 through 23. 
	HourOfDay int64 `json:"hour_of_day"`
}
