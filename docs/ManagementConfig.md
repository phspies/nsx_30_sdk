# ManagementConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** | The _revision property describes the current revision of the resource. To prevent clients from overwriting each other&#x27;s changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. | [optional] [default to null]
**PublishFqdns** | **bool** | True if Management nodes publish their fqdns(instead of default IP addresses) across NSX for its reachability. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

