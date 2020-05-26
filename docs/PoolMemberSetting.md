# PoolMemberSetting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | **string** | Member admin state | [optional] [default to ADMIN_STATE.ENABLED]
**IpAddress** | **string** | Pool member IP address | [default to null]
**Port** | **string** | Pool member port number | [optional] [default to null]
**Weight** | **int64** | Only applicable to static pool members. If supplied for a pool defined by a grouping object, update API would fail.  | [optional] [default to null]
**DisplayName** | **string** | Only applicable to static pool members. If supplied for a pool defined by a grouping object, update API would fail.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

