# NodeRouteProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Src** | **string** | Source address to prefer when sending to destinations of route | [optional] [default to null]
**FromAddress** | **string** | From address | [optional] [default to null]
**Proto** | **string** | Routing protocol identifier of route | [optional] [default to PROTO.BOOT]
**RouteType** | **string** | Route type | [default to null]
**Metric** | **string** | Metric value of route | [optional] [default to null]
**Destination** | **string** | Destination covered by route | [optional] [default to null]
**InterfaceId** | **string** | Network interface id of route | [optional] [default to null]
**RouteId** | **string** | Unique identifier for the route | [optional] [default to null]
**Netmask** | **string** | Netmask of destination covered by route | [optional] [default to null]
**Scope** | **string** | Scope of destinations covered by route | [optional] [default to null]
**Gateway** | **string** | Address of next hop | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

