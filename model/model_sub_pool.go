package nsxt

type SubPool struct {
	// Percentage utlization of sub-pool based on the number of services configured and the hard limits, if any. 
	UsagePercentage float64 `json:"usage_percentage,omitempty"`
	// Credits remaining on the sub-pool that can be used to deploy services of corresponding sub-pool type. 
	RemainingCreditNumber int32 `json:"remaining_credit_number,omitempty"`
	// Type of the sub-pool configured on edge node.
	SubPoolType string `json:"sub_pool_type,omitempty"`
}
