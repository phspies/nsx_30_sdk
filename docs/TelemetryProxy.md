# TelemetryProxy

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
**Username** | **string** | Specify the user name used to authenticate with the proxy server, if required.  | [optional] [default to null]
**Password** | **string** | Specify the password used to authenticate with the proxy server, if required. A GET call on /telemetry/config returns a non-meaningful password to maintain security. To change the password to a new value, issue a PUT call after updating this field. To remove the password, issue a PUT call after emptying this field. To retain a previously set password, issue a PUT call keeping the non-meaningful value obtained from the GET call.  | [optional] [default to null]
**Scheme** | **string** | The scheme accepted by the proxy server. Specify one of HTTP and HTTPS.  | [default to null]
**Hostname** | **string** | Specify the fully qualified domain name, or ip address, of the proxy server.  | [default to null]
**Port** | **int32** | Specify the port of the proxy server. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

