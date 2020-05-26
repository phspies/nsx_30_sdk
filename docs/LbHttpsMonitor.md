# LbHttpsMonitor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorPort** | **string** | If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported. For ICMP monitor, monitor_port is not required.  | [optional] [default to null]
**FallCount** | **int64** | num of consecutive checks must fail before marking it down | [optional] [default to 3]
**Interval** | **int64** | the frequency at which the system issues the monitor check (in second) | [optional] [default to 5]
**RiseCount** | **int64** | num of consecutive checks must pass before marking it up | [optional] [default to 3]
**Timeout** | **int64** | the number of seconds the target has in which to respond to the monitor request  | [optional] [default to 15]
**ResponseStatusCodes** | **[]int32** | The HTTP response status code should be a valid HTTP status code.  | [optional] [default to null]
**ServerAuthCrlIds** | **[]string** | A Certificate Revocation List (CRL) can be specified in the server-side SSL profile binding to disallow compromised server certificates.  | [optional] [default to null]
**ServerAuthCaIds** | **[]string** | If server auth type is REQUIRED, server certificate must be signed by one of the trusted Certificate Authorities (CAs), also referred to as root CAs, whose self signed certificates are specified.  | [optional] [default to null]
**ServerAuth** | **string** | server authentication mode | [optional] [default to SERVER_AUTH.IGNORE]
**RequestBody** | **string** | String to send as part of HTTP health check request body. Valid only for certain HTTP methods like POST.  | [optional] [default to null]
**ResponseBody** | **string** | If HTTP response body match string (regular expressions not supported) is specified (using LbHttpMonitor.response_body) then the healthcheck HTTP response body is matched against the specified string and server is considered healthy only if there is a match. If the response body string is not specified, HTTP healthcheck is considered successful if the HTTP response status code is 2xx, but it can be configured to accept other status codes as successful.  | [optional] [default to null]
**Ciphers** | **[]string** | supported SSL cipher list to servers | [optional] [default to null]
**RequestHeaders** | [**[]LbHttpRequestHeader**](LbHttpRequestHeader.md) | Array of HTTP request headers | [optional] [default to null]
**ClientCertificateId** | **string** | client certificate can be specified to support client authentication.  | [optional] [default to null]
**RequestMethod** | **string** | the health check method for HTTP monitor type | [optional] [default to REQUEST_METHOD.GET]
**IsFips** | **bool** | This flag is set to true when all the ciphers and protocols are FIPS compliant. It is set to false when one of the ciphers or protocols are not FIPS compliant..  | [optional] [default to null]
**CertificateChainDepth** | **int64** | authentication depth is used to set the verification depth in the server certificates chain.  | [optional] [default to 3]
**IsSecure** | **bool** | This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure.  | [optional] [default to null]
**RequestUrl** | **string** | URL used for HTTP monitor | [optional] [default to null]
**CipherGroupLabel** | **string** | It is a label of cipher group which is mostly consumed by GUI.  | [optional] [default to null]
**RequestVersion** | **string** | HTTP request version | [optional] [default to REQUEST_VERSION.11_]
**Protocols** | **[]string** | SSL versions TLS1.1 and TLS1.2 are supported and enabled by default. SSLv2, SSLv3, and TLS1.0 are supported, but disabled by default.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

