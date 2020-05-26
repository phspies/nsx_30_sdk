# BgpNeighbor

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
**GracefulRestartMode** | **string** | BGP Graceful Restart mode. If specified, then it will take precedence over global Graceful Restart mode configured in logical router BgpConfig otherwise BgpConfig level Graceful Restart mode will be applicable for peer.  | [optional] [default to null]
**RemoteAs** | **int64** | This is a deprecated property, Please use &#x27;remote_as_num&#x27; instead. | [optional] [default to null]
**FilterOutIpprefixlistId** | **string** | This is a deprecated property, Please use &#x27;address_family&#x27; instead. | [optional] [default to null]
**HoldDownTimer** | **int64** | Wait period (seconds) before declaring peer dead | [optional] [default to 180]
**SourceAddresses** | **[]string** | BGP neighborship will be formed from all these source addresses to this neighbour. | [optional] [default to null]
**MaximumHopLimit** | **int32** | This value is set on TTL(time to live) of BGP header. When router receives the BGP packet, it decrements the TTL. The default value of TTL is one when BPG request is initiated.So in the case of a BGP peer multiple hops away and and value of TTL is one, then  next router in the path will decrement the TTL to 0, realize it cant forward the packet and will drop it. If the hop count value to reach neighbor is equal to or less than the maximum_hop_limit value then intermediate router decrements the TTL count by one and forwards the request to BGP neighour. If the hop count value is greater than the maximum_hop_limit value then intermediate router discards the request when TTL becomes 0.  | [optional] [default to 1]
**Enabled** | **bool** | Flag to enable this BGP Neighbor | [optional] [default to true]
**RemoteAsNum** | **string** | 4 Byte ASN of the neighbor in ASPLAIN/ASDOT Format | [optional] [default to null]
**AddressFamilies** | [**[]BgpNeighborAddressFamily**](BgpNeighborAddressFamily.md) | User can enable the neighbor for the specific address families and also define filters per address family. When the neighbor is created, it is default enabled for IPV4_UNICAST address family for backward compatibility reasons. User can change that if required, by defining the address family configuration.  | [optional] [default to null]
**BfdConfig** | [***BfdConfigParameters**](BfdConfigParameters.md) |  | [optional] [default to null]
**LogicalRouterId** | **string** | Logical router id | [optional] [default to null]
**FilterInIpprefixlistId** | **string** | This is a deprecated property, Please  use &#x27;address_family&#x27; instead. | [optional] [default to null]
**FilterOutRoutemapId** | **string** | This is a deprecated property, Please use &#x27;address_family&#x27; instead. | [optional] [default to null]
**FilterInRoutemapId** | **string** | This is a deprecated property, Please use &#x27;address_family&#x27; instead. | [optional] [default to null]
**KeepAliveTimer** | **int64** | Frequency (seconds) with which keep alive messages are sent to peers | [optional] [default to 60]
**Password** | **string** | User can create (POST) the neighbor with or without the password. The view (GET) on the neighbor, would never reveal if the password is set or not. The password can be set later using edit neighbor workFlow (PUT) On the edit neighbor (PUT), if the user does not specify the password property, the older value is retained. Maximum length of this field is 20 characters.  | [optional] [default to null]
**SourceAddress** | **string** | Deprecated - do not provide a value for this field. Use source_addresses instead. | [optional] [default to null]
**AllowAsIn** | **bool** | Flag to enable allowas_in option for BGP neighbor | [optional] [default to false]
**EnableBfd** | **bool** | Flag to enable BFD for this BGP Neighbor. Enable this if the neighbor supports BFD as this will lead to faster convergence. | [optional] [default to false]
**NeighborAddress** | **string** | Neighbor IP Address | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

