# NdpHeader

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DstIp** | **string** | The IP address of the destination of the solicitation. It MUST NOT be a multicast address. | [optional] [default to null]
**MsgType** | **string** | This field specifies the type of the Neighbor discover message being sent. NEIGHBOR_SOLICITATION - Neighbor Solicitation message to discover the link-layer address of an on-link IPv6 node or to confirm a previously determined link-layer address. NEIGHBOR_ADVERTISEMENT - Neighbor Advertisement message in response to a Neighbor Solicitation message. | [optional] [default to MSG_TYPE.SOLICITATION]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

