package nsxt

type HostedEntityInfo struct {
	// Unique identifier of entity
	EntityUuid string `json:"entity_uuid,omitempty"`
	// The type of entity hosted could be MP, CCP, VMC App etc.
	EntityType string `json:"entity_type,omitempty"`
}
