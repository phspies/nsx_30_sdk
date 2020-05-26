# PortMirroringFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterAction** | **string** | If set to MIRROR, packets will be mirrored. If set to DO_NOT_MIRROR, packets will not be mirrored. | [optional] [default to FILTER_ACTION.MIRROR]
**IpProtocol** | **string** | The transport protocols of TCP or UDP, used to match the transport protocol of a packet. If not provided, no filtering by IP protocols is performed. | [optional] [default to null]
**SrcIps** | [***IpAddresses**](IPAddresses.md) |  | [optional] [default to null]
**DstIps** | [***IpAddresses**](IPAddresses.md) |  | [optional] [default to null]
**DstPorts** | **string** | Destination port in the form of a port or port range, used to match the destination port of a packet. If not provided, no filtering by destination port is performed. | [optional] [default to null]
**SrcPorts** | **string** | Source port in the form of a port or port range, used to match the source port of a packet. If not provided, no filtering by source port is performed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

