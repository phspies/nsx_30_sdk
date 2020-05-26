package nsxt

type CryptoAlgorithm struct {
	// Supported key sizes for the algorithm.
	KeySize []int64 `json:"key_size,omitempty"`
	// Crypto algorithm name.
	Name string `json:"name,omitempty"`
}
