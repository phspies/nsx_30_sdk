package nsxt

// Data type collection configuration
type DataTypeCollectionConfiguration struct {
	// The frequency in seconds at which data is collected
	CollectionFrequency int64 `json:"collection_frequency"`
	// Defines the type of data being collected
	DataType string `json:"data_type"`
}
