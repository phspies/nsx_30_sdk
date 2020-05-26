package nsxt

// Information about the management plane this node is communciating with
type MpaConfigProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// The nodes client type.
	RmqClientType *interface{} `json:"RmqClientType,omitempty"`
	// The list of messaging brokers this controller is configured with.
	RmqBrokerCluster []BrokerProperties `json:"RmqBrokerCluster,omitempty"`
	// The shared secret to use when autnenticating to the management plane's message bus. Not returned in REST responses.
	SharedSecret string `json:"SharedSecret,omitempty"`
	// The account name to use when authenticating to the management plane's message bus.
	AccountName *interface{} `json:"AccountName,omitempty"`
}
