# IpAddressElement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** | The _revision property describes the current revision of the resource. To prevent clients from overwriting each other&#x27;s changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. | [optional] [default to null]
**IpAddress** | **string** | IPElement can be a single IP address, IP address range or a Subnet. Its type can be of IPv4 or IPv6. Supported list of formats are \&quot;192.168.1.1\&quot;, \&quot;192.168.1.1-192.168.1.100\&quot;, \&quot;192.168.0.0/24\&quot;, \&quot;fe80::250:56ff:fe83:318c\&quot;, \&quot;fe80::250:56ff:fe83:3181-fe80::250:56ff:fe83:318c\&quot;, \&quot;fe80::250:56ff:fe83:318c/64\&quot;  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

