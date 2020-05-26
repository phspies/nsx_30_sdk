# LbIcmpMonitor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorPort** | **string** | If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported. For ICMP monitor, monitor_port is not required.  | [optional] [default to null]
**FallCount** | **int64** | num of consecutive checks must fail before marking it down | [optional] [default to 3]
**Interval** | **int64** | the frequency at which the system issues the monitor check (in second) | [optional] [default to 5]
**RiseCount** | **int64** | num of consecutive checks must pass before marking it up | [optional] [default to 3]
**Timeout** | **int64** | the number of seconds the target has in which to respond to the monitor request  | [optional] [default to 15]
**DataLength** | **int64** | The data size(in byte) of the ICMP healthcheck packet | [optional] [default to 56]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

