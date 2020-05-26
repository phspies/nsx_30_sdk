# LbSourceIpPersistenceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersistenceShared** | **bool** | The persistence shared flag identifies whether the persistence table is shared among virtual-servers referring this profile. If persistence shared flag is not set in the cookie persistence profile bound to a virtual server, it defaults to cookie persistence that is private to each virtual server and is qualified by the pool. This is accomplished by load balancer inserting a cookie with name in the format &amp;lt;name&amp;gt;.&amp;lt;virtual_server_id&amp;gt;.&amp;lt;pool_id&amp;gt;. If persistence shared flag is set in the cookie persistence profile, in cookie insert mode, cookie persistence could be shared across multiple virtual servers that are bound to the same pools. The cookie name would be changed to &amp;lt;name&amp;gt;.&amp;lt;profile-id&amp;gt;.&amp;lt;pool-id&amp;gt;. If persistence shared flag is not set in the sourceIp persistence profile bound to a virtual server, each virtual server that the profile is bound to maintains its own private persistence table. If persistence shared flag is set in the sourceIp persistence profile, all virtual servers the profile is bound to share the same persistence table. If persistence shared flag is not set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using both virtual server ID and profile ID. If persistence shared flag is set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using profile ID. It means that virtual servers which consume the same profile in the LbRule with this flag enabled are sharing the same persistence table.  | [optional] [default to false]
**ResourceType** | **string** | The resource_type property identifies persistence profile type.  | [default to null]
**Purge** | **string** | persistence purge setting | [optional] [default to PURGE.FULL]
**HaPersistenceMirroringEnabled** | **bool** | Persistence entries are not synchronized to the HA peer by default.  | [optional] [default to false]
**Timeout** | **int64** | When all connections complete (reference count reaches 0), persistence entry timer is started with the expiration time.  | [optional] [default to 300]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

