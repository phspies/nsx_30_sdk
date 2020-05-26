# DuplicateAddressBindingEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Source from which the address binding entry was obtained | [optional] [default to SOURCE.UNKNOWN]
**Binding** | [***PacketAddressClassifier**](PacketAddressClassifier.md) |  | [optional] [default to null]
**BindingTimestamp** | **int64** | Timestamp at which the binding was discovered via snooping or manually specified by the user  | [optional] [default to null]
**ConflictingPort** | **string** | Provides the ID of the port on which the same address bidning exists  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

