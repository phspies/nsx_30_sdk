package nsxt

// Answer of nslookup
type DnsQueryAnswer struct {
	// Unparsed answer string from raw_answer. 
	RawString string `json:"raw_string,omitempty"`
	// Matched name of the given address. 
	Name string `json:"name,omitempty"`
	// Can be resolved ip address. 
	Address string `json:"address,omitempty"`
}
