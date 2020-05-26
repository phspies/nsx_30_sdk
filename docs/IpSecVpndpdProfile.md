# IpSecVpndpdProfile

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
**RetryCount** | **int64** | Maximum number of DPD messages retry attempts. This value is applicable for both dpd probe modes, periodic and on-demand. | [optional] [default to 5]
**Enabled** | **bool** | If true, enable dead peer detection. | [optional] [default to true]
**DpdProbeMode** | **string** | DPD probe mode is used to query the liveliness of the peer. Two modes are possible - PERIODIC - is used to query the liveliness of the peer at regular intervals (dpd_probe_interval). It does not take into consideration traffic coming from the peer. The benefit of this mode over the on-demand mode is earlier detection of dead peers. However, use of periodic DPD incurs extra overhead. When communicating to large numbers of peers, please consider using on-demand DPD instead. ON_DEMAND - is used to query the liveliness of the peer by instructing the local endpoint to send DPD message to a peer if there is traffic to send to the peer AND the peer was idle for dpd_probe_interval seconds (i.e. there was no traffic from the peer for dpd_probe_interval seconds)  | [optional] [default to DPD_PROBE_MODE.PERIODIC]
**DpdProbeInterval** | **int64** | When the DPD probe mode is periodic, this interval is the number of seconds between DPD messages. When the DPD probe mode is on-demand, this interval is the number of seconds during which traffic is not received from the peer before DPD retry messages are sent if there is IPSec traffic to send. For PERIODIC Mode:    Minimum: 3    Maximum: 360    Default: 60 For ON_DEMAND Mode:    Minimum: 1    Maximum: 10    Default: 3  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

