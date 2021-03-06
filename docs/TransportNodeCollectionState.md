# TransportNodeCollectionState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationErrors** | [**[]ValidationError**](ValidationError.md) | Transport node profile(TNP) will not be applied to a discovered node(DN) if some validations are not passed. In this case transport node is not created or existing transport node is not updated with TNP configurations. | [optional] [default to null]
**State** | **string** | If the host preparation or transport node creation is going on for any host then state will be \&quot;IN_PROGRESS\&quot;.  If setting desired state of the transport node failed for any of the host then state will be \&quot;FAILED_TO_CREATE\&quot;  If realization of transport node failed for any of the host then state will be \&quot;FAILED_TO_REALIZE\&quot;  If Transport node is successfully created for all of the hosts in compute collection then state will be \&quot;SUCCESS\&quot;  You can override the configuration for one or more hosts in the compute collection by update TN(transport node) request on individual TN. If TN is successfully created for all hosts in compute collection and one or more hosts have overridden configuration then transport node collection state will be \&quot;PROFILE_MISMATCH\&quot;.  | [optional] [default to null]
**AggregateProgressPercentage** | **int64** | Average of all transport node deployment progress in a cluster. Applicable only if transport node profile is applied on a cluster. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

