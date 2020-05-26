# IpSecVpnTunnelTrafficStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PacketsSentOtherError** | **int64** | Total number of packets dropped while sending for any reason. | [optional] [default to null]
**PacketsOut** | **int64** | Total number of outgoing packets on outbound Security association (SA). | [optional] [default to null]
**DroppedPacketsOut** | **int64** | Total number of outgoing packets dropped on outbound security association. | [optional] [default to null]
**IntegrityFailures** | **int64** | Total number of packets dropped due to integrity failures. | [optional] [default to null]
**NomatchingPolicyErrors** | **int64** | Number of packets dropped because of no matching policy is available. | [optional] [default to null]
**SaMismatchErrorsIn** | **int64** | Totoal number of security association (SA) mismatch errors on incoming packets. | [optional] [default to null]
**PeerSubnet** | **string** | Peer subnet to which a tunnel belongs. | [optional] [default to null]
**ReplayErrors** | **int64** | Total number of packets dropped due to replay check on that Security association (SA). | [optional] [default to null]
**BytesOut** | **int64** | Total number of outgoing bytes on outbound Security association (SA). | [optional] [default to null]
**LocalSubnet** | **string** | Local subnet to which a tunnel belongs. | [optional] [default to null]
**DroppedPacketsIn** | **int64** | Total number of incoming packets dropped on inbound security association. | [optional] [default to null]
**EncryptionFailures** | **int64** | Total number of packets dropped because of failure in encryption. | [optional] [default to null]
**SaMismatchErrorsOut** | **int64** | Totoal number of security association (SA) mismatch errors on outgoing packets. | [optional] [default to null]
**TunnelDownReason** | **string** | Gives the detailed reason about the tunnel when it is down. If tunnel is UP tunnel down reason will be empty. | [optional] [default to null]
**PacketsReceiveOtherError** | **int64** | Total number of incoming packets dropped on inbound Security association (SA)(misc). | [optional] [default to null]
**BytesIn** | **int64** | Total number of incoming bytes on inbound Security association (SA). | [optional] [default to null]
**DecryptionFailures** | **int64** | Total number of packets dropped due to decryption failures. | [optional] [default to null]
**SeqNumberOverflowError** | **int64** | Total number of packets dropped while sending due to overflow in sequence number. | [optional] [default to null]
**PacketsIn** | **int64** | Total number of incoming packets on inbound Security association (SA). | [optional] [default to null]
**TunnelStatus** | **string** | Specifies the status of tunnel. If all the SA (Security association) are negotiated then tunnels status will be UP. If negotiation fails for the SAs status will be DOWN, if SAs are in negotiating phase tunnels status will be NEGOTIATING. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

