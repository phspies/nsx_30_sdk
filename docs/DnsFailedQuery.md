# DnsFailedQuery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeSpent** | **int64** | The time the query took before it got a failed answer, in ms. | [optional] [default to null]
**RecordType** | **string** | The record type be queried, e.g. A, CNAME, SOA, etc. | [optional] [default to null]
**ClientIp** | **string** | The client host ip address from which the query was issued.  | [optional] [default to null]
**UpstreamServerIp** | **string** | The upstream server ip address to which the query was forwarded. If the query could not be serviced from the DNS forwarder cache, this property will contain the IP address of the DNS server that serviced the request. If the request was serviced from the cache, this property will be absent.  | [optional] [default to null]
**ErrorMessage** | **string** | The detailed error message of the failed query, if any. | [optional] [default to null]
**Address** | **string** | The address be queried, can be a FQDN or an ip address. | [optional] [default to null]
**Timestamp** | **string** | Timestamp of the request, in YYYY-MM-DD HH:MM:SS.zzz format. | [default to null]
**ErrorType** | **string** | The type of the query failure, e.g. NXDOMAIN, etc. | [optional] [default to null]
**SourceIp** | **string** | The source ip address that is used to forward a query to an upstream server.  | [optional] [default to null]
**ForwarderIp** | **string** | The DNS forwarder ip address to which the query was first received.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

