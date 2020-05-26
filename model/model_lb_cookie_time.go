package nsxt

type LbCookieTime struct {
	// Both session cookie and persistence cookie are supported, Use LbSessionCookieTime for session cookie time setting, Use LbPersistenceCookieTime for persistence cookie time setting 
	Type_ string `json:"type"`
}
