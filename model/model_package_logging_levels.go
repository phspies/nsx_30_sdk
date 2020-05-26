package nsxt

type PackageLoggingLevels struct {
	// Logging levels per package
	LoggingLevel string `json:"logging_level,omitempty"`
	// Package name
	PackageName string `json:"package_name,omitempty"`
}
