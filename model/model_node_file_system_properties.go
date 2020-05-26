package nsxt

// File system properties
type NodeFileSystemProperties struct {
	// File system mount
	Mount string `json:"mount,omitempty"`
	// File system size in kilobytes
	Total int64 `json:"total,omitempty"`
	// File system type
	Type_ string `json:"type,omitempty"`
	// File system id
	FileSystem string `json:"file_system,omitempty"`
	// Amount of file system used in kilobytes
	Used int64 `json:"used,omitempty"`
}
