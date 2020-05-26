# InstanceRuntime

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
**ServiceVmId** | **string** | Service-VM/SVM id of deployed virtual-machine. | [optional] [default to null]
**DeploymentStatus** | **string** | Service-Instance Runtime deployment status of the Service-VM. It shows the latest status during the process of deployment, redeploy, upgrade, and un-deployment of VM. | [optional] [default to null]
**VmNicInfo** | [***VmNicInfo**](VmNicInfo.md) |  | [optional] [default to null]
**MaintenanceMode** | **string** | The maintenance mode indicates whether the corresponding service VM is in maintenance mode. The service VM will not be used to service new requests if it is in maintenance mode.  | [optional] [default to null]
**RuntimeStatus** | **string** | Service-Instance Runtime status of the deployed Service-VM. | [optional] [default to null]
**ErrorMessage** | **string** | Error message for the Service Instance Runtime if any. | [optional] [default to null]
**ServiceInstanceId** | **string** | Id of an instantiation of a registered service. | [optional] [default to null]
**RuntimeHealthStatusByPartner** | **string** | Service-Instance runtime health status set by partner to indicate whether the service is running properly or not.  | [optional] [default to null]
**UnhealthyReason** | **string** | Reason provided by partner for the service being unhealthy. This could be due to various reasons such as connectivity lost as an example.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

