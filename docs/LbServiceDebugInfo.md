# LbServiceDebugInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | [**[]LbPool**](LbPool.md) | The pools which are associated to the given load balancer service would be included. The pools could be defined in virtual server default pool, sorry pool or load balancer rule action.  | [optional] [default to null]
**TcpProfiles** | [**[]LbTcpProfile**](LbTcpProfile.md) | The TCP profiles are associated to virtual servers  | [optional] [default to null]
**VirtualServers** | [**[]LbVirtualServer**](LbVirtualServer.md) | The virtual servers which are associated to the given load balancer service would be included.  | [optional] [default to null]
**ClientSslProfiles** | [**[]LbClientSslProfile**](LbClientSslProfile.md) | The client SSL profiles are associated to virtual servers  | [optional] [default to null]
**ServerSslProfiles** | [**[]LbServerSslProfile**](LbServerSslProfile.md) | The server SSL profiles are associated to virtual servers  | [optional] [default to null]
**Service** | [***LbService**](LbService.md) |  | [optional] [default to null]
**Rules** | [**[]LbRule**](LbRule.md) | The load balancer rules are associated to virtual servers  | [optional] [default to null]
**ApplicationProfiles** | [**[]LbAppProfile**](LbAppProfile.md) | The application profiles are associated to virtual servers  | [optional] [default to null]
**PersistenceProfiles** | [**[]LbPersistenceProfile**](LbPersistenceProfile.md) | The persistence profiles are associated to virtual servers  | [optional] [default to null]
**Monitors** | [**[]LbMonitor**](LbMonitor.md) | The load balancer monitors are associated to pools.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

