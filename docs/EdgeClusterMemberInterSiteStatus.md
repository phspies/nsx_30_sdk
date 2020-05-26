# EdgeClusterMemberInterSiteStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportNode** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**EstablishedBgpSessions** | **int64** | Total number of current established inter-site IBGP sessions. | [optional] [default to null]
**Status** | **string** | Edge node IBGP status | [optional] [default to null]
**TotalBgpSessions** | **int64** | Total number of inter-site IBGP sessions. | [optional] [default to null]
**NeighborStatus** | [**[]BgpNeighborStatusLiteDto**](BgpNeighborStatusLiteDto.md) | Inter-site BGP neighbor status. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

