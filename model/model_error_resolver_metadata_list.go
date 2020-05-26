package nsxt

// List of errors with their metadata
type ErrorResolverMetadataList struct {
	// List of errors with their corresponding metadata.
	Errors []ErrorResolverMetadata `json:"errors"`
}
