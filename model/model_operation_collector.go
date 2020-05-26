package nsxt

// The operation collector is defined to receive stats from hosts. (eg. vRNI-collector collects all the system metrics)
type OperationCollector struct {
	// Port for the operation collector.
	CollectorPort int32 `json:"collector_port"`
	// IP address for the operation collector.
	CollectorIp string `json:"collector_ip"`
}
