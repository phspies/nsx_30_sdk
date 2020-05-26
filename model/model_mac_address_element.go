package nsxt

// MAC Address
type MacAddressElement struct {
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// A MAC address. Must be 6 pairs of hexadecimal digits, upper or lower case, separated by colons or dashes. Examples: 01:23:45:67:89:ab, 01-23-45-67-89-AB. 
	MacAddress string `json:"mac_address"`
}
