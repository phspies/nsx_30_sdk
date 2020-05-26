# LogicalRouter

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
**EdgeClusterMemberIndices** | **[]int64** | For stateful services, the logical router should be associated with edge cluster. For TIER 1 logical router, for manual placement of service router within the cluster, edge cluster member indices needs to be provided else same will be auto-allocated. You can provide maximum two indices for HA ACTIVE_STANDBY. For TIER0 logical router this property is no use and placement is derived from logical router uplink or loopback port.  | [optional] [default to null]
**Ipv6Profiles** | [***IPv6Profiles**](IPv6Profiles.md) |  | [optional] [default to null]
**AllocationProfile** | [***EdgeClusterMemberAllocationProfile**](EdgeClusterMemberAllocationProfile.md) |  | [optional] [default to null]
**FirewallSections** | [**[]ResourceReference**](ResourceReference.md) | List of Firewall sections related to Logical Router. | [optional] [default to null]
**FailoverMode** | **string** | Determines the behavior when a logical router instance restarts after a failure. If set to PREEMPTIVE, the preferred node will take over, even if it causes another failure. If set to NON_PREEMPTIVE, then the instance that restarted will remain secondary. This property must not be populated unless the high_availability_mode property is set to ACTIVE_STANDBY. If high_availability_mode property is set to ACTIVE_STANDBY and this property is not specified then default will be NON_PREEMPTIVE.  | [optional] [default to null]
**AdvancedConfig** | [***LogicalRouterConfig**](LogicalRouterConfig.md) |  | [optional] [default to null]
**RouterType** | **string** | TIER0 for external connectivity. TIER1 for two tier topology with TIER0 on top. VRF for isolation of routing table on TIER0.  | [default to null]
**PreferredEdgeClusterMemberIndex** | **int64** | Preferred edge cluster member index which is required for PREEMPTIVE failover mode. Used for Tier0 routers only.  | [optional] [default to null]
**HighAvailabilityMode** | **string** | High availability mode | [optional] [default to null]
**EdgeClusterId** | **string** | Used for tier0 routers | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

