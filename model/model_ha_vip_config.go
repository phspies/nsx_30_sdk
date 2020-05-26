package nsxt

type HaVipConfig struct {
	// Array of IP address subnets which will be used as floating IP addresses. | Note - this configuration is applicable only for Active-Standby LogicalRouter. | For Active-Active LogicalRouter this configuration will be rejected.
	HaVipSubnets []VipSubnet `json:"ha_vip_subnets"`
	// Identifiers of logical router uplink ports which are to be paired to provide | redundancy. Floating IP will be owned by one of these uplink ports (depending upon | which node is Active).
	RedundantUplinkPortIds []string `json:"redundant_uplink_port_ids"`
	// Flag to enable this ha vip config.
	Enabled bool `json:"enabled,omitempty"`
}
