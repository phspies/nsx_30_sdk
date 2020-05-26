package nsxt

// Metadata fetched from an external system like Syslog or LogInsight.
type ErrorResolverSystemMetadata struct {
	// The value fetched from another system
	Value string `json:"value,omitempty"`
}
