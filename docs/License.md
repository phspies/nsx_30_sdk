# License

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Features** | **string** | semicolon delimited feature list | [optional] [default to null]
**Description** | **string** | license edition | [optional] [default to null]
**ProductVersion** | **string** | product version | [optional] [default to null]
**Expiry** | **int64** | date that license expires | [optional] [default to null]
**IsEval** | **bool** | true for evalution license | [optional] [default to null]
**IsMh** | **bool** | multi-hypervisor support | [optional] [default to null]
**LicenseKey** | **string** | license key | [optional] [default to null]
**IsExpired** | **bool** | whether the license has expired | [optional] [default to null]
**ProductName** | **string** | product name | [optional] [default to null]
**CapacityType** | **string** | License metrics specifying the capacity type of license key. Types are: - VM - CPU - USER(Concurrent User) - CORE  | [optional] [default to null]
**Quantity** | **int64** | license capacity; 0 for unlimited | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

