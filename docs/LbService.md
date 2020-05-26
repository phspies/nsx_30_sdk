# LbService

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
**AccessLogEnabled** | **bool** | Whether access log is enabled | [optional] [default to null]
**Attachment** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**ErrorLogLevel** | **string** | Load balancer engine writes information about encountered issues of different severity levels to the error log. This setting is used to define the severity level of the error log.  | [optional] [default to ERROR_LOG_LEVEL.INFO]
**VirtualServerIds** | **[]string** | virtual servers can be associated to LbService(which is similar to physical/virtual load balancer), Lb virtual servers, pools and other entities could be defined independently, the virtual server identifier list here would be used to maintain the relationship of LbService and other Lb entities.  | [optional] [default to null]
**RelaxScaleValidation** | **bool** | If relax_scale_validation is true, the scale validations for virtual servers/pools/pool members/rules are relaxed for load balancer service. When load balancer service is deployed on edge nodes, the scale of virtual servers/pools/pool members for the load balancer service should not exceed the scale number of the largest load balancer size which could be configured on a certain edge form factor. For example, the largest load balancer size supported on a MEDIUM edge node is MEDIUM. So one SMALL load balancer deployed on MEDIUM edge nodes can support the scale number of MEDIUM load balancer. It is not recommended to enable active monitors if relax_scale_validation is true due to performance consideration. If relax_scale_validation is false, scale numbers should be validated for load balancer service.  | [optional] [default to false]
**Enabled** | **bool** | Whether the load balancer service is enabled | [optional] [default to true]
**Size** | **string** | The size of load balancer service | [optional] [default to SIZE.SMALL]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

