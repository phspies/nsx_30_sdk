# RaConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HopLimit** | **int64** | The maximum number of hops through which packets can pass before being discarded.  | [optional] [default to 64]
**RouterLifetime** | **int64** | Router lifetime value in seconds. A value of 0 indicates the router is not a default router for the receiving end. Any other value in this field specifies the lifetime, in seconds, associated with this router as a default router.  | [optional] [default to 1800]
**RaInterval** | **int64** | Interval between 2 Router advertisement in seconds.  | [optional] [default to 600]
**PrefixPreferredTime** | **int64** | The time interval in seconds, in which the prefix is advertised as preferred.  | [optional] [default to 604800]
**PrefixLifetime** | **int64** | The time interval in seconds, in which the prefix is advertised as valid.  | [optional] [default to 2592000]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

