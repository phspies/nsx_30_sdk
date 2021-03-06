# LbPool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | Unique identifier of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**MemberGroup** | [***PoolMemberGroup**](PoolMemberGroup.md) |  | [optional] [default to null]
**SnatTranslation** | [***LbSnatTranslation**](LbSnatTranslation.md) |  | [optional] [default to null]
**Algorithm** | **string** | Load balancing algorithm, configurable per pool controls how the incoming connections are distributed among the members.  | [optional] [default to ALGORITHM.ROUND_ROBIN]
**Members** | [**[]PoolMember**](PoolMember.md) | Server pool consists of one or more pool members. Each pool member is identified, typically, by an IP address and a port.  | [optional] [default to null]
**PassiveMonitorId** | **string** | Passive healthchecks are disabled by default and can be enabled by attaching a passive health monitor to a server pool. Each time a client connection to a pool member fails, its failed count is incremented. For pools bound to L7 virtual servers, a connection is considered to be failed and failed count is incremented if any TCP connection errors (e.g. TCP RST or failure to send data) or SSL handshake failures occur. For pools bound to L4 virtual servers, if no response is received to a TCP SYN sent to the pool member or if a TCP RST is received in response to a TCP SYN, then the pool member is considered to have failed and the failed count is incremented.  | [optional] [default to null]
**TcpMultiplexingNumber** | **int64** | The maximum number of TCP connections per pool that are idly kept alive for sending future client requests.  | [optional] [default to 6]
**ActiveMonitorIds** | **[]string** | In case of active healthchecks, load balancer itself initiates new connections (or sends ICMP ping) to the servers periodically to check their health, completely independent of any data traffic. Active healthchecks are disabled by default and can be enabled for a server pool by binding a health monitor to the pool. If multiple active monitors are configured, the pool member status is UP only when the health check status for all the monitors are UP.  | [optional] [default to null]
**TcpMultiplexingEnabled** | **bool** | TCP multiplexing allows the same TCP connection between load balancer and the backend server to be used for sending multiple client requests from different client TCP connections.  | [optional] [default to false]
**MinActiveMembers** | **int64** | A pool is considered active if there are at least certain minimum number of members.  | [optional] [default to 1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

