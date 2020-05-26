package nsxt

type Filter struct {
	// The name of the filter.
	Name string `json:"name"`
	// The value of the filter.
	Value string `json:"value"`
}
