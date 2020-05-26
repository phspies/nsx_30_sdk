# NodeNetworkInterfaceProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**PhysicalAddress** | **string** | Interface MAC address | [optional] [default to null]
**BroadcastAddress** | **string** | Interface broadcast address | [optional] [default to null]
**LinkStatus** | **string** | Interface administration status | [optional] [default to null]
**DefaultGateway** | **string** | Interface&#x27;s default gateway | [optional] [default to null]
**BondPrimary** | **string** | Bond&#x27;s primary device name in active-backup bond mode | [optional] [default to null]
**BondSlaves** | **[]string** | Bond&#x27;s slave devices | [optional] [default to null]
**IpAddresses** | [**[]IPv4AddressProperties**](IPv4AddressProperties.md) | Interface IP addresses | [optional] [default to null]
**Vlan** | **int64** | VLAN Id | [optional] [default to null]
**BondMode** | **string** | Bond mode | [optional] [default to null]
**InterfaceId** | **string** | Interface ID | [optional] [default to null]
**AdminStatus** | **string** | Interface administration status | [optional] [default to null]
**Plane** | **string** | Interface plane | [optional] [default to null]
**IsKni** | **bool** | Interface is a KNI | [optional] [default to null]
**IpConfiguration** | **string** | Interface configuration | [default to null]
**Mtu** | **int64** | Interface MTU | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

