# SwitchingToVmcModeParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModeId** | **string** | Possible enum values in a \&quot;/config/nsx_appliance_mode\&quot; file | [default to null]
**SreOrg** | [***OrgInfo**](OrgInfo.md) |  | [optional] [default to null]
**DefaultOrgId** | **string** | Org ID of a Client - commonly UUID. | [optional] [default to null]
**EaOrg** | [***OrgInfo**](OrgInfo.md) |  | [optional] [default to null]
**GssOrg** | [***OrgInfo**](OrgInfo.md) |  | [optional] [default to null]
**ProxyHost** | **string** | IP/host of PoP (Point-of-Presence) HTTP proxy server | [optional] [default to null]
**CspTimeDrift** | **int64** | CSP time drift in milliseconds | [optional] [default to null]
**SddcId** | **string** | SDDC id | [optional] [default to null]
**BasicAuthWhitelistIps** | **[]string** | List of whitelist IPs for basic auth | [optional] [default to null]
**BaseUrl** | **string** | Protocol and domain name (or IP address) of a CSP server, like \&quot;https://console-stg.cloud.vmware.com\&quot;. | [optional] [default to null]
**ProxyPort** | **int64** | Port of PoP (Point-of-Presence) Http proxy server | [optional] [default to null]
**CspOrgUri** | **string** | Relative path on CSP server to the Org location. Can be \&quot;/csp/gateway/am/api/orgs/\&quot;. | [optional] [default to null]
**CspClientCredential** | [***Oauth2Credentials**](Oauth2Credentials.md) |  | [optional] [default to null]
**AuthCode** | [***Oauth2Credentials**](Oauth2Credentials.md) |  | [optional] [default to null]
**ModeChangeOnly** | **bool** | When this parameter is set to true, only a change of the node mode happens without any update to the auth properties. When this param is not set to true i.e. set to false or not provided, mode change and update to the auth properties will both happen. | [optional] [default to null]
**CspClientIncomingCredentials** | **[]string** | List of incoming client IDs | [optional] [default to null]
**ServiceDefinitionId** | **string** | Service definition id | [optional] [default to null]
**ResourceType** | **string** | Node Mode type | [optional] [default to RESOURCE_TYPE.SWITCHING_TO_VMC_MODE_PARAMETERS]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

