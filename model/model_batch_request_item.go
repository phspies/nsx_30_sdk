package nsxt

// A single request within a batch of operations
type BatchRequestItem struct {
	Body *interface{} `json:"body,omitempty"`
	// relative uri (path and args), of the call including resource id (if this is a POST/DELETE), exclude hostname and port and prefix, exploded form of parameters
	Uri string `json:"uri"`
	// http method type
	Method string `json:"method"`
}
