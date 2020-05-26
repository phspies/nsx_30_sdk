# ServiceInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnFailurePolicy** | **string** | Failure policy of the service instance - if it has to be different from the service. By default the service instance inherits the FailurePolicy of the service it belongs to. | [optional] [default to null]
**TransportType** | **string** | Transport to be used by this service instance for deploying the Service-VM. This field is to be set Not Applicable(NA) if the service only caters to functionality EPP(Endpoint Protection). | [default to null]
**ResourceType** | **string** | ServiceInstance is used when NSX handles the lifecyle of   appliance. Deployment and appliance related all the information is necessary. ByodServiceInstance is a custom instance to be used when NSX is not handling   the lifecycles of appliance/s. User will manage their own appliance (BYOD)   to connect with NSX. VirtualServiceInstance is a a custom instance to be used when NSX is not   handling the lifecycle of an appliance and when the user is not bringing   their own appliance.  | [default to null]
**ServiceId** | **string** | The Service to which the service instance is associated. | [optional] [default to null]
**DeploymentSpecName** | **string** | Name of the deployment spec to be used by this service instance. | [default to null]
**InstanceDeploymentTemplate** | [***DeploymentTemplate**](DeploymentTemplate.md) |  | [default to null]
**ImplementationType** | **string** | Implementation to be used by this service instance for deploying the Service-VM. | [default to null]
**AttachmentPoint** | **string** | Attachment point to be used by this service instance for deploying the Service-VM. | [default to null]
**InstanceDeploymentConfig** | [***InstanceDeploymentConfig**](InstanceDeploymentConfig.md) |  | [optional] [default to null]
**DeploymentMode** | **string** | Deployment mode specifies where the partner appliance will be deployed in HA or non-HA i.e standalone mode. | [default to DEPLOYMENT_MODE.ACTIVE_STANDBY]
**DeployedTo** | [**[]ResourceReference**](ResourceReference.md) | List of resource references where service instance be deployed. Ex. Tier 0 Logical Router in case of N-S ServiceInsertion. | [default to null]
**ServiceDeploymentId** | **string** | Id of the Service Deployment using which the instances were deployed. Its available only for instances that were deployed using service deployment API. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

