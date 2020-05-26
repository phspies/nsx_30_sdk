package nsxt

type MetadataProxyStatisticsPerLogicalSwitch struct {
	// requests to nova server
	RequestsToNovaServer int64 `json:"requests_to_nova_server"`
	// responses to clients
	ResponsesToClients int64 `json:"responses_to_clients"`
	// succeeded responses from  nova server
	SucceededResponsesFromNovaServer int64 `json:"succeeded_responses_from_nova_server"`
	// uuid of attached logical switch
	LogicalSwitchId string `json:"logical_switch_id"`
	// requests from clients
	RequestsFromClients int64 `json:"requests_from_clients"`
	// error responses from  nova server
	ErrorResponsesFromNovaServer int64 `json:"error_responses_from_nova_server"`
}
