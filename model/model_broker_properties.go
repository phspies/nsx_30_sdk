package nsxt

// Information about a management plane node this transport node is configured to communicate with
type BrokerProperties struct {
	// Indicates whether this broker is the master.
	BrokerIsMaster string `json:"BrokerIsMaster,omitempty"`
	// IP address or hostname of the message bus broker on the management plane node.
	BrokerIpAddress string `json:"BrokerIpAddress"`
	// Type of host running the broker.
	BrokerVirtualHost string `json:"BrokerVirtualHost,omitempty"`
	// Certificate thumbprint of the message bus broker on the management plane node.
	BrokerSslCertThumbprint string `json:"BrokerSslCertThumbprint"`
	// Port number of the message bus broker on the management plane node.
	BrokerPort string `json:"BrokerPort,omitempty"`
	// Fully qualified domain name of the message bus broker on the management plane node.
	BrokerFqdn string `json:"BrokerFqdn,omitempty"`
}
