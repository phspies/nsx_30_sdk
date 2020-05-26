# ServiceDeploymentSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentSpecs** | [**[]SvmDeploymentSpec**](SVMDeploymentSpec.md) | Deployment Specs holds information required to deploy the Service-VMs. i.e. OVF url where the partner Service-VM OVF is hosted. The host type on which the OVF can be deployed, Form factor to name a few. | [optional] [default to null]
**NicMetadataList** | [**[]NicMetadata**](NicMetadata.md) | NIC metadata associated with the deployment spec. | [optional] [default to null]
**DeploymentTemplate** | [**[]DeploymentTemplate**](DeploymentTemplate.md) | Deployment Template holds the attributes specific to partner for which the service is created. These attributes are opaque to NSX Manager. | [default to null]
**SvmVersion** | **string** | Partner needs to specify the Service VM version which will get deployed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

