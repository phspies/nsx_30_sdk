package nsxt

// A set of IPv4 or IPv6 addresses defined by a network CIDR.
type IpPoolSubnet struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// The collection of upto 3 DNS servers for the subnet.
	DnsNameservers []string `json:"dns_nameservers,omitempty"`
	// Represents network address and the prefix length which will be associated with a layer-2 broadcast domain
	Cidr string `json:"cidr"`
	// The default gateway address on a layer-3 router.
	GatewayIp string `json:"gateway_ip,omitempty"`
	// A collection of IPv4 or IPv6 IP Pool Ranges.
	AllocationRanges []IpPoolRange `json:"allocation_ranges"`
	// The DNS suffix for the DNS server.
	DnsSuffix string `json:"dns_suffix,omitempty"`
}
