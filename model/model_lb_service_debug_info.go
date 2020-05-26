package nsxt

// The information for a given load balancer service could be used for debugging and troubleshooting. It includes load balancer service, associated virtual servers, associated pools, associated profiles such as persistence, SSL, application, associated monitors and associated rules. 
type LbServiceDebugInfo struct {
	// The pools which are associated to the given load balancer service would be included. The pools could be defined in virtual server default pool, sorry pool or load balancer rule action. 
	Pools []LbPool `json:"pools,omitempty"`
	// The TCP profiles are associated to virtual servers 
	TcpProfiles []LbTcpProfile `json:"tcp_profiles,omitempty"`
	// The virtual servers which are associated to the given load balancer service would be included. 
	VirtualServers []LbVirtualServer `json:"virtual_servers,omitempty"`
	// The client SSL profiles are associated to virtual servers 
	ClientSslProfiles []LbClientSslProfile `json:"client_ssl_profiles,omitempty"`
	// The server SSL profiles are associated to virtual servers 
	ServerSslProfiles []LbServerSslProfile `json:"server_ssl_profiles,omitempty"`
	Service *LbService `json:"service,omitempty"`
	// The load balancer rules are associated to virtual servers 
	Rules []LbRule `json:"rules,omitempty"`
	// The application profiles are associated to virtual servers 
	ApplicationProfiles []LbAppProfile `json:"application_profiles,omitempty"`
	// The persistence profiles are associated to virtual servers 
	PersistenceProfiles []LbPersistenceProfile `json:"persistence_profiles,omitempty"`
	// The load balancer monitors are associated to pools. 
	Monitors []LbMonitor `json:"monitors,omitempty"`
}
