# ConditionalForwarderZone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpstreamServers** | **[]string** | Ip address of the upstream DNS servers the DNS forwarder accesses.  | [default to null]
**SourceIp** | **string** | The source ip used by the fowarder of the zone. If no source ip specified, the ip address of listener of the DNS forwarder will be used.  | [optional] [default to null]
**DomainNames** | **[]string** | A forwarder domain name should be a valid FQDN. If reverse lookup is needed for this zone, reverse lookup domain name like X.in-addr.arpa can be defined. Here the X represents a subnet.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

