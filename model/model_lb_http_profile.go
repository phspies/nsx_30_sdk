package nsxt

type LbHttpProfile struct {
	// An application profile can be bound to a virtual server to specify the application protocol characteristics. It is used to influence how load balancing is performed. Currently, three types of application profiles are supported: LbFastTCPProfile, LbFastUDPProfile and LbHttpProfile. LbFastTCPProfile or LbFastUDPProfile is typically used when the application is using a custom protocol or a standard protocol not supported by the load balancer. It is also used in cases where the user only wants L4 load balancing mainly because L4 load balancing has much higher performance and scalability, and/or supports connection mirroring. LbHttpProfile is used for both HTTP and HTTPS applications. Though application rules, if bound to the virtual server, can be used to accomplish the same goal, LbHttpProfile is intended to simplify enabling certain common use cases. 
	ResourceType string `json:"resource_type"`
	// When buffering is disabled, the response is passed to a client synchronously, immediately as it is received. When buffering is enabled, LB receives a response from the backend server as soon as possible, saving it into the buffers. 
	ResponseBuffering bool `json:"response_buffering,omitempty"`
	// If server doesn't send any packet within this time, the connection is closed. 
	ResponseTimeout int64 `json:"response_timeout,omitempty"`
	// It is used to specify the HTTP application idle timeout, it means that how long the load balancer will keep the connection idle to wait for the client to send the next keep-alive request. It is not a TCP socket setting. 
	IdleTimeout int64 `json:"idle_timeout,omitempty"`
	// If it is not specified, it means that request body size is unlimited. 
	RequestBodySize int64 `json:"request_body_size,omitempty"`
	// A response with header larger than response_header_size will be dropped. 
	ResponseHeaderSize int64 `json:"response_header_size,omitempty"`
	// NTLM is an authentication protocol that can be used over HTTP. If the flag is set to true, LB will use NTLM challenge/response methodology. 
	Ntlm bool `json:"ntlm,omitempty"`
	// A request with header equal to or below this size is guaranteed to be processed. A request with header larger than request_header_size will be processed up to 32K bytes on best effort basis. 
	RequestHeaderSize int64 `json:"request_header_size,omitempty"`
	// If a website is temporarily down or has moved, incoming requests for that virtual server can be temporarily redirected to a URL 
	HttpRedirectTo string `json:"http_redirect_to,omitempty"`
	// When X-Forwareded-For is configured, X-Forwarded-Proto and X-Forwarded-Port information is added automatically. The two additional header information can be also modified or deleted in load balancer rules. 
	XForwardedFor string `json:"x_forwarded_for,omitempty"`
	// Certain secure applications may want to force communication over SSL, but instead of rejecting non-SSL connections, they may choose to redirect the client automatically to use SSL. 
	HttpRedirectToHttps bool `json:"http_redirect_to_https,omitempty"`
}
