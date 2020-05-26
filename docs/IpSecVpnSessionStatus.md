# IpSecVpnSessionStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpsecVpnSessionId** | **string** | UUID of vpn session. | [optional] [default to null]
**DisplayName** | **string** | Display name of vpn session. | [optional] [default to null]
**FailedTunnels** | **int64** | Number of failed tunnels. | [optional] [default to null]
**NegotiatedTunnels** | **int64** | Number of negotiated tunnels. | [optional] [default to null]
**SessionStatus** | **string** | Gives session status consolidated using IKE status and tunnel status. It can be UP, DOWN, DEGRADED. If IKE and all tunnels are UP status will be UP, if all down it will be DOWN, otherwise it will be DEGRADED. | [optional] [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated. | [optional] [default to null]
**AggregateTrafficCounters** | [***IpSecVpnTrafficCounters**](IPSecVPNTrafficCounters.md) |  | [optional] [default to null]
**IkeStatus** | [***IpSecVpnikeSessionStatus**](IPSecVPNIKESessionStatus.md) |  | [optional] [default to null]
**TotalTunnels** | **int64** | Total number of tunnels. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

