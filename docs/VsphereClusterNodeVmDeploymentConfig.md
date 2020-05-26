# VsphereClusterNodeVmDeploymentConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementType** | **string** | Specifies the config for the platform through which to deploy the VM  | [default to null]
**DnsServers** | **[]string** | List of DNS servers. If DHCP is used, the default DNS servers associated with the DHCP server will be used instead. Required if using static IP.  | [optional] [default to null]
**SearchDomains** | **[]string** | List of domain names that are used to complete unqualified host names.  | [optional] [default to null]
**ManagementNetworkId** | **string** | Distributed portgroup identifier to which the management vnic of cluster node VM will be connected.  | [default to null]
**EnableSsh** | **bool** | If true, the SSH service will automatically be started on the VM. Enabling SSH service is not recommended for security reasons.  | [optional] [default to false]
**AllowSshRootLogin** | **bool** | If true, the root user will be allowed to log into the VM. Allowing root SSH logins is not recommended for security reasons.  | [optional] [default to false]
**ComputeId** | **string** | The cluster node VM will be deployed on the specified cluster or resourcepool for specified VC server.  | [default to null]
**NtpServers** | **[]string** | List of NTP servers. To use hostnames, a DNS server must be defined. If not using DHCP, a DNS server should be specified under dns_servers.  | [optional] [default to null]
**DiskProvisioning** | **string** | Specifies the disk provisioning type of the VM.  | [optional] [default to DISK_PROVISIONING.THIN]
**VcId** | **string** | The VC-specific identifiers will be resolved on this VC, so all other identifiers specified in the config must belong to this vCenter server.  | [default to null]
**StorageId** | **string** | The cluster node VM will be deployed on the specified datastore in the specified VC server. User must ensure that storage is accessible by the specified cluster/host.  | [default to null]
**DefaultGatewayAddresses** | **[]string** | The default gateway for the VM to be deployed must be specified if all the other VMs it communicates with are not in the same subnet. Do not specify this field and management_port_subnets to use DHCP. Note: only single IPv4 default gateway address is supported and it must belong to management network. IMPORTANT: VMs deployed using DHCP are currently not supported, so this parameter should be specified.  | [optional] [default to null]
**ManagementPortSubnets** | [**[]IpSubnet**](IPSubnet.md) | IP Address and subnet configuration for the management port. Do not specify this field and default_gateway_addresses to use DHCP. Note: only one IPv4 address is supported for the management port. IMPORTANT: VMs deployed using DHCP are currently not supported, so this parameter should be specified.  | [optional] [default to null]
**HostId** | **string** | The cluster node VM will be deployed on the specified host in the specified VC server within the cluster if host_id is specified. Note: User must ensure that storage and specified networks are accessible by this host.  | [optional] [default to null]
**Hostname** | **string** | Desired host name/FQDN for the VM to be deployed | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

