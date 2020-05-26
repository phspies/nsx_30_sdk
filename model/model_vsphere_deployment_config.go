package nsxt

// The Vsphere deployment configuration determines where to deploy the edge node. It contains settings that are applied during install time. If using DHCP, you must leave the following fields unset: search_domains, management_port_subnets, dns_servers and default_gateway_addresses. Use EdgeNodeSettings to specify host name, SSH, NTP and DNS settings for both deployment and consequent update. These settings are editable for manually deployed edge nodes as well. EdgeNodeSettings reports current values configured on the edge node. The following fields are deprecated Hostname, SSH, NTP and DNS settings. 
type VsphereDeploymentConfig struct {
	PlacementType string `json:"placement_type"`
	// List of distributed portgroup or VLAN logical identifiers to which the datapath serving vnics of edge node vm will be connected. 
	DataNetworkIds []string `json:"data_network_ids"`
	ResourceAllocation *ResourceAssignment `json:"resource_allocation,omitempty"`
	// List of DNS servers. This field is deprecated. Use dns_servers property in EdgeNodeSettings section when creating or updating transport nodes. 
	DnsServers []string `json:"dns_servers,omitempty"`
	// List of NTP servers. This field is deprecated. Use ntp_servers property in EdgeNodeSettings section when creating or updating transport nodes. 
	NtpServers []string `json:"ntp_servers,omitempty"`
	// Distributed portgroup identifier to which the management vnic of edge node vm will be connected. This portgroup must have connectivity with MP and CCP. A VLAN logical switch identifier may also be specified. 
	ManagementNetworkId string `json:"management_network_id"`
	// Enabling SSH service is not recommended for security reasons. This field is deprecated. Use enable_ssh property in EdgeNodeSettings section when creating or updating transport nodes. 
	EnableSsh bool `json:"enable_ssh,omitempty"`
	// Allowing root SSH logins is not recommended for security reasons. This field is deprecated. Use allow_ssh_root_login property in EdgeNodeSettings section when creating transport nodes. 
	AllowSshRootLogin bool `json:"allow_ssh_root_login,omitempty"`
	// The edge node vm will be deployed on the specified cluster or resourcepool. Note - all the hosts must have nsx fabric prepared in the specified cluster. 
	ComputeId string `json:"compute_id"`
	// List of domain names that are used to complete unqualified host names. This field is deprecated. Use search_domains property in EdgeNodeSettings section when creating or updating transport nodes. 
	SearchDomains []string `json:"search_domains,omitempty"`
	ReservationInfo *ReservationInfo `json:"reservation_info,omitempty"`
	// The vc specific identifiers will be resolved on this VC. So all other identifiers specified here must belong to this vcenter server. 
	VcId string `json:"vc_id"`
	// The edge node vm will be deployed on the specified datastore. User must ensure that storage is accessible by the specified cluster/host. 
	StorageId string `json:"storage_id"`
	// The default gateway for edge node must be specified if all the nodes it communicates with are not in the same subnet. Note: Only single IPv4 default gateway address is supported and it must belong to management network. 
	DefaultGatewayAddresses []string `json:"default_gateway_addresses,omitempty"`
	// IP Address and subnet configuration for the management port. Note: only one IPv4 address is supported for the management port. 
	ManagementPortSubnets []IpSubnet `json:"management_port_subnets,omitempty"`
	// The edge node vm will be deployed on the specified Host within the cluster if host_id is specified. Note - User must ensure that storage and specified networks are accessible by this host. 
	HostId string `json:"host_id,omitempty"`
	// Host name or FQDN for edge node.
	Hostname string `json:"hostname,omitempty"`
}