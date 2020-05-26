package nsxt

// The configuration data for setting the global realization state barrier 
type RealizationStateBarrierConfig struct {
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// The interval in milliseconds used for auto incrementing the barrier number 
	Interval int64 `json:"interval"`
}
