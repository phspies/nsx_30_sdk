package nsxt

// A profile holding DNS configuration which allows firewall to use DNS response packets and manage firewall cache. e.g. TTL used to expire snooped entries from cache.
type FirewallDnsProfile struct {
	// Resource type to use as profile type
	ResourceType string `json:"resource_type"`
	DnsTtlConfig *DnsTtlConfig `json:"dns_ttl_config,omitempty"`
}
