package nsxt

// Next hop to prefix lists mapping.
type NextHopPrefixListsMapping struct {
	// Array of Prefix list UUIDs.
	PrefixLists []string `json:"prefix_lists"`
	// Next hop address.
	NextHop string `json:"next_hop"`
}
