# TunnelProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Status** | **string** | Status of tunnel | [optional] [default to null]
**EgressInterface** | **string** | Corresponds to the interface where local_ip_address is routed. | [optional] [default to null]
**RemoteNodeId** | **string** | UUID of the remote transport node | [optional] [default to null]
**Bfd** | [***BfdProperties**](BFDProperties.md) |  | [optional] [default to null]
**LocalIp** | **string** | Local IP address of tunnel | [optional] [default to null]
**LastUpdatedTime** | **int64** | Time at which the Tunnel status has been fetched last time. | [optional] [default to null]
**Name** | **string** | Name of tunnel | [optional] [default to null]
**RemoteNodeDisplayName** | **string** | Represents the display name of the remote transport node at the other end of the tunnel. | [optional] [default to null]
**Encap** | **string** | Tunnel encap | [optional] [default to null]
**LatencyType** | **string** | Latency type. | [optional] [default to null]
**LatencyValue** | **int64** | The latency value is set only when latency_type is VALID. | [optional] [default to null]
**RemoteIp** | **string** | Remote IP address of tunnel | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

