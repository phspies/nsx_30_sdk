# EdgeNodeSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchDomains** | **[]string** | List of domain names that are used to complete unqualified host names.  | [optional] [default to null]
**DnsServers** | **[]string** | List of DNS servers.  | [optional] [default to null]
**NtpServers** | **[]string** | List of NTP servers.  | [optional] [default to null]
**Hostname** | **string** | Host name or FQDN for edge node. | [optional] [default to null]
**EnableSsh** | **bool** | Enabling SSH service is not recommended for security reasons.  | [optional] [default to false]
**AllowSshRootLogin** | **bool** | Allowing root SSH logins is not recommended for security reasons. Edit of this property is not supported when updating transport node. Use the CLI to change this property.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

