package nsxt

// Supported attributes and sub-attributes for NSProfile
type NsSupportedAttributes struct {
	// The type represent pre-defined or user defined list of supported attributes and sub-attributes that can be used while creating NSProfile 
	NsAttributes []NsAttributes `json:"ns_attributes"`
}
