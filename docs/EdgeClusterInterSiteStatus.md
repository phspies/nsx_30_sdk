# EdgeClusterInterSiteStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTimestamp** | **int64** | Timestamp when the edge cluster inter-site status was last updated.  | [optional] [default to null]
**OverallStatus** | **string** | Overall status of all edge nodes IBGP status in the edge cluster.  | [optional] [default to null]
**EdgeClusterName** | **string** | Name of the edge cluster whose status is being reported. | [optional] [default to null]
**MemberStatus** | [**[]EdgeClusterMemberInterSiteStatus**](EdgeClusterMemberInterSiteStatus.md) | Per edge node inter-site status. | [optional] [default to null]
**EdgeClusterId** | **string** | Id of the edge cluster whose status is being reported. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

