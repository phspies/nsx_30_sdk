# NotificationWatcher

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Optional description that can be associated with this NotificationWatcher. | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | System generated identifier to identify a notification watcher uniquely.  | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**SendTimeout** | **int64** | Optional time duration (in seconds) to specify request timeout to notification watcher. If the send reaches the timeout, will try to send refresh_needed as true in the next time interval. The default value is 30 seconds. | [optional] [default to 30]
**Uri** | **string** | URI notification requests should be made on the specified server. | [default to null]
**CertificateSha256Thumbprint** | **string** | Contains the hex-encoded SHA256 thumbprint of the HTTPS certificate. It must be specified if use_https is set to true. | [optional] [default to null]
**Method** | **string** | Type of method notification requests should be made on the specified server. The value must be set to POST. | [default to null]
**SendInterval** | **int64** | Optional time interval (in seconds) for which notification URIs will be accumulated. At the end of the time interval the accumulated notification URIs will be sent to this NotificationWatcher in the form of zero (nothing accumulated) or more notification requests as soon as possible. If it is not specified, the NotificationWatcher should expected to receive notifications at any time. | [optional] [default to null]
**MaxSendUriCount** | **int64** | If the number of notification URIs accumulated in specified send_interval exceeds max_send_uri_count, then multiple notification requests (each with max_send_uri_count or less number of notification URIs) will be sent to this NotificationWatcher. The default value is 5000. | [optional] [default to 5000]
**AuthenticationScheme** | [***NotificationAuthenticationScheme**](NotificationAuthenticationScheme.md) |  | [default to null]
**Server** | **string** | IP address or fully qualified domain name of the partner/customer watcher. | [default to null]
**Port** | **int64** | Optional integer port value to specify a non-standard HTTP or HTTPS port. | [optional] [default to null]
**UseHttps** | **bool** | Optional field, when set to true indicates REST API server should use HTTPS. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

