package nsxt

// Identity Firewall standalone hosts switch setting. This setting enables or disables Identity Firewall feature on all standalone hosts. 
type IdfwStandaloneHostsSwitchSetting struct {
	// IDFW standalone hosts switch (true=Enabled / false=Disabled).
	StandaloneHostsEnabled bool `json:"standalone_hosts_enabled"`
}
