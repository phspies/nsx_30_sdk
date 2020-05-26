package nsxt

type EdgeNode struct {
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
	NodeSettings *EdgeNodeSettings `json:"node_settings,omitempty"`
	DeploymentConfig *EdgeNodeDeploymentConfig `json:"deployment_config,omitempty"`
	// List of logical router ids to which this edge node is allocated.
	AllocationList []string `json:"allocation_list,omitempty"`
	// Supported edge deployment type.
	DeploymentType string `json:"deployment_type,omitempty"`
}
