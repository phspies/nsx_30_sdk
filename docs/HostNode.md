# HostNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveredIpAddresses** | **[]string** | Discovered IP Addresses of the fabric node, version 4 or 6 | [optional] [default to null]
**IpAddresses** | **[]string** | IP Addresses of the Node, version 4 or 6. This property is mandatory for all nodes except for automatic deployment of edge virtual machine node. For automatic deployment, the ip address from management_port_subnets property will be considered.  | [optional] [default to null]
**ExternalId** | **string** | ID of the Node maintained on the Node and used to recognize the Node | [optional] [default to null]
**Fqdn** | **string** | Fully qualified domain name of the fabric node | [optional] [default to null]
**ResourceType** | **string** | Fabric node type, for example &#x27;HostNode&#x27;, &#x27;EdgeNode&#x27; or &#x27;PublicCloudGatewayNode&#x27; | [default to null]
**DiscoveredNodeId** | **string** | Id of discovered node which was converted to create this node | [optional] [default to null]
**ManagedByServer** | **string** | The id of the vCenter server managing the ESXi type HostNode | [optional] [default to null]
**HostCredential** | [***HostNodeLoginCredential**](HostNodeLoginCredential.md) |  | [optional] [default to null]
**OsVersion** | **string** | Version of the hypervisor operating system | [optional] [default to null]
**OsType** | **string** | Hypervisor type, for example ESXi or RHEL KVM | [default to null]
**WindowsInstallLocation** | **string** | Specify an installation folder to install the NSX kernel modules for Windows Server. By default, it is C:\\Program Files\\VMware\\NSX\\. | [optional] [default to null]
**MaintenanceModeState** | **string** | Indicates host node&#x27;s maintenance mode state. The state is ENTERING when a task to put the host in maintenance-mode is in progress.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

