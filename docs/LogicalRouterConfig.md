# LogicalRouterConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalTransitNetwork** | **string** | CIDR block defining service router to distributed router links. If the value for this field is not provided, then it will be considered as default IPv4 CIDR- \&quot;169.254.0.0/28\&quot; for logical router with ACTIVE_STANDBY HA mode \&quot;169.254.0.0/24\&quot; for logical router with ACTIVE_ACTIVE HA mode  | [optional] [default to null]
**TransportZoneId** | **string** | Transport zone of the logical router. If specified then all downlink switches should belong to this transport zone and an error will be thrown if transport zone of the downlink switch doesn&#x27;t match with this transport zone. All internal and external transit switches will be created in this transport zone. | [optional] [default to null]
**HaVipConfigs** | [**[]HaVipConfig**](HaVipConfig.md) | This configuration can be defined only for Active-Standby LogicalRouter to provide | redundancy. For mulitple uplink ports, multiple HaVipConfigs must be defined | and each config will pair exactly two uplink ports. The VIP will move and will | always be owned by the Active node. Note - when HaVipConfig[s] are defined, | configuring dynamic-routing is disallowed. | [optional] [default to null]
**ExternalTransitNetworks** | **[]string** | CIDR block defining addresses for Tier0 to Tier1 links. If the value for this field is not provided, then it will be considered as default IPv4 CIDR \&quot;100.64.0.0/16\&quot;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

