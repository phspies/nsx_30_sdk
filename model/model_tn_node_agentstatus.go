package nsxt

type TnNodeAgentstatus struct {
	// Show the Node Agent connected VM vif status.
	HyperbusConnectionStatus string `json:"hyperbus_connection_status"`
	// Connected VM vif id.
	VifId string `json:"vif_id"`
}
