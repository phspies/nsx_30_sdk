package nsxt

type RouteMapSequenceSet struct {
	// For incoming and import route_maps on receiving both v6 global and v6 link-local address for the route, prefer to use the global address as the next hop. By default, it prefers the link-local next hop. 
	PreferGlobalV6NextHop bool `json:"prefer_global_v6_next_hop,omitempty"`
	// Local preference indicates the degree of preference for one BGP route over other BGP routes. The path/route with highest local preference value is preferred/selected. If local preference value is not specified then it will be considered as 100 by default. 
	LocalPreference int64 `json:"local_preference,omitempty"`
	// Weight used to select certain path
	Weight int32 `json:"weight,omitempty"`
	// Set large BGP community, community value shoud be in aa:bb:nn format where aa, bb, nn are unsigned integers with range [1-4294967295].
	LargeCommunity string `json:"large_community,omitempty"`
	// As Path Prepending to influence path selection
	AsPathPrepend string `json:"as_path_prepend,omitempty"`
	// Set normal BGP community either well-known community name or community value in aa:nn(2byte:2byte) format. 
	Community string `json:"community,omitempty"`
	// Multi Exit Discriminator (MED)
	MultiExitDiscriminator int64 `json:"multi_exit_discriminator,omitempty"`
}
