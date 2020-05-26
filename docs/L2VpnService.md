# L2VpnService

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
**EnableFullMesh** | **bool** | Full mesh topology auto disables traffic replication between connected peers. However, this property is deprecated. Please refer enable_hub property instead to control client to client forwarding via the server. The value of enable_full_mesh will not be used anymore. If enable_hub is not provided explicitly, the default value of it will be used.  | [optional] [default to false]
**EnableHub** | **bool** | This property only applies in SERVER mode. If set to true, traffic from any client will be replicated to all other clients. If set to false, traffic received from clients is only replicated to the local VPN endpoint.  | [optional] [default to false]
**LogicalRouterId** | **string** | Logical router id | [default to null]
**Mode** | **string** | Specify an L2VPN service mode as SERVER or CLIENT. L2VPN service in SERVER mode requires user to configure L2VPN session explicitly. L2VPN service in CLIENT mode can use peercode generated from SERVER to configure L2VPN session.  | [optional] [default to MODE.SERVER]
**LogicalTapIpPool** | **[]string** | IP Pool to allocate local and peer endpoint IPs for L2VpnSession logical Tap. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

