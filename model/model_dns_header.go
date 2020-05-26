package nsxt

type DnsHeader struct {
	// This is used to specify the type of the address. V4 - The address provided is an IPv4 domain name/IP address, the Type in query or response will be A V6 - The address provided is an IPv6 domain name/IP address, the Type in query or response will be AAAA
	AddressType string `json:"address_type,omitempty"`
	// Specifies the message type whether it is a query or a response.
	MessageType string `json:"message_type,omitempty"`
	// This is used to define what is being asked or responded.
	Address string `json:"address,omitempty"`
}