package nsxt

// This action is used to select SSL mode. Three types of SSL mode actions can be specified in Transport phase, ssl passthrough, ssl offloading and ssl end-to-end. 
type LbSslModeSelectionAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// SSL Passthrough: LB establishes a TCP connection with client and another connection with selected backend server. LB won't inspect the stream data between client and backend server, but just pass it through. Backend server exchanges SSL connection with client. SSL Offloading: LB terminiates the connections from client, and establishes SSL connection with it. After receiving the HTTP request, LB connects the selected backend server and talk with it via HTTP without SSL. LB estalishes new connection to selected backend server for each HTTP request, in case ntlm or multiplexing are NOT configured. SSL End-to-End: LB terminiates the connections from client, and establishes SSL connection with it. After receiving the HTTP request, LB connects the selected backend server and talk with it via HTTPS. LB estalishes new SSL connection to selected backend server for each HTTP request, in case ntlm or multiplexing are NOT configured. 
	SslMode string `json:"ssl_mode"`
}
