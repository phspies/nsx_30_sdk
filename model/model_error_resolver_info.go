package nsxt

// Metadata related to a given error_id
type ErrorResolverInfo struct {
	// The error id for which metadata information is needed
	ErrorId int64 `json:"error_id"`
	// Indicates whether there is a resolver associated with the error or not
	ResolverPresent bool `json:"resolver_present"`
	UserMetadata *ErrorResolverUserMetadata `json:"user_metadata,omitempty"`
}
