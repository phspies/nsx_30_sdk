# ServiceDeployment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | Unique identifier of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**Perimeter** | **string** | This indicates the deployment perimeter, such as a VC cluster or a host. | [optional] [default to PERIMETER.HOST]
**DeploymentSpecName** | **string** | Name of the deployment spec to be used for deployment, which specifies the OVF provided by the partner and the form factor. | [default to null]
**DeploymentMode** | **string** | Mode of deployment. Currently, only stand alone deployment is supported. It is a single VM deployed through this deployment spec. In future, HA configurations will be supported here. | [optional] [default to DEPLOYMENT_MODE.STAND_ALONE]
**InstanceDeploymentTemplate** | [***DeploymentTemplate**](DeploymentTemplate.md) |  | [default to null]
**ServiceDeploymentConfig** | [***ServiceDeploymentConfig**](ServiceDeploymentConfig.md) |  | [default to null]
**ServiceId** | **string** | The Service to which the service deployment is associated. | [optional] [default to null]
**ClusteredDeploymentCount** | **int64** | Number of instances in case of clustered deployment. | [optional] [default to 1]
**DeployedTo** | [**[]ResourceReference**](ResourceReference.md) | List of resource references where service instance be deployed. Ex. Tier 0 Logical Router in case of N-S ServiceInsertion. Service Attachment in case of E-W ServiceInsertion. | [optional] [default to null]
**DeploymentType** | **string** | Specifies whether the service VM should be deployed on each host such that it provides partner service locally on the host, or whether the service VMs can be deployed as a cluster. If deployment_type is CLUSTERED, then the clustered_deployment_count should be provided. | [optional] [default to DEPLOYMENT_TYPE.CLUSTERED]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

