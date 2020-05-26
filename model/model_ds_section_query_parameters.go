package nsxt

// Section query parameters
type DsSectionQueryParameters struct {
	// Flag to cascade delete of this object to all it's child objects.
	Cascade bool `json:"cascade,omitempty"`
}
