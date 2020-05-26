package nsxt

// Define DHCP options other than option 121.
type GenericDhcpOption struct {
	// Code of the dhcp option.
	Code int64 `json:"code"`
	// Value of the option.
	Values []string `json:"values"`
}
