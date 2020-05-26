# SvmDeploymentSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OvfUrl** | **string** | Location of the partner VM OVF to be deployed. | [default to null]
**Name** | **string** | Deployment Spec name for ease of use, since multiple DeploymentSpec can be specified. | [optional] [default to null]
**MinHostVersion** | **string** | Minimum host version supported by this ovf. If a host in the deployment cluster is having version less than this, then service deployment will not happen on that host. | [optional] [default to 6.5]
**ServiceFormFactor** | **string** | Supported ServiceInsertion Form Factor for the OVF deployment. The default FormFactor is Medium. | [optional] [default to SERVICE_FORM_FACTOR.MEDIUM]
**HostType** | **string** | Host Type on which the specified OVF can be deployed. | [default to null]
**SvmVersion** | **string** | Partner needs to specify the Service VM version which will get deployed. | [optional] [default to 1.0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

