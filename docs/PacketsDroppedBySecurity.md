# PacketsDroppedBySecurity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpoofGuardDropped** | [**[]PacketTypeAndCounter**](PacketTypeAndCounter.md) | The packets dropped by \&quot;Spoof Guard\&quot;; supported packet types are IPv4, IPv6, ARP, ND, non-IP. | [optional] [default to null]
**DhcpServerDroppedIpv4** | **int64** | The number of IPv4 packets dropped by \&quot;DHCP server block\&quot;. | [optional] [default to null]
**DhcpServerDroppedIpv6** | **int64** | The number of IPv6 packets dropped by \&quot;DHCP server block\&quot;. | [optional] [default to null]
**DhcpClientDroppedIpv4** | **int64** | The number of IPv4 packets dropped by \&quot;DHCP client block\&quot;. | [optional] [default to null]
**BpduFilterDropped** | **int64** | The number of packets dropped by \&quot;BPDU filter\&quot;. | [optional] [default to null]
**DhcpClientDroppedIpv6** | **int64** | The number of IPv6 packets dropped by \&quot;DHCP client block\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

