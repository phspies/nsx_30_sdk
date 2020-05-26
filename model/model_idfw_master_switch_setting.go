package nsxt

// Identity Firewall master switch setting.  This setting enables or disables Identity Firewall feature across the system.  It affects compute collections, hypervisor and virtual machines.  This operation is expensive and also has big impact and implication on system perforamce. 
type IdfwMasterSwitchSetting struct {
	// IDFW master switch (true=Enabled / false=Disabled).
	IdfwMasterSwitchEnabled bool `json:"idfw_master_switch_enabled"`
}
