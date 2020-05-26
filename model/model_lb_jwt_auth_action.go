package nsxt

// This action is used to control access to backend server resources using JSON Web Token(JWT) authentication. The JWT authentication is done before any HTTP manipulation if the HTTP request matches the given condition in LbRule. Any verification failed, the HTTP process will be terminated, and HTTP response with 401 status code and WWW-Authentication header will be returned to client. 
type LbJwtAuthAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// JWT is an open standard that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. Load balancer will search for every specified tokens one by one for the jwt message until found. This parameter is optional. In case not found or this field is not configured, load balancer searches the Bearer header by default in the http request \"Authorization: Bearer &lt;token&gt;\". 
	Tokens []string `json:"tokens,omitempty"`
	// Specify whether to pass the JWT to backend server or remove it. By default, it is false which means will not pass the JWT to backend servers. 
	PassJwtToPool bool `json:"pass_jwt_to_pool,omitempty"`
	// A description of the protected area. If no realm is specified, clients often display a formatted hostname instead. The configured realm is returned when client request is rejected with 401 http status. In the response, it will be \"WWW-Authentication: Bearer realm=&lt;realm&gt;\". 
	Realm string `json:"realm,omitempty"`
	Key *LbJwtKey `json:"key,omitempty"`
}
