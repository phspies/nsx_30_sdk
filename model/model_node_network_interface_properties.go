package nsxt

// Node network interface properties
type NodeNetworkInterfaceProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Interface MAC address
	PhysicalAddress string `json:"physical_address,omitempty"`
	// Interface broadcast address
	BroadcastAddress string `json:"broadcast_address,omitempty"`
	// Interface administration status
	LinkStatus string `json:"link_status,omitempty"`
	// Interface's default gateway
	DefaultGateway string `json:"default_gateway,omitempty"`
	// Bond's primary device name in active-backup bond mode
	BondPrimary string `json:"bond_primary,omitempty"`
	// Bond's slave devices
	BondSlaves []string `json:"bond_slaves,omitempty"`
	// Interface IP addresses
	IpAddresses []IPv4AddressProperties `json:"ip_addresses,omitempty"`
	// VLAN Id
	Vlan int64 `json:"vlan,omitempty"`
	// Bond mode
	BondMode string `json:"bond_mode,omitempty"`
	// Interface ID
	InterfaceId string `json:"interface_id,omitempty"`
	// Interface administration status
	AdminStatus string `json:"admin_status,omitempty"`
	// Interface plane
	Plane string `json:"plane,omitempty"`
	// Interface is a KNI
	IsKni bool `json:"is_kni,omitempty"`
	// Interface configuration
	IpConfiguration string `json:"ip_configuration"`
	// Interface MTU
	Mtu int64 `json:"mtu,omitempty"`
}