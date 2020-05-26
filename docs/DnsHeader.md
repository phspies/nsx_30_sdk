# DnsHeader

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType** | **string** | This is used to specify the type of the address. V4 - The address provided is an IPv4 domain name/IP address, the Type in query or response will be A V6 - The address provided is an IPv6 domain name/IP address, the Type in query or response will be AAAA | [optional] [default to ADDRESS_TYPE.V4]
**MessageType** | **string** | Specifies the message type whether it is a query or a response. | [optional] [default to MESSAGE_TYPE.QUERY]
**Address** | **string** | This is used to define what is being asked or responded. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

