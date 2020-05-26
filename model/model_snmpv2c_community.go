package nsxt

// This object contains SNMP v2c community identifier, shared secret and access properties.
type Snmpv2cCommunity struct {
	// Access permissions for polling NSX nodes over SNMP v2c.
	Access string `json:"access,omitempty"`
	// Unique, non-sensitive community name to identify community.
	CommunityName string `json:"community_name"`
	// Community string. This is considered a shared secret and therefore sensitive information. This field is required when adding a community. When updating a community, do not include this field in the request. If this field is present in an update request, it will be considered as a new value for community string.
	CommunityString string `json:"community_string,omitempty"`
}
