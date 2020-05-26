package nsxt

type LbActiveMonitor struct {
	// Load balancers monitor the health of backend servers to ensure traffic is not black holed. There are two types of healthchecks: active and passive. Passive healthchecks depend on failures in actual client traffic (e.g. RST from server in response to a client connection) to detect that the server or the application is down. In case of active healthchecks, load balancer itself initiates new connections (or sends ICMP ping) to the servers periodically to check their health, completely independent of any data traffic. Currently, active health monitors are supported for HTTP, HTTPS, TCP, UDP and ICMP protocols. 
	ResourceType string `json:"resource_type"`
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
}
