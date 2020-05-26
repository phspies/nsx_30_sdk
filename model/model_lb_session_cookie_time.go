package nsxt

type LbSessionCookieTime struct {
	// Both session cookie and persistence cookie are supported, Use LbSessionCookieTime for session cookie time setting, Use LbPersistenceCookieTime for persistence cookie time setting 
	Type_ string `json:"type"`
	// Instead of using HTTP Cookie max-age and relying on client to expire the cookie, max idle time and/or max lifetime of the cookie can be used. Max idle time, if configured, specifies the maximum interval the cookie is valid for from the last time it was seen in a request. It is available for insert mode. 
	CookieMaxIdle int64 `json:"cookie_max_idle,omitempty"`
	// Max life time, if configured, specifies the maximum interval the cookie is valid for from the first time the cookie was seen in a request. It is available for insert mode. 
	CookieMaxLife int64 `json:"cookie_max_life,omitempty"`
}
