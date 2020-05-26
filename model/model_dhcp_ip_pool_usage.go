package nsxt

type DhcpIpPoolUsage struct {
	// allocated percentage. COULD BE INACCURATE, REFERENCE ONLY.
	AllocatedPercentage int64 `json:"allocated_percentage"`
	// pool size
	PoolSize int64 `json:"pool_size"`
	// allocated number. COULD BE INACCURATE, REFERENCE ONLY.
	AllocatedNumber int64 `json:"allocated_number"`
	// uuid of dhcp ip pool
	DhcpIpPoolId string `json:"dhcp_ip_pool_id"`
}
