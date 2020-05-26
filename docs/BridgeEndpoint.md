# BridgeEndpoint

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
**HaEnable** | **bool** | This field will not be used if an edge cluster is being used for the bridge endpoint  | [optional] [default to true]
**BridgeClusterId** | **string** | This field will not be used if an edge cluster is being used for the bridge endpoint  | [optional] [default to null]
**VlanTransportZoneId** | **string** | This field will not be used if a bridge cluster is being used for the bridge endpoint  | [optional] [default to null]
**BridgeEndpointProfileId** | **string** | This field will not be used if a bridge cluster is being used for the bridge endpoint  | [optional] [default to null]
**UplinkTeamingPolicyName** | **string** | This name has to be one of the switching uplink teaming policy names listed inside the TransportZone. If this field is not specified, bridge will use the first pnic in host-switch config. This field will not be used if a bridge cluster is being used for the bridge endpoint | [optional] [default to null]
**VlanTrunkSpec** | [***VlanTrunkSpec**](VlanTrunkSpec.md) |  | [optional] [default to null]
**Vlan** | **int64** | This property is used for VLAN specification of bridge endpoint. It&#x27;s mutually exclusive with &#x27;vlan_trunk_spec&#x27;, either &#x27;vlan&#x27; or &#x27;vlan_trunk_spec&#x27; should be specified.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

