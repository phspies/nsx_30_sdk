# BgpConfig

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
**InterSrIbgp** | [***InterSrRoutingConfig**](InterSRRoutingConfig.md) |  | [optional] [default to null]
**AsNumber** | **int64** | This is a deprecated property, Please use &#x27;as_num&#x27; instead. For VRF logical router, the as_number from parent logical router will be effective. | [optional] [default to null]
**RouteAggregation** | [**[]BgpRouteAggregation**](BgpRouteAggregation.md) | List of routes to be aggregated | [optional] [default to null]
**LogicalRouterId** | **string** | Logical router id | [optional] [default to null]
**GracefulRestart** | **bool** | Flag to enable graceful restart. This field is deprecated, kindly use graceful_restart_config parameter for graceful restart configuration. If both parameters are set and consistent with each other [i.e. graceful_restart&#x3D;false and graceful_restart_mode&#x3D;HELPER_ONLY OR graceful_restart&#x3D;true and graceful_restart_mode&#x3D;GR_AND_HELPER] then this is allowed, but if inconsistent with each other then this is not allowed and validation error will be thrown. For VRF logical router, the settings from parent logical router will be effective.  | [optional] [default to null]
**AsNum** | **string** | For VRF logical router, the as_num from parent logical router will be effective. | [optional] [default to null]
**Enabled** | **bool** | While creation of BGP config this flag will be set to - true for Tier0 logical router with Active-Active high-availability mode - false for Tier0 logical router with Active-Standby high-availanility mode. User can change this value while updating the config. If this property is not specified in the payload, the default value will be considered as false irrespective of the high-availability mode.  | [optional] [default to false]
**GracefulRestartConfig** | [***GracefulRestartConfig**](GracefulRestartConfig.md) |  | [optional] [default to null]
**MultipathRelax** | **bool** | For TIER0 logical router, default is true. For VRF logical router, the settings from parent logical router will be effective. | [optional] [default to null]
**Ecmp** | **bool** | While creation of BGP config this flag will be set to true User can change this value while updating BGP config. If this property is not specified in the payload, the default value will be considered as true.  | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

