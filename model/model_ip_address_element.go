package nsxt

// IP Address
type IpAddressElement struct {
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// IPElement can be a single IP address, IP address range or a Subnet. Its type can be of IPv4 or IPv6. Supported list of formats are \"192.168.1.1\", \"192.168.1.1-192.168.1.100\", \"192.168.0.0/24\", \"fe80::250:56ff:fe83:318c\", \"fe80::250:56ff:fe83:3181-fe80::250:56ff:fe83:318c\", \"fe80::250:56ff:fe83:318c/64\" 
	IpAddress string `json:"ip_address"`
}
