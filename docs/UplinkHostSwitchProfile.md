# UplinkHostSwitchProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredCapabilities** | **[]string** |  | [optional] [default to null]
**ResourceType** | **string** | Supported HostSwitch profiles. | [default to null]
**Lags** | [**[]Lag**](Lag.md) | list of LACP group | [optional] [default to null]
**TransportVlan** | **int64** | VLAN used for tagging Overlay traffic of associated HostSwitch | [optional] [default to 0]
**Teaming** | [***TeamingPolicy**](TeamingPolicy.md) |  | [default to null]
**OverlayEncap** | **string** | The protocol used to encapsulate overlay traffic | [optional] [default to OVERLAY_ENCAP.GENEVE]
**NamedTeamings** | [**[]NamedTeamingPolicy**](NamedTeamingPolicy.md) | List of named uplink teaming policies that can be used by logical switches | [optional] [default to null]
**Mtu** | **int32** | Maximum Transmission Unit used for uplinks | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

