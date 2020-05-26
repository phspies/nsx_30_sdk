package nsxt

// Error along with its metadata
type ErrorResolverMetadata struct {
	// The error id as reported by the entity where the error occurred.
	ErrorId int64 `json:"error_id"`
	SystemMetadata *ErrorResolverSystemMetadata `json:"system_metadata,omitempty"`
	// The entity/node UUID where the error has occurred.
	EntityId string `json:"entity_id"`
	UserMetadata *ErrorResolverUserMetadata `json:"user_metadata,omitempty"`
}
