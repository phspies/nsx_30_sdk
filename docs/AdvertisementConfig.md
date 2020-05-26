# AdvertisementConfig

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
**AdvertiseNsxConnectedRoutes** | **bool** | Flag to advertise all connected routes | [optional] [default to false]
**AdvertiseLbVip** | **bool** | Flag to advertise lb vip ips | [optional] [default to false]
**AdvertiseStaticRoutes** | **bool** | Flag to advertise all static routes | [optional] [default to false]
**LogicalRouterId** | **string** | TIER1 logical router id on which to enable this configuration | [optional] [default to null]
**AdvertiseDnsForwarder** | **bool** | Flag to advertise all routes of dns forwarder listener ips and source ips | [optional] [default to false]
**AdvertiseNatRoutes** | **bool** | Flag to advertise all routes of nat | [optional] [default to false]
**AdvertiseIpsecLocalIp** | **bool** | Flag to advertise all IPSec VPN local endpoint ips to linked TIER0 logical router | [optional] [default to false]
**Enabled** | **bool** | Flag to enable this configuration | [optional] [default to false]
**AdvertiseLbSnatIp** | **bool** | Flag to advertise all lb SNAT ips | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

