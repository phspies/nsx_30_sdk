# IpSecVpnSessionStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IkeTrafficStatistics** | [***IpSecVpnikeTrafficStatistics**](IPSecVPNIKETrafficStatistics.md) |  | [optional] [default to null]
**DisplayName** | **string** | Display name of vpn session. | [optional] [default to null]
**PolicyStatistics** | [**[]IpSecVpnPolicyTrafficStatistics**](IPSecVPNPolicyTrafficStatistics.md) | Gives aggregate traffic statistics across all ipsec tunnels and individual tunnel statistics. | [optional] [default to null]
**PartialStats** | **bool** | Partial statistics if true specifies that the statistics are only from active node. | [optional] [default to null]
**IpsecVpnSessionId** | **string** | UUID of vpn session. | [optional] [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated. | [optional] [default to null]
**IkeStatus** | [***IpSecVpnikeSessionStatus**](IPSecVPNIKESessionStatus.md) |  | [optional] [default to null]
**AggregateTrafficCounters** | [***IpSecVpnTrafficCounters**](IPSecVPNTrafficCounters.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

