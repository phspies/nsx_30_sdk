package nsxt

// Specify which HostSwitch from this TransportNode is used handle traffic for given TransportZone
type TransportZoneEndPoint struct {
	// Unique ID identifying the transport zone for this endpoint
	TransportZoneId string `json:"transport_zone_id"`
	// Identifiers of the transport zone profiles associated with this transport zone endpoint on this transport node.
	TransportZoneProfileIds []TransportZoneProfileTypeIdEntry `json:"transport_zone_profile_ids,omitempty"`
}
