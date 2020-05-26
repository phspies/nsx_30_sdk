package nsxt

type CapacityThreshold struct {
	// Set the maximum threshold percentage. Specify a value between 0 and 100. Usage percentage above this value is tagged as critical. 
	MaxThresholdPercentage float64 `json:"max_threshold_percentage"`
	// Indicate the object type for which threshold is to be set. 
	ThresholdType string `json:"threshold_type"`
	// Set the minimum threshold percentage. Specify a value between 0 and 100. Usage percentage above this value is tagged as warning. 
	MinThresholdPercentage float64 `json:"min_threshold_percentage"`
}
