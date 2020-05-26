package nsxt

type LbSourceIpPersistenceProfile struct {
	// The persistence shared flag identifies whether the persistence table is shared among virtual-servers referring this profile. If persistence shared flag is not set in the cookie persistence profile bound to a virtual server, it defaults to cookie persistence that is private to each virtual server and is qualified by the pool. This is accomplished by load balancer inserting a cookie with name in the format &lt;name&gt;.&lt;virtual_server_id&gt;.&lt;pool_id&gt;. If persistence shared flag is set in the cookie persistence profile, in cookie insert mode, cookie persistence could be shared across multiple virtual servers that are bound to the same pools. The cookie name would be changed to &lt;name&gt;.&lt;profile-id&gt;.&lt;pool-id&gt;. If persistence shared flag is not set in the sourceIp persistence profile bound to a virtual server, each virtual server that the profile is bound to maintains its own private persistence table. If persistence shared flag is set in the sourceIp persistence profile, all virtual servers the profile is bound to share the same persistence table. If persistence shared flag is not set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using both virtual server ID and profile ID. If persistence shared flag is set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using profile ID. It means that virtual servers which consume the same profile in the LbRule with this flag enabled are sharing the same persistence table. 
	PersistenceShared bool `json:"persistence_shared,omitempty"`
	// The resource_type property identifies persistence profile type. 
	ResourceType string `json:"resource_type"`
	// persistence purge setting
	Purge string `json:"purge,omitempty"`
	// Persistence entries are not synchronized to the HA peer by default. 
	HaPersistenceMirroringEnabled bool `json:"ha_persistence_mirroring_enabled,omitempty"`
	// When all connections complete (reference count reaches 0), persistence entry timer is started with the expiration time. 
	Timeout int64 `json:"timeout,omitempty"`
}
