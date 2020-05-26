# ApiServiceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalApiConcurrencyLimit** | **int64** | The maximum number of concurrent API requests that will be serviced. If the number of API requests being processed exceeds this limit, new API requests will be refused and a 503 Service Unavailable response will be returned to the client.  To disable API concurrency limiting, set this value to 0. | [optional] [default to 199]
**LockoutImmuneAddresses** | **[]string** | The list of IP addresses which are not subjected to a lockout on failed login attempts. | [optional] [default to null]
**BasicAuthenticationEnabled** | **bool** | Identifies whether basic authentication is enabled or disabled in API calls. | [optional] [default to true]
**RedirectHost** | **string** | Host name or IP address to use for redirect location headers, or empty string to derive from current request. To disable, set redirect_host to the empty string (\&quot;\&quot;). | [optional] 
**CipherSuites** | [**[]CipherSuite**](CipherSuite.md) | The TLS cipher suites that the API service will negotiate. | [optional] [default to null]
**CookieBasedAuthenticationEnabled** | **bool** | Identifies whether cookie-based authentication is enabled or disabled in API calls. When cookie-based authentication is disabled, new sessions cannot be created via /api/session/create. | [optional] [default to true]
**SessionTimeout** | **int64** | NSX session inactivity timeout | [optional] [default to 1800]
**ClientApiRateLimit** | **int64** | The maximum number of API requests that will be serviced per second for a given authenticated client.  If more API requests are received than can be serviced, a 429 Too Many Requests HTTP response will be returned. To disable API rate limiting, set this value to 0. | [optional] [default to 100]
**ClientApiConcurrencyLimit** | **int64** | The maximum number of concurrent API requests that will be serviced for a given authenticated client.  If the number of API requests being processed exceeds this limit, new API requests will be refused and a 503 Service Unavailable response will be returned to the client. To disable API concurrency limiting, set this value to 0. | [optional] [default to 40]
**ProtocolVersions** | [**[]ProtocolVersion**](ProtocolVersion.md) | The TLS protocol versions that the API service will negotiate. | [optional] [default to null]
**ConnectionTimeout** | **int64** | NSX connection timeout, in seconds. To disable timeout, set to 0. | [optional] [default to 30]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

