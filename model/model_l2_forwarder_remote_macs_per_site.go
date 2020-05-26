package nsxt

type L2ForwarderRemoteMacsPerSite struct {
	// Remote standby IP addresses.
	RemoteStandbyIps []string `json:"remote_standby_ips,omitempty"`
	// 32 bit unique RTEP group id of the logical switch per site. 
	RtepGroupId int64 `json:"rtep_group_id,omitempty"`
	RemoteSite *ResourceReference `json:"remote_site,omitempty"`
	// Remote active IP addresses.
	RemoteActiveIps []string `json:"remote_active_ips,omitempty"`
	// Remote mac addresses.
	RemoteMacAddresses []string `json:"remote_mac_addresses,omitempty"`
}
