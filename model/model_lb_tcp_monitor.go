package nsxt

type LbTcpMonitor struct {
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
	// Expected data, if specified, can be anywhere in the response and it has to be a string, regular expressions are not supported. 
	Receive string `json:"receive,omitempty"`
	// If both send and receive are not specified, then just a TCP connection is established (3-way handshake) to validate server is healthy, no data is sent. 
	Send string `json:"send,omitempty"`
}
