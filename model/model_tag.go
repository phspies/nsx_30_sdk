package nsxt

// Arbitrary key-value pairs that may be attached to an entity
type Tag struct {
	// Tag searches may optionally be restricted by scope
	Scope string `json:"scope,omitempty"`
	// Identifier meaningful to user with maximum length of 256 characters
	Tag string `json:"tag,omitempty"`
}
