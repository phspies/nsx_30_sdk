# SnmpServiceProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V3AuthProtocol** | **string** | SNMP v3 auth protocol | [default to V3_AUTH_PROTOCOL.SHA1]
**Communities** | **[]string** | SNMP v1, v2c community strings | [optional] [default to null]
**V3Configured** | **bool** | SNMP v3 is configured or not | [optional] [default to null]
**V3PrivProtocol** | **string** | SNMP v3 private protocol | [default to V3_PRIV_PROTOCOL.AES128]
**V3Users** | [**[]SnmpV3User**](SnmpV3User.md) | SNMP v3 users information | [optional] [default to null]
**V2Configured** | **bool** | SNMP v2 is configured or not | [optional] [default to null]
**StartOnBoot** | **bool** | Start when system boots | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

