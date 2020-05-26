# L2VpnSessionStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficStatisticsPerLogicalSwitch** | [**[]L2VpnPerLsTrafficStatistics**](L2VPNPerLSTrafficStatistics.md) | Traffic statistics per logical switch. | [optional] [default to null]
**DisplayName** | **string** | L2VPN display name. | [optional] [default to null]
**PartialStats** | **bool** | Partial statistics is set to true if onle active node responds while standby does not. In case of both nodes responded statistics will be summed and partial stats will be false. If cluster has only active node, partial statistics will always be false. | [optional] [default to null]
**SessionId** | **string** | Session identifier for L2VPN. | [optional] [default to null]
**TapTrafficCounters** | [**[]L2VpnTapTrafficStatistics**](L2VPNTapTrafficStatistics.md) | Tunnel port traffic counters. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

