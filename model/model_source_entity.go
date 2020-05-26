package nsxt

// Service insertion data path inserts unique 'source node id' value into each packet before it received by Service VM. This value can be resolved to multiple Source Entities. It represents source of the packets.
type SourceEntity struct {
	// Type of source entity. Currently source value can be resolved to VIF and Virtual Machine.
	SourceEntityType string `json:"source_entity_type,omitempty"`
	// UUID of Source entity
	SourceEntityId string `json:"source_entity_id,omitempty"`
}
