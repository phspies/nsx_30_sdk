# TransportNode

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
**HostSwitchSpec** | [***HostSwitchSpec**](HostSwitchSpec.md) |  | [optional] [default to null]
**NodeId** | **string** | Unique Id of the fabric node | [optional] [default to null]
**NodeDeploymentInfo** | [***Node**](Node.md) |  | [optional] [default to null]
**MaintenanceMode** | **string** | The property is read-only, used for querying result. User could update transport node maintenance mode by UpdateTransportNodeMaintenanceMode call. | [optional] [default to null]
**FailureDomainId** | **string** | Set failure domain of edge transport node which will help in auto placement of TIER1 logical routers, DHCP Servers and MDProxies, if failure domain based allocation is enabled in edge cluster API. It is only supported for edge transport node and not for host transport node. In case failure domain is not set by user explicitly, it will be always assigned with default system created failure domain.  | [optional] [default to null]
**RemoteTunnelEndpoint** | [***TransportNodeRemoteTunnelEndpointConfig**](TransportNodeRemoteTunnelEndpointConfig.md) |  | [optional] [default to null]
**IsOverridden** | **bool** | This flag is relevant to only those hosts which are part of a compute collection which has transport node profile (TNP) applied on it. If you change the transport node configuration and it is different than cluster level TNP then this flag will be set to true  | [optional] [default to null]
**TransportZoneEndpoints** | [**[]TransportZoneEndPoint**](TransportZoneEndPoint.md) | This is deprecated. TransportZoneEndPoints should be specified per host switch at StandardHostSwitch through Transport Node or Transport Node Profile configuration. This will ONLY include the TransportZoneEndpoints that were were specified here during the TransportNode configuration. If TransportZoneEndpoints are specified directly in {$ref: StandardHostSwitch}, such TransportZoneEndpoints will not be populated here.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

