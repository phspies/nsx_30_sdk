# VsphereDeploymentConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementType** | **string** |  | [default to null]
**DataNetworkIds** | **[]string** | List of distributed portgroup or VLAN logical identifiers to which the datapath serving vnics of edge node vm will be connected.  | [default to null]
**ResourceAllocation** | [***ResourceAssignment**](ResourceAssignment.md) |  | [optional] [default to null]
**DnsServers** | **[]string** | List of DNS servers. This field is deprecated. Use dns_servers property in EdgeNodeSettings section when creating or updating transport nodes.  | [optional] [default to null]
**NtpServers** | **[]string** | List of NTP servers. This field is deprecated. Use ntp_servers property in EdgeNodeSettings section when creating or updating transport nodes.  | [optional] [default to null]
**ManagementNetworkId** | **string** | Distributed portgroup identifier to which the management vnic of edge node vm will be connected. This portgroup must have connectivity with MP and CCP. A VLAN logical switch identifier may also be specified.  | [default to null]
**EnableSsh** | **bool** | Enabling SSH service is not recommended for security reasons. This field is deprecated. Use enable_ssh property in EdgeNodeSettings section when creating or updating transport nodes.  | [optional] [default to false]
**AllowSshRootLogin** | **bool** | Allowing root SSH logins is not recommended for security reasons. This field is deprecated. Use allow_ssh_root_login property in EdgeNodeSettings section when creating transport nodes.  | [optional] [default to false]
**ComputeId** | **string** | The edge node vm will be deployed on the specified cluster or resourcepool. Note - all the hosts must have nsx fabric prepared in the specified cluster.  | [default to null]
**SearchDomains** | **[]string** | List of domain names that are used to complete unqualified host names. This field is deprecated. Use search_domains property in EdgeNodeSettings section when creating or updating transport nodes.  | [optional] [default to null]
**ReservationInfo** | [***ReservationInfo**](ReservationInfo.md) |  | [optional] [default to null]
**VcId** | **string** | The vc specific identifiers will be resolved on this VC. So all other identifiers specified here must belong to this vcenter server.  | [default to null]
**StorageId** | **string** | The edge node vm will be deployed on the specified datastore. User must ensure that storage is accessible by the specified cluster/host.  | [default to null]
**DefaultGatewayAddresses** | **[]string** | The default gateway for edge node must be specified if all the nodes it communicates with are not in the same subnet. Note: Only single IPv4 default gateway address is supported and it must belong to management network.  | [optional] [default to null]
**ManagementPortSubnets** | [**[]IpSubnet**](IPSubnet.md) | IP Address and subnet configuration for the management port. Note: only one IPv4 address is supported for the management port.  | [optional] [default to null]
**HostId** | **string** | The edge node vm will be deployed on the specified Host within the cluster if host_id is specified. Note - User must ensure that storage and specified networks are accessible by this host.  | [optional] [default to null]
**Hostname** | **string** | Host name or FQDN for edge node. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

