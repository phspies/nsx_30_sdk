package nsxt

// Abstract base type for telemetry schedule configuration
type TelemetrySchedule struct {
	// Specify one of DailyTelemetrySchedule, WeeklyTelemetrySchedule, or MonthlyTelemetrySchedule.
	FrequencyType string `json:"frequency_type"`
}
