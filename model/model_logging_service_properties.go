package nsxt

// Service properties
type LoggingServiceProperties struct {
	// Package logging levels
	PackageLoggingLevel []PackageLoggingLevels `json:"package_logging_level,omitempty"`
	// Service logging level
	LoggingLevel string `json:"logging_level"`
	// Modified package logging levels
	ModifiedPackageLoggingLevels string `json:"modified_package_logging_levels,omitempty"`
}
