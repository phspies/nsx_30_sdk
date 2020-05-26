# Snmpv3Properties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthProtocol** | **string** | Authentication protocol used for SNMP v3 communication. | [optional] [default to AUTH_PROTOCOL.SHA1]
**PrivProtocol** | **string** | Privacy protocol used for SNMP v3 communication. | [optional] [default to PRIV_PROTOCOL.AES128]
**Users** | [**[]Snmpv3User**](Snmpv3User.md) | List of SNMP v3 users allowed to poll NSX nodes over SNMP. Also, users specified in a SNMP v3 target must exist in this list. | [optional] [default to null]
**Targets** | [**[]Snmpv3Target**](Snmpv3Target.md) | List of SNMP v3 targets/receivers where SNMP v3 traps/notifications will be sent from NSX nodes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

