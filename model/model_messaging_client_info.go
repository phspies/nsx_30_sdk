package nsxt

type MessagingClientInfo struct {
	// Type of messaging client
	ClientType string `json:"client_type,omitempty"`
	// Account name in messaging client
	AccountName string `json:"account_name,omitempty"`
}
