# LbVirtualServer

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
**AccessListControl** | [***LbAccessListControl**](LbAccessListControl.md) |  | [optional] [default to null]
**IpProtocol** | **string** | Assigned Internet Protocol in IP header, TCP, UDP are supported.  | [optional] [default to IP_PROTOCOL.TCP]
**LogSignificantEventOnly** | **bool** | The property log_significant_event_only can take effect only when access_log_enabled is true. If log_significant_event_only is true, significant events are logged in access log. For L4 virtual server, significant event means unsuccessful(error or dropped) TCP/UDP connections. For L7 virtual server, significant event means unsuccessful connections or HTTP/HTTPS requests which have error response code(e.g. 4xx, 5xx).  | [optional] [default to false]
**ServerTcpProfileId** | **string** | Only L7 virtual server could be configured with customized server side TCP profile.  | [optional] [default to null]
**DefaultPoolMemberPorts** | **[]string** | If default_pool_member_ports are configured, both default_pool_member_port and default_pool_member_ports in the response payload would include port settings, notice that the value of default_pool_member_port is the first element of default_pool_member_ports.  | [optional] [default to null]
**PersistenceProfileId** | **string** | Persistence profile is used to allow related client connections to be sent to the same backend server.  | [optional] [default to null]
**ServerSslProfileBinding** | [***ServerSslProfileBinding**](ServerSslProfileBinding.md) |  | [optional] [default to null]
**PoolId** | **string** | The server pool(LbPool) contains backend servers. Server pool consists of one or more servers, also referred to as pool members, that are similarly configured and are running the same application.  | [optional] [default to null]
**ClientTcpProfileId** | **string** | Only L7 virtual server could be configured with customized client side TCP profile.  | [optional] [default to null]
**DefaultPoolMemberPort** | **string** | This is a deprecated property, please use &#x27;default_pool_member_ports&#x27; instead. If default_pool_member_port is configured and default_pool_member_ports are not specified, both default_pool_member_port and default_pool_member_ports in response payload would return the same port value. If both are specified, default_pool_member_ports setting would take effect with higher priority.  | [optional] [default to null]
**AccessLogEnabled** | **bool** | Whether access log is enabled | [optional] [default to false]
**MaxConcurrentConnections** | **int64** | To ensure one virtual server does not over consume resources, affecting other applications hosted on the same LBS, connections to a virtual server can be capped. If it is not specified, it means that connections are unlimited.  | [optional] [default to null]
**RuleIds** | **[]string** | Load balancer rules allow customization of load balancing behavior using match/action rules. Currently, load balancer rules are supported for only layer 7 virtual servers with LbHttpProfile.  | [optional] [default to null]
**MaxNewConnectionRate** | **int64** | To ensure one virtual server does not over consume resources, connections to a member can be rate limited. If it is not specified, it means that connection rate is unlimited.  | [optional] [default to null]
**SorryPoolId** | **string** | When load balancer can not select a backend server to serve the request in default pool or pool in rules, the request would be served by sorry server pool.  | [optional] [default to null]
**ClientSslProfileBinding** | [***ClientSslProfileBinding**](ClientSslProfileBinding.md) |  | [optional] [default to null]
**ApplicationProfileId** | **string** | The application profile defines the application protocol characteristics. It is used to influence how load balancing is performed. Currently, LbFastTCPProfile, LbFastUDPProfile and LbHttpProfile, etc are supported.  | [default to null]
**IpAddress** | **string** | virtual server IP address | [default to null]
**Port** | **string** | This is a deprecated property, please use &#x27;ports&#x27; instead. Port setting could be single port for both L7 mode and L4 mode. For L4 mode, a single port range is also supported. The port setting could be a single port or port range such as \&quot;80\&quot;, \&quot;1234-1236\&quot;. If port is configured and ports are not specified, both port and ports in response payload would return the same port value. If both port and ports are configured, ports setting would take effect with higher priority.  | [optional] [default to null]
**Enabled** | **bool** | whether the virtual server is enabled | [optional] [default to true]
**Ports** | **[]string** | Port setting could be a single port for both L7 mode and L4 mode. For L4 mode, multiple ports or port ranges are also supported such as \&quot;80\&quot;, \&quot;443\&quot;, \&quot;1234-1236\&quot;. If ports is configured, both port and ports in the response payload would include port settings, notice that the port field value is the first element of ports.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

