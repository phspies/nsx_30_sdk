package nsxt

// An address binding entry is a combination of the IP-MAC-VLAN binding for a logical port. The address bindings can be obtained via various methods like ARP snooping, DHCP snooping etc. or by user configuration. 
type AddressBindingEntry struct {
	// Source from which the address binding entry was obtained
	Source string `json:"source,omitempty"`
	Binding *PacketAddressClassifier `json:"binding,omitempty"`
	// Timestamp at which the binding was discovered via snooping or manually specified by the user 
	BindingTimestamp int64 `json:"binding_timestamp,omitempty"`
}
