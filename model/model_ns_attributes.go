package nsxt

// NS Attributes data holder structure
type NsAttributes struct {
	// Reference to sub attributes for the attribute
	SubAttributes []NsAttributesData `json:"sub_attributes,omitempty"`
	AttributesData *NsAttributesData `json:"attributes_data"`
}
