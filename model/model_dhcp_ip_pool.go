package nsxt

// DHCP ip-pool to define dynamic ip allocation ranges.
type DhcpIpPool struct {
	// Lease time, in seconds, [60-(2^32-1)]. Default is 86400.
	LeaseTime int64 `json:"lease_time,omitempty"`
	// Gateway ip address of the allocation.
	GatewayIp string `json:"gateway_ip,omitempty"`
	Options *DhcpOptions `json:"options,omitempty"`
	// Ip-ranges to define dynamic ip allocation ranges.
	AllocationRanges []IpPoolRange `json:"allocation_ranges"`
	// Warning threshold. Alert will be raised if the pool usage reaches the given threshold. 
	WarningThreshold int64 `json:"warning_threshold,omitempty"`
	// Error threshold. Alert will be raised if the pool usage reaches the given threshold. 
	ErrorThreshold int64 `json:"error_threshold,omitempty"`
}
