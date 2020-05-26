package nsxt

type BgpNeighborAddressFamily struct {
	// Id of the IPPrefix List to be used for IN direction filter
	InFilterIpprefixlistId string `json:"in_filter_ipprefixlist_id,omitempty"`
	// Id of the RouteMap to be used for OUT direction filter
	OutFilterRoutemapId string `json:"out_filter_routemap_id,omitempty"`
	// Address family type
	Type_ string `json:"type"`
	// Id of the RouteMap to be used for IN direction filter
	InFilterRoutemapId string `json:"in_filter_routemap_id,omitempty"`
	// Enable this address family
	Enabled bool `json:"enabled,omitempty"`
	// Id of the IPPrefixList to be used for OUT direction filter
	OutFilterIpprefixlistId string `json:"out_filter_ipprefixlist_id,omitempty"`
	// Maximum number of routes supported on the address family
	MaximumRoutes int64 `json:"maximum_routes,omitempty"`
}
