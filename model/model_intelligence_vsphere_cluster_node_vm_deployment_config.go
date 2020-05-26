package nsxt

// The Vsphere deployment configuration determines where to deploy the NSX-Intelligence cluster node VM through a vCenter server. It contains settings that are applied during install time. If using DHCP, the following fields must be left unset - dns_servers, management_port_subnets, and default_gateway_addresses 
type IntelligenceVsphereClusterNodeVmDeploymentConfig struct {
	// Specifies the config for the platform through which to deploy the VM 
	PlacementType string `json:"placement_type"`
	// The NSX-Intelligence cluster node VM OVF URL to download and install the OVF file. This field is deprecated now. Please upload OVA file using \"/repository/bundles\" API and then try deployment without providing this field. 
	OvfUrl string `json:"ovf_url,omitempty"`
	// List of DNS servers. If DHCP is used, the default DNS servers associated with the DHCP server will be used instead. Required if using static IP. 
	DnsServers []string `json:"dns_servers,omitempty"`
	// Desired display name for NSX-Intelligence VM to be deployed 
	DisplayName string `json:"display_name,omitempty"`
	// List of NTP servers. To use hostnames, a DNS server must be defined. If not using DHCP, a DNS server should be specified under dns_servers. 
	NtpServers []string `json:"ntp_servers,omitempty"`
	// Desired host name/FQDN for the VM to be deployed 
	Hostname string `json:"hostname"`
	// If true, the SSH service will automatically be started on the VM. Enabling SSH service is not recommended for security reasons. 
	EnableSsh bool `json:"enable_ssh,omitempty"`
	// If true, the root user will be allowed to log into the VM. Allowing root SSH logins is not recommended for security reasons. 
	AllowSshRootLogin bool `json:"allow_ssh_root_login,omitempty"`
	// The NSX-Intelligence cluster node VM will be deployed on the specified cluster or resourcepool for specified VC server. 
	ComputeId string `json:"compute_id"`
	// Specifies the disk provisioning type of the VM. 
	DiskProvisioning string `json:"disk_provisioning,omitempty"`
	// The VC-specific identifiers will be resolved on this VC, so all other identifiers specified in the config must belong to this vCenter server. 
	VcId string `json:"vc_id"`
	// The NSX-Intelligence cluster node VM will be deployed on the specified datastore in the specified VC server. User must ensure that storage is accessible by the specified cluster/host. 
	StorageId string `json:"storage_id"`
	// The default gateway for the VM to be deployed must be specified if all the other VMs it communicates with are not in the same subnet. Do not specify this field and management_port_subnets to use DHCP. Note: only single IPv4 default gateway address is supported and it must belong to management network. IMPORTANT: VMs deployed using DHCP are currently not supported, so this parameter should be specified. 
	DefaultGatewayAddresses []string `json:"default_gateway_addresses,omitempty"`
	// IP Address and subnet configuration for the management port. Do not specify this field and default_gateway_addresses to use DHCP. Note: only one IPv4 address is supported for the management port. IMPORTANT: VMs deployed using DHCP are currently not supported, so this parameter should be specified. 
	ManagementPortSubnets []IpSubnet `json:"management_port_subnets,omitempty"`
	// The NSX-Intelligence cluster node VM will be deployed on the specified host in the specified VC server within the cluster if host_id is specified. Note: User must ensure that storage and specified networks are accessible by this host. 
	HostId string `json:"host_id,omitempty"`
	// Distributed portgroup identifier to which the management vnic of NSX-Intelligence cluster node VM will be connected. 
	ManagementNetworkId string `json:"management_network_id"`
}
