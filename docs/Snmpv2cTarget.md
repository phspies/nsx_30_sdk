# Snmpv2cTarget

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunityName** | **string** | Unique non-sensitive community name to identify community. | [default to null]
**CommunityString** | **string** | Community string (shared secret). This field is required when adding a community target. When updating a community target, do not include this field in the request. If this field is present in an update request, it will be considered as a new value for community string. | [optional] [default to null]
**Port** | **int64** | SNMP v2c target server&#x27;s port number. | [optional] [default to 162]
**Server** | **string** | SNMP v2c target server&#x27;s IP or FQDN. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

