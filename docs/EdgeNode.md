# EdgeNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveredIpAddresses** | **[]string** | Discovered IP Addresses of the fabric node, version 4 or 6 | [optional] [default to null]
**IpAddresses** | **[]string** | IP Addresses of the Node, version 4 or 6. This property is mandatory for all nodes except for automatic deployment of edge virtual machine node. For automatic deployment, the ip address from management_port_subnets property will be considered.  | [optional] [default to null]
**ExternalId** | **string** | ID of the Node maintained on the Node and used to recognize the Node | [optional] [default to null]
**Fqdn** | **string** | Fully qualified domain name of the fabric node | [optional] [default to null]
**ResourceType** | **string** | Fabric node type, for example &#x27;HostNode&#x27;, &#x27;EdgeNode&#x27; or &#x27;PublicCloudGatewayNode&#x27; | [default to null]
**NodeSettings** | [***EdgeNodeSettings**](EdgeNodeSettings.md) |  | [optional] [default to null]
**DeploymentConfig** | [***EdgeNodeDeploymentConfig**](EdgeNodeDeploymentConfig.md) |  | [optional] [default to null]
**AllocationList** | **[]string** | List of logical router ids to which this edge node is allocated. | [optional] [default to null]
**DeploymentType** | **string** | Supported edge deployment type. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

