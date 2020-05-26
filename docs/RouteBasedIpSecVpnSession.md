# RouteBasedIpSecVpnSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerEndpointId** | **string** | Peer endpoint identifier. | [default to null]
**IpsecVpnServiceId** | **string** | Identifier of VPN Service linked with local endpoint. | [optional] [default to null]
**LocalEndpointId** | **string** | Local endpoint identifier. | [default to null]
**TcpMssClamping** | [***TcpMssClamping**](TcpMssClamping.md) |  | [optional] [default to null]
**Enabled** | **bool** | Enable/Disable IPSec VPN session. | [optional] [default to true]
**ResourceType** | **string** | A Policy Based VPN requires to define protect rules that match   local and peer subnets. IPSec security associations is   negotiated for each pair of local and peer subnet. A Route Based VPN is more flexible, more powerful and recommended over   policy based VPN. IP Tunnel port is created and all traffic routed via   tunnel port is protected. Routes can be configured statically   or can be learned through BGP. A route based VPN is must for establishing   redundant VPN session to remote site.  | [default to null]
**TunnelPorts** | [**[]TunnelPortConfig**](TunnelPortConfig.md) | IP Tunnel ports. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

