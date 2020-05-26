package nsxt

type CapacityDashboardUsage struct {
	// Indicate the current usage count of object type. 
	CurrentUsageCount int64 `json:"current_usage_count"`
	// This indicates the maximum threshold percentage for object type. 
	MaxThresholdPercentage float64 `json:"max_threshold_percentage"`
	// Display name for NSX object type. 
	DisplayName string `json:"display_name"`
	// Severity calculated from percentage usage 
	Severity string `json:"severity"`
	// This is the maximum supported count for object type in consideration. 
	MaxSupportedCount int64 `json:"max_supported_count"`
	// Indicate the object type for which usage is calculated. 
	UsageType string `json:"usage_type"`
	// This indicates the minimum threshold percentage for object type. 
	MinThresholdPercentage float64 `json:"min_threshold_percentage"`
	// Current usage percentage for object type 
	CurrentUsagePercentage float64 `json:"current_usage_percentage"`
}