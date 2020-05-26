package nsxt

type LbUdpMonitor struct {
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
	// Expected data, can be anywhere in the response and it has to be a string, regular expressions are not supported. UDP healthcheck is considered failed if there is no server response within the timeout period. 
	Receive string `json:"receive"`
	// The data to be sent to the monitored server. 
	Send string `json:"send"`
}
