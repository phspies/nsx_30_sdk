# TransportNodeRemoteTunnelEndpointConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamedTeamingPolicy** | **string** | Specifying this field will override the default teaming policy of the host switch and will be used by remote tunnel endpoint traffic. | [optional] [default to null]
**HostSwitchName** | **string** | The host switch name should reference an existing host switch specified in the transport node configuration. The name will be used to identify the host switch responsible for processing remote tunnel endpoint traffic. | [default to null]
**RtepVlan** | **int64** | The transport VLAN id used for tagging intersite overlay traffic between remote tunnel endpoints. | [default to null]
**IpAssignmentSpec** | [***IpAssignmentSpec**](IpAssignmentSpec.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

