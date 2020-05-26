package nsxt

// IP assignment specification for Static IP Pool.
type StaticIpPoolSpec struct {
	ResourceType string `json:"resource_type"`
	IpPoolId string `json:"ip_pool_id"`
}
