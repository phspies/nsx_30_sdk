package nsxt

// Contains migration related information about logical constructs
type LogicalConstructMigrationStats struct {
	// Type of the vSphere object (e.g. dvportgroup).
	SourceType string `json:"source_type,omitempty"`
	// Number of objects of source type.
	SourceCount string `json:"source_count,omitempty"`
	// Type of the Policy object corresponding to the source type (e.g. Segment).
	TargetType string `json:"target_type,omitempty"`
	// Functional area that this vSphere object falls into
	Vertical string `json:"vertical,omitempty"`
	// Number of objects of target type.
	TargetCount string `json:"target_count,omitempty"`
}