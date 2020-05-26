# LogicalSwitch

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
**SwitchType** | **string** | This field indicates purpose of a LogicalSwitch. It is set by manager internally or user can provide this field. If not set, DEFAULT type is assigned. NSX components can use this field to create LogicalSwitch that provides component specific functionality. DEFAULT type LogicalSwitches are created for basic L2 connectivity by API users. SERVICE_PLANE type LogicalSwitches are system created service plane LogicalSwitches for Service Insertion service. User can not create SERVICE_PLANE type of LogicalSwitch. DHCP_RELAY type LogicalSwitches are created by external user like Policy with special permissions or by system and will be treated as internal LogicalSwitches. Such LogicalSwitch will not be exposed to vSphere user. GLOBAL type LogicalSwitches are created to span multiple NSX domains to connect multiple remote sites.  | [optional] [default to null]
**ReplicationMode** | **string** | Replication mode of the Logical Switch | [optional] [default to null]
**ExtraConfigs** | [**[]ExtraConfig**](ExtraConfig.md) | This property could be used for vendor specific configuration in key value string pairs, the setting in extra_configs will be automatically inheritted by logical ports in the logical switch.  | [optional] [default to null]
**UplinkTeamingPolicyName** | **string** | This name has to be one of the switching uplink teaming policy names listed inside the logical switch&#x27;s TransportZone. If this field is not specified, the logical switch will not have a teaming policy associated with it and the host switch&#x27;s default teaming policy will be used. | [optional] [default to null]
**AddressBindings** | [**[]PacketAddressClassifier**](PacketAddressClassifier.md) | Address bindings for the Logical switch | [optional] [default to null]
**IpPoolId** | **string** | IP pool id that associated with a LogicalSwitch. | [optional] [default to null]
**Vlan** | **int64** | This property is dedicated to VLAN based network, to set VLAN of logical network. It is mutually exclusive with &#x27;vlan_trunk_spec&#x27;.  | [optional] [default to null]
**Hybrid** | **bool** | If this flag is set to true, then all the logical switch ports attached to this logical switch will behave in a hybrid fashion. The hybrid logical switch port indicates to NSX that the VM intends to operate in underlay mode, but retains the ability to forward egress traffic to the NSX overlay network. This flag can be enabled only for the logical switches in the overlay type transport zone which has host switch mode as STANDARD and also has either CrossCloud or CloudScope tag scopes. Only the NSX public cloud gateway (PCG) uses this flag, other host agents like ESX, KVM and Edge will ignore it. This property cannot be modified once the logical switch is created.  | [optional] [default to false]
**MacPoolId** | **string** | Mac pool id that associated with a LogicalSwitch. | [optional] [default to null]
**Vni** | **int32** | Only for OVERLAY network. A VNI will be auto-allocated from the default VNI pool if not given; otherwise the given VNI has to be inside the default pool and not used by any other LogicalSwitch.  | [optional] [default to null]
**VlanTrunkSpec** | [***VlanTrunkSpec**](VlanTrunkSpec.md) |  | [optional] [default to null]
**AdminState** | **string** | Represents Desired state of the Logical Switch | [default to null]
**TransportZoneId** | **string** | Id of the TransportZone to which this LogicalSwitch is associated | [default to null]
**Span** | **[]string** | Each manager ID represents the NSX Local Manager the logical switch connects. This will be populated by the manager. | [optional] [default to null]
**SwitchingProfileIds** | [**[]SwitchingProfileTypeIdEntry**](SwitchingProfileTypeIdEntry.md) |  | [optional] [default to null]
**GlobalVni** | **int32** | The VNI is used for intersite traffic and the global logical switch ID. The global VNI pool is agnostic of the local VNI pool, and there is no need to have an exclusive VNI range. For example, VNI x can be the global VNI for logical switch B and the local VNI for logical switch A. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

