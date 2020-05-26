# HostNodeLoginCredential

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username of the account on the host node | [optional] [default to null]
**Password** | **string** | The authentication password of the host node | [optional] [default to null]
**Thumbprint** | **string** | For ESXi hosts, the thumbprint of the ESXi management service. For KVM hosts, the SSH key fingerprint. If thumbprint is not provided then connection to host may not be established and API call will fail.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

