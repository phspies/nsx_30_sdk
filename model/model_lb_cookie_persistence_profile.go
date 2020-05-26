package nsxt

type LbCookiePersistenceProfile struct {
	// The persistence shared flag identifies whether the persistence table is shared among virtual-servers referring this profile. If persistence shared flag is not set in the cookie persistence profile bound to a virtual server, it defaults to cookie persistence that is private to each virtual server and is qualified by the pool. This is accomplished by load balancer inserting a cookie with name in the format &lt;name&gt;.&lt;virtual_server_id&gt;.&lt;pool_id&gt;. If persistence shared flag is set in the cookie persistence profile, in cookie insert mode, cookie persistence could be shared across multiple virtual servers that are bound to the same pools. The cookie name would be changed to &lt;name&gt;.&lt;profile-id&gt;.&lt;pool-id&gt;. If persistence shared flag is not set in the sourceIp persistence profile bound to a virtual server, each virtual server that the profile is bound to maintains its own private persistence table. If persistence shared flag is set in the sourceIp persistence profile, all virtual servers the profile is bound to share the same persistence table. If persistence shared flag is not set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using both virtual server ID and profile ID. If persistence shared flag is set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using profile ID. It means that virtual servers which consume the same profile in the LbRule with this flag enabled are sharing the same persistence table. 
	PersistenceShared bool `json:"persistence_shared,omitempty"`
	// The resource_type property identifies persistence profile type. 
	ResourceType string `json:"resource_type"`
	// If garble is set to true, cookie value (server IP and port) would be encrypted. If garble is set to false, cookie value would be plain text. 
	CookieGarble bool `json:"cookie_garble,omitempty"`
	// If fallback is true, once the cookie points to a server that is down (i.e. admin state DISABLED or healthcheck state is DOWN), then a new server is selected by default to handle that request. If fallback is false, it will cause the request to be rejected if cookie points to a server 
	CookieFallback bool `json:"cookie_fallback,omitempty"`
	// cookie persistence mode
	CookieMode string `json:"cookie_mode,omitempty"`
	// HTTP cookie domain could be configured, only available for insert mode. 
	CookieDomain string `json:"cookie_domain,omitempty"`
	// cookie name
	CookieName string `json:"cookie_name"`
	CookieTime *LbCookieTime `json:"cookie_time,omitempty"`
	// HTTP cookie path could be set, only available for insert mode. 
	CookiePath string `json:"cookie_path,omitempty"`
}
