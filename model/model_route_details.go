package nsxt

// BGP route details.
type RouteDetails struct {
	// BGP Multi Exit Discriminator attribute.
	Med int64 `json:"med,omitempty"`
	// Next hop IP address.
	NextHop string `json:"next_hop,omitempty"`
	// CIDR network address.
	Network string `json:"network,omitempty"`
	// BGP Weight attribute.
	Weight int64 `json:"weight,omitempty"`
	// BGP Local Preference attribute.
	LocalPref int64 `json:"local_pref,omitempty"`
	// BGP AS path attribute.
	AsPath string `json:"as_path,omitempty"`
}
