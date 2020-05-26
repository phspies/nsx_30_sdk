package nsxt

// Information about the management plane this controller is communciating with
type ManagementPlaneProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// The account name to use when authenticating to the management plane's message bus.
	Account *interface{} `json:"account,omitempty"`
	// The shared secret to use when autnenticating to the management plane's message bus. Not returned in REST responses.
	Secret string `json:"secret,omitempty"`
	// The list of messaging brokers this controller is configured with.
	Brokers []ManagementPlaneBrokerProperties `json:"brokers,omitempty"`
}
