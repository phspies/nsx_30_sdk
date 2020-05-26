# IpSecVpnTunnelProfile

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
**EncapsulationMode** | **string** | Encapsulation Mode to be used for encryption of packet. Tunnel mode protects internal routing information by encrypting IP header of original packet. | [optional] [default to ENCAPSULATION_MODE.MODE]
**TransformProtocol** | **string** | IPSec transform specifies IPSec security protocol. | [optional] [default to TRANSFORM_PROTOCOL.ESP]
**DigestAlgorithms** | **[]string** | Algorithm to be used for message digest. Default digest algorithm is implicitly covered by default encryption algorithm \&quot;AES_GCM_128\&quot;. | [optional] [default to null]
**EncryptionAlgorithms** | **[]string** | Encryption algorithm to encrypt/decrypt the messages exchanged between IPSec VPN initiator and responder during tunnel negotiation. Default is AES_GCM_128. | [optional] [default to null]
**EnablePerfectForwardSecrecy** | **bool** | If true, perfect forward secrecy (PFS) is enabled. | [optional] [default to true]
**DhGroups** | **[]string** | Diffie-Hellman group to be used if PFS is enabled. Default is GROUP14. | [optional] [default to null]
**DfPolicy** | **string** | Defragmentation policy helps to handle defragmentation bit present in the inner packet. COPY copies the defragmentation bit from the inner IP packet into the outer packet. CLEAR ignores the defragmentation bit present in the inner packet. | [optional] [default to DF_POLICY.COPY]
**SaLifeTime** | **int64** | SA life time specifies the expiry time of security association. Default is 3600 seconds.  | [optional] [default to 3600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

