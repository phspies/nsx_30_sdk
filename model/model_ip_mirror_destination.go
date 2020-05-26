package nsxt

type IpMirrorDestination struct {
	// Resource types of mirror destination
	ResourceType string `json:"resource_type"`
	// The destination IPs of the mirror packet will be sent to.
	DestinationIps []string `json:"destination_ips"`
	// You can choose GRE, ERSPAN II or ERSPAN III.
	EncapsulationType string `json:"encapsulation_type"`
	// Used by physical switch for the mirror traffic forwarding. Must be provided and only effective when encapsulation type is ERSPAN type II or type III. 
	ErspanId int32 `json:"erspan_id,omitempty"`
	// User-configurable 32-bit key only for GRE
	GreKey int32 `json:"gre_key,omitempty"`
}
