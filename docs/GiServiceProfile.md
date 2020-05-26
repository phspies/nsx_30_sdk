# GiServiceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Service Profile type, for example &#x27;GiServiceProfile&#x27;, &#x27;ServiceInsertionServiceProfile&#x27; | [default to null]
**VendorTemplateKey** | **string** | Different VMs in data center can have Different protection levels as specified by administrator in the policy. The identifier for the policy with which the partner appliance identifies this policy. This identifier will be passed to the partner appliance at runtime to specify which protection level is applicable for the VM being protected. | [optional] [default to null]
**ServiceId** | **string** | The service to which the service profile belongs. | [optional] [default to null]
**VendorTemplateId** | **string** | ID of the vendor template, created by partner while registering the service. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

