package nsxt

// Contains a set of information of a VM on the network interfaces present on the partner appliance that needs to be configured by the NSX Manager.
type VmNicInfo struct {
	// Set of information of a VM on the network interfaces present on the partner appliance that needs to be configured by the NSX Manager.
	NicInfos []NicInfo `json:"nic_infos"`
}
