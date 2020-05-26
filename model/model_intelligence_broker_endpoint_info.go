package nsxt

// An endpoint to connect to NSX-Intelligence broker. Either FQDN or IP address can be used in the endpoint info. 
type IntelligenceBrokerEndpointInfo struct {
	// The port number where the broker is listening to. 
	Port int64 `json:"port"`
	// The IP address or the full qualified domain name of broker. 
	Address string `json:"address"`
}
