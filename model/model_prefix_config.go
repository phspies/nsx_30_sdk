package nsxt

type PrefixConfig struct {
	// Action for the IPPrefix
	Action string `json:"action"`
	// Greater than or equal to
	Ge int64 `json:"ge,omitempty"`
	// Less than or equal to
	Le int64 `json:"le,omitempty"`
	// If absent, the action applies to all addresses.
	Network string `json:"network,omitempty"`
}
