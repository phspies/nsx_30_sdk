# LbCookiePersistenceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersistenceShared** | **bool** | The persistence shared flag identifies whether the persistence table is shared among virtual-servers referring this profile. If persistence shared flag is not set in the cookie persistence profile bound to a virtual server, it defaults to cookie persistence that is private to each virtual server and is qualified by the pool. This is accomplished by load balancer inserting a cookie with name in the format &amp;lt;name&amp;gt;.&amp;lt;virtual_server_id&amp;gt;.&amp;lt;pool_id&amp;gt;. If persistence shared flag is set in the cookie persistence profile, in cookie insert mode, cookie persistence could be shared across multiple virtual servers that are bound to the same pools. The cookie name would be changed to &amp;lt;name&amp;gt;.&amp;lt;profile-id&amp;gt;.&amp;lt;pool-id&amp;gt;. If persistence shared flag is not set in the sourceIp persistence profile bound to a virtual server, each virtual server that the profile is bound to maintains its own private persistence table. If persistence shared flag is set in the sourceIp persistence profile, all virtual servers the profile is bound to share the same persistence table. If persistence shared flag is not set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using both virtual server ID and profile ID. If persistence shared flag is set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using profile ID. It means that virtual servers which consume the same profile in the LbRule with this flag enabled are sharing the same persistence table.  | [optional] [default to false]
**ResourceType** | **string** | The resource_type property identifies persistence profile type.  | [default to null]
**CookieGarble** | **bool** | If garble is set to true, cookie value (server IP and port) would be encrypted. If garble is set to false, cookie value would be plain text.  | [optional] [default to true]
**CookieFallback** | **bool** | If fallback is true, once the cookie points to a server that is down (i.e. admin state DISABLED or healthcheck state is DOWN), then a new server is selected by default to handle that request. If fallback is false, it will cause the request to be rejected if cookie points to a server  | [optional] [default to true]
**CookieMode** | **string** | cookie persistence mode | [optional] [default to COOKIE_MODE.INSERT]
**CookieDomain** | **string** | HTTP cookie domain could be configured, only available for insert mode.  | [optional] [default to null]
**CookieName** | **string** | cookie name | [default to null]
**CookieTime** | [***LbCookieTime**](LbCookieTime.md) |  | [optional] [default to null]
**CookiePath** | **string** | HTTP cookie path could be set, only available for insert mode.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

