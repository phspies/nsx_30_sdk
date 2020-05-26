package nsxt

// Specify limit, shares and reservation for all kinds of traffic. Values for limit and reservation are expressed in percentage. And for shares, the value is expressed as a number between 1-100. The overall reservation among all traffic types should not exceed 75%. Otherwise, the API request will be rejected. 
type ResourceAllocation struct {
	// Minimum guaranteed bandwidth percentage
	Reservation float64 `json:"reservation"`
	TrafficType *HostInfraTrafficType `json:"traffic_type"`
	// The limit property specifies the maximum bandwidth allocation for a given traffic type and is expressed in percentage. The default value for this field is set to -1 which means the traffic is unbounded for the traffic type. All other negative values for this property is not supported and will be rejected by the API. 
	Limit float64 `json:"limit"`
	// Shares
	Shares int32 `json:"shares"`
}