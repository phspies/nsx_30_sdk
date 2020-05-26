package nsxt

type LbHttpMonitor struct {
	// If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported. For ICMP monitor, monitor_port is not required. 
	MonitorPort string `json:"monitor_port,omitempty"`
	// num of consecutive checks must fail before marking it down
	FallCount int64 `json:"fall_count,omitempty"`
	// the frequency at which the system issues the monitor check (in second)
	Interval int64 `json:"interval,omitempty"`
	// num of consecutive checks must pass before marking it up
	RiseCount int64 `json:"rise_count,omitempty"`
	// the number of seconds the target has in which to respond to the monitor request 
	Timeout int64 `json:"timeout,omitempty"`
	// The HTTP response status code should be a valid HTTP status code. 
	ResponseStatusCodes []int32 `json:"response_status_codes,omitempty"`
	// the health check method for HTTP monitor type
	RequestMethod string `json:"request_method,omitempty"`
	// String to send as part of HTTP health check request body. Valid only for certain HTTP methods like POST. 
	RequestBody string `json:"request_body,omitempty"`
	// If HTTP response body match string (regular expressions not supported) is specified (using LbHttpMonitor.response_body) then the healthcheck HTTP response body is matched against the specified string and server is considered healthy only if there is a match. If the response body string is not specified, HTTP healthcheck is considered successful if the HTTP response status code is 2xx, but it can be configured to accept other status codes as successful. 
	ResponseBody string `json:"response_body,omitempty"`
	// URL used for HTTP monitor
	RequestUrl string `json:"request_url,omitempty"`
	// HTTP request version
	RequestVersion string `json:"request_version,omitempty"`
	// Array of HTTP request headers
	RequestHeaders []LbHttpRequestHeader `json:"request_headers,omitempty"`
}
