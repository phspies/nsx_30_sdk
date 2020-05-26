# IpSecVpnPeerEndpoint

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
**Psk** | **string** | IPSec Pre-shared key. Maximum length of this field is 128 characters. | [optional] [default to null]
**PeerId** | **string** | Peer identifier. | [default to null]
**IpsecTunnelProfileId** | **string** | Tunnel profile id to be used. By default it will point to system default profile. | [optional] [default to null]
**AuthenticationMode** | **string** | Authentication mode used for the peer authentication. For PSK (Pre Shared Key) authentication mode, &#x27;psk&#x27; property is mandatory and for the CERTIFICATE authentication mode, &#x27;peer_id&#x27; property is mandatory. | [optional] [default to AUTHENTICATION_MODE.PSK]
**PeerAddress** | **string** | IPV4 address of peer endpoint on remote site. | [default to null]
**ConnectionInitiationMode** | **string** | Connection initiation mode used by local endpoint to establish ike connection with peer endpoint. INITIATOR - In this mode local endpoint initiates tunnel setup and will also respond to incoming tunnel setup requests from peer gateway. RESPOND_ONLY - In this mode, local endpoint shall only respond to incoming tunnel setup requests. It shall not initiate the tunnel setup. ON_DEMAND - In this mode local endpoint will initiate tunnel creation once first packet matching the policy rule is received and will also respond to incoming initiation request.  | [optional] [default to CONNECTION_INITIATION_MODE.INITIATOR]
**DpdProfileId** | **string** | Dead peer detection (DPD) profile id. Default will be set according to system default policy. | [optional] [default to null]
**IkeProfileId** | **string** | IKE profile id to be used. Default will be set according to system default policy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

