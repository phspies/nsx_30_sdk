package nsxt

type CapacityUsageMetaInfo struct {
	// Indicates the maximum global threshold percentage 
	MaxGlobalThresholdPercentage float64 `json:"max_global_threshold_percentage"`
	// Indicates the minimum global threshold percentage 
	MinGlobalThresholdPercentage float64 `json:"min_global_threshold_percentage"`
	// Timestamp at which capacity usage was last calculated
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp"`
}
