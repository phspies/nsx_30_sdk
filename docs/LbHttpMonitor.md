# LbHttpMonitor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorPort** | **string** | If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported. For ICMP monitor, monitor_port is not required.  | [optional] [default to null]
**FallCount** | **int64** | num of consecutive checks must fail before marking it down | [optional] [default to 3]
**Interval** | **int64** | the frequency at which the system issues the monitor check (in second) | [optional] [default to 5]
**RiseCount** | **int64** | num of consecutive checks must pass before marking it up | [optional] [default to 3]
**Timeout** | **int64** | the number of seconds the target has in which to respond to the monitor request  | [optional] [default to 15]
**ResponseStatusCodes** | **[]int32** | The HTTP response status code should be a valid HTTP status code.  | [optional] [default to null]
**RequestMethod** | **string** | the health check method for HTTP monitor type | [optional] [default to REQUEST_METHOD.GET]
**RequestBody** | **string** | String to send as part of HTTP health check request body. Valid only for certain HTTP methods like POST.  | [optional] [default to null]
**ResponseBody** | **string** | If HTTP response body match string (regular expressions not supported) is specified (using LbHttpMonitor.response_body) then the healthcheck HTTP response body is matched against the specified string and server is considered healthy only if there is a match. If the response body string is not specified, HTTP healthcheck is considered successful if the HTTP response status code is 2xx, but it can be configured to accept other status codes as successful.  | [optional] [default to null]
**RequestUrl** | **string** | URL used for HTTP monitor | [optional] [default to null]
**RequestVersion** | **string** | HTTP request version | [optional] [default to REQUEST_VERSION.11_]
**RequestHeaders** | [**[]LbHttpRequestHeader**](LbHttpRequestHeader.md) | Array of HTTP request headers | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

