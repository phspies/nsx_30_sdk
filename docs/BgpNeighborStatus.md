# BgpNeighborStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionState** | **string** | Current state of the BGP session. | [optional] [default to null]
**MessagesReceived** | **int64** | Count of messages received from the neighbor | [optional] [default to null]
**KeepAliveInterval** | **int64** | Time in ms to wait for HELLO packet from BGP peer | [optional] [default to null]
**NeighborRouterId** | **string** | Router ID of the BGP neighbor. | [optional] [default to null]
**TotalOutPrefixCount** | **int64** | Sum of out prefixes counts across all address families. | [optional] [default to null]
**LrComponentId** | **string** | Logical router component(Service Router/Distributed Router) id | [optional] [default to null]
**EstablishedConnectionCount** | **int64** | Count of connections established | [optional] [default to null]
**MessagesSent** | **int64** | Count of messages sent to the neighbor | [optional] [default to null]
**TimeSinceEstablished** | **int64** | Time(in milliseconds) since connection was established. | [optional] [default to null]
**HoldTime** | **int64** | Time in ms to wait for HELLO from BGP peer. If a HELLO packet is not seen from BGP Peer withing hold_time then BGP neighbor will be marked as down. | [optional] [default to null]
**GracefulRestart** | **bool** | Current state of graceful restart where graceful_restart &#x3D; true indicates graceful restart is enabled and graceful_restart &#x3D; false indicates graceful restart is disabled. This is deprecated field, use graceful_restart_mode instead.  | [optional] [default to null]
**GracefulRestartMode** | **string** | Current state of graceful restart of BGP neighbor. Possible values are - 1. GR_AND_HELPER - Graceful restart with Helper 2. HELPER_ONLY - Helper only 3. DISABLE - Disabled  | [optional] [default to null]
**ConnectionDropCount** | **int64** | Count of connection drop | [optional] [default to null]
**RemotePort** | **int64** | TCP port number of remote BGP Connection | [optional] [default to null]
**TotalInPrefixCount** | **int64** | Sum of in prefixes counts across all address families. | [optional] [default to null]
**RemoteSite** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**TransportNode** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**LocalPort** | **int64** | TCP port number of Local BGP connection | [optional] [default to null]
**RemoteAsNumber** | **string** | AS number of the BGP neighbor | [optional] [default to null]
**AnnouncedCapabilities** | **[]string** | BGP capabilities sent to BGP neighbor. | [optional] [default to null]
**NegotiatedCapability** | **[]string** | BGP capabilities negotiated with BGP neighbor. | [optional] [default to null]
**AddressFamilies** | [**[]BgpAddressFamily**](BgpAddressFamily.md) | Address families of BGP neighbor | [optional] [default to null]
**SourceAddress** | **string** | The Ip address of logical port | [optional] [default to null]
**NeighborAddress** | **string** | The IP of the BGP neighbor | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

