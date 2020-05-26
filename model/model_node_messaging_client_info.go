package nsxt

type NodeMessagingClientInfo struct {
	// A list of messaging clients owned by this entity
	Clients []MessagingClientInfo `json:"clients,omitempty"`
	// Entity type of this messaging client
	EntityType string `json:"entity_type,omitempty"`
}
