package nsxt

type TrustManagementData struct {
	// List of supported algorithms.
	SupportedAlgorithms []CryptoAlgorithm `json:"supported_algorithms,omitempty"`
}
