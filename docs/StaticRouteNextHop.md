# StaticRouteNextHop

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackholeAction** | **string** | Action to be taken on matching packets for NULL routes. | [optional] [default to null]
**AdministrativeDistance** | **int64** | Administrative Distance for the next hop IP | [optional] [default to 1]
**IpAddress** | **string** | Next Hop IP | [optional] [default to null]
**BfdEnabled** | **bool** | Status of bfd for this next hop where bfd_enabled &#x3D; true indicate bfd is enabled for this next hop and bfd_enabled &#x3D; false indicate bfd peer is disabled or not configured for this next hop. | [optional] [default to false]
**LogicalRouterPortId** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

