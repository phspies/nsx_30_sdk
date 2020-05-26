package nsxt

type ManagementConfig struct {
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// True if Management nodes publish their fqdns(instead of default IP addresses) across NSX for its reachability.
	PublishFqdns bool `json:"publish_fqdns"`
}
