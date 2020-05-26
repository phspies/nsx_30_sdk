package nsxt

type Principal struct {
	// Certificate list.
	Attributes []KeyValue `json:"attributes"`
}
