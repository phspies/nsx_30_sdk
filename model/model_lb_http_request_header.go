package nsxt

type LbHttpRequestHeader struct {
	// Value of HTTP request header
	HeaderValue string `json:"header_value"`
	// Name of HTTP request header
	HeaderName string `json:"header_name"`
}
