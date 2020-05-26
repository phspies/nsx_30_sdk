package nsxt

// To filter the routes advertised by the TIER1 LR to TIER0 LR. Filtering will be based on the type of route and the prefix operator configured.
type AdvertisementRuleFilter struct {
	// GE prefix operator filters all the routes having network subset of any of the networks configured in Advertise rule. EQ prefix operator filter all the routes having network equal to any of the network configured in Advertise rule.
	PrefixOperator string `json:"prefix_operator"`
	// Array of route types to filter routes
	MatchRouteTypes []string `json:"match_route_types"`
}
