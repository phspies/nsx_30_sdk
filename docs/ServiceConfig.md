# ServiceConfig

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
**AppliedTo** | [**[]ResourceReference**](ResourceReference.md) | The list of entities that the configurations should be applied to. This can either be a NSGroup or any other entity like TransportNode, LogicalPorts etc.  | [optional] [default to null]
**Precedence** | **int64** | Every ServiceConfig has a priority based upon its precedence value. Lower the value of precedence, higher will be its priority. If user doesnt specify the precedence, it is generated automatically by system. The precedence is generated based upon the type of profile used in ServiceConfig. Precedence are auto-generated in decreasing order with difference of 100. Automatically generated precedence value will be 100 less than the current minimum value of precedence of ServiceConfig of a given profile type in system.There cannot be duplicate precedence for ServiceConfig of same profile type.  | [optional] [default to null]
**Profiles** | [**[]NsxProfileReference**](NSXProfileReference.md) | These are the NSX Profiles which will be added to service config, which will be applied to entities/groups provided to applied_to field of service config.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

