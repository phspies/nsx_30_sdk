# IpPoolSubnet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**DnsNameservers** | **[]string** | The collection of upto 3 DNS servers for the subnet. | [optional] [default to null]
**Cidr** | **string** | Represents network address and the prefix length which will be associated with a layer-2 broadcast domain | [default to null]
**GatewayIp** | **string** | The default gateway address on a layer-3 router. | [optional] [default to null]
**AllocationRanges** | [**[]IpPoolRange**](IpPoolRange.md) | A collection of IPv4 or IPv6 IP Pool Ranges. | [default to null]
**DnsSuffix** | **string** | The DNS suffix for the DNS server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

