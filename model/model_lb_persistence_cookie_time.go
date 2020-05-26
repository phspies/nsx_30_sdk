package nsxt

type LbPersistenceCookieTime struct {
	// Both session cookie and persistence cookie are supported, Use LbSessionCookieTime for session cookie time setting, Use LbPersistenceCookieTime for persistence cookie time setting 
	Type_ string `json:"type"`
	// HTTP cookie max-age to expire cookie, only available for insert mode. 
	CookieMaxIdle int64 `json:"cookie_max_idle"`
}
