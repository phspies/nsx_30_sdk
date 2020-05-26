# IpSecVpnPolicyRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | [***OwnerResourceLink**](OwnerResourceLink.md) |  | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Id** | **string** | Unique policy id. | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Sources** | [**[]IpSecVpnPolicySubnet**](IPSecVPNPolicySubnet.md) | List of local subnets. | [optional] [default to null]
**Action** | **string** | PROTECT - Protect rules are defined per policy based IPSec VPN session. BYPASS - Bypass rules are defined per IPSec VPN service and affects all policy based IPSec VPN sessions. Bypass rules are prioritized over protect rules.  | [optional] [default to ACTION.PROTECT]
**Enabled** | **bool** | A flag to enable/disable the policy rule. | [optional] [default to true]
**Logged** | **bool** | A flag to enable/disable the logging for the policy rule. | [optional] [default to false]
**Destinations** | [**[]IpSecVpnPolicySubnet**](IPSecVPNPolicySubnet.md) | List of peer subnets. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

