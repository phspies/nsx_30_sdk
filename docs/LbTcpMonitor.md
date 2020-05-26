# LbTcpMonitor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorPort** | **string** | If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported. For ICMP monitor, monitor_port is not required.  | [optional] [default to null]
**FallCount** | **int64** | num of consecutive checks must fail before marking it down | [optional] [default to 3]
**Interval** | **int64** | the frequency at which the system issues the monitor check (in second) | [optional] [default to 5]
**RiseCount** | **int64** | num of consecutive checks must pass before marking it up | [optional] [default to 3]
**Timeout** | **int64** | the number of seconds the target has in which to respond to the monitor request  | [optional] [default to 15]
**Receive** | **string** | Expected data, if specified, can be anywhere in the response and it has to be a string, regular expressions are not supported.  | [optional] [default to null]
**Send** | **string** | If both send and receive are not specified, then just a TCP connection is established (3-way handshake) to validate server is healthy, no data is sent.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

