# BfdHealthMonitoringProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**ProbeInterval** | **int64** | The time interval (in millisec) between probe packets for tunnels between transport nodes. | [optional] [default to 1000]
**LatencyEnabled** | **bool** | The flag is to turn on/off latency. A POST or PUT request with \&quot;latency_enabled\&quot; true will enable NSX to send the networking latency data to thrid-party monitoring tools like vRNI. | [optional] [default to null]
**Enabled** | **bool** | Whether the heartbeat is enabled. A POST or PUT request with \&quot;enabled\&quot; false (with no probe intervals) will set (POST) or reset (PUT) the probe_interval to their default value. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

