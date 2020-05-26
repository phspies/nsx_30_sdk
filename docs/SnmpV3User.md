# Snmpv3User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivPassword** | **string** | Privacy password used for SNMP v3 communication. This field is required when adding a user. When updating a user, do not include this field in the request. If this field is present in an update request, it will be considered as a new value for privacy password. | [optional] [default to null]
**Access** | **string** | Access permissions for polling NSX nodes over SNMP v3. | [optional] [default to ACCESS.ONLY]
**AuthPassword** | **string** | Authentication password used for SNMP v3 communication. This field is required when adding a user. When updating a user, do not include this field in the request. If this field is present in an update request, it will be considered as a new value for authentication password. | [optional] [default to null]
**UserId** | **string** | Unique SNMP v3 user id. | [default to null]
**SecurityLevel** | **string** | Security level indicates whether SNMP communication involves authentication and privacy protocols for this user. Value \&quot;AUTH_PRIV\&quot; indicates both authentication and privacy protocols will be used for SNMP communication. | [optional] [default to SECURITY_LEVEL.PRIV]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

