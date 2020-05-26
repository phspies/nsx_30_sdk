package nsxt

// Host node
type HostNode struct {
	// Discovered IP Addresses of the fabric node, version 4 or 6
	DiscoveredIpAddresses []string `json:"discovered_ip_addresses,omitempty"`
	// IP Addresses of the Node, version 4 or 6. This property is mandatory for all nodes except for automatic deployment of edge virtual machine node. For automatic deployment, the ip address from management_port_subnets property will be considered. 
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// ID of the Node maintained on the Node and used to recognize the Node
	ExternalId string `json:"external_id,omitempty"`
	// Fully qualified domain name of the fabric node
	Fqdn string `json:"fqdn,omitempty"`
	// Fabric node type, for example 'HostNode', 'EdgeNode' or 'PublicCloudGatewayNode'
	ResourceType string `json:"resource_type"`
	// Id of discovered node which was converted to create this node
	DiscoveredNodeId string `json:"discovered_node_id,omitempty"`
	// The id of the vCenter server managing the ESXi type HostNode
	ManagedByServer string `json:"managed_by_server,omitempty"`
	HostCredential *HostNodeLoginCredential `json:"host_credential,omitempty"`
	// Version of the hypervisor operating system
	OsVersion string `json:"os_version,omitempty"`
	// Hypervisor type, for example ESXi or RHEL KVM
	OsType string `json:"os_type"`
	// Specify an installation folder to install the NSX kernel modules for Windows Server. By default, it is C:\\Program Files\\VMware\\NSX\\.
	WindowsInstallLocation string `json:"windows_install_location,omitempty"`
	// Indicates host node's maintenance mode state. The state is ENTERING when a task to put the host in maintenance-mode is in progress. 
	MaintenanceModeState string `json:"maintenance_mode_state,omitempty"`
}
