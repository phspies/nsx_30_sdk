# LogicalPort

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
**LogicalSwitchId** | **string** | Id of the Logical switch that this port belongs to. | [default to null]
**InitState** | **string** | Set initial state when a new logical port is created. &#x27;UNBLOCKED_VLAN&#x27; means new port will be unblocked on traffic in creation, also VLAN will be set with corresponding logical switch setting. This port setting can only be configured at port creation (POST), and cannot be modified.  | [optional] [default to null]
**SwitchingProfileIds** | [**[]SwitchingProfileTypeIdEntry**](SwitchingProfileTypeIdEntry.md) |  | [optional] [default to null]
**Attachment** | [***LogicalPortAttachment**](LogicalPortAttachment.md) |  | [optional] [default to null]
**InternalId** | **string** | The internal_id of the logical port may or may not be identical to it&#x27;s managed resource ID. If a VirtualMachine connected to logical port migrates from one site to another, then on the destination site, it will be connected to different logical port managed resource. However, the internal_id field will be persisted across vmotion.  | [optional] [default to null]
**ExtraConfigs** | [**[]ExtraConfig**](ExtraConfig.md) | This property could be used for vendor specific configuration in key value string pairs. Logical port setting will override logical switch setting if the same key was set on both logical switch and logical port.  | [optional] [default to null]
**AddressBindings** | [**[]PacketAddressClassifier**](PacketAddressClassifier.md) | Each address binding must contain both an IPElement and MAC address. VLAN ID is optional. This binding configuration can be used by features such as spoof-guard and overrides any discovered bindings. Any non unique entries are deduplicated to generate a unique set of address bindings and then stored. For IPv6 addresses, a subnet address cannot have host bits set. A maximum of 128 unique address bindings is allowed per port.  | [optional] [default to null]
**IgnoreAddressBindings** | [**[]PacketAddressClassifier**](PacketAddressClassifier.md) | IP Discovery module uses various mechanisms to discover address bindings being used on each port. If a user would like to ignore any specific discovered address bindings or prevent the discovery of a particular set of discovered bindings, then those address bindings can be provided here. Currently IP range in CIDR format is not supported.  | [optional] [default to null]
**AdminState** | **string** | Represents Desired state of the logical port | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

