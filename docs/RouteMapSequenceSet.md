# RouteMapSequenceSet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferGlobalV6NextHop** | **bool** | For incoming and import route_maps on receiving both v6 global and v6 link-local address for the route, prefer to use the global address as the next hop. By default, it prefers the link-local next hop.  | [optional] [default to false]
**LocalPreference** | **int64** | Local preference indicates the degree of preference for one BGP route over other BGP routes. The path/route with highest local preference value is preferred/selected. If local preference value is not specified then it will be considered as 100 by default.  | [optional] [default to null]
**Weight** | **int32** | Weight used to select certain path | [optional] [default to null]
**LargeCommunity** | **string** | Set large BGP community, community value shoud be in aa:bb:nn format where aa, bb, nn are unsigned integers with range [1-4294967295]. | [optional] [default to null]
**AsPathPrepend** | **string** | As Path Prepending to influence path selection | [optional] [default to null]
**Community** | **string** | Set normal BGP community either well-known community name or community value in aa:nn(2byte:2byte) format.  | [optional] [default to null]
**MultiExitDiscriminator** | **int64** | Multi Exit Discriminator (MED) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

