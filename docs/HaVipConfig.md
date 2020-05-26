# HaVipConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HaVipSubnets** | [**[]VipSubnet**](VIPSubnet.md) | Array of IP address subnets which will be used as floating IP addresses. | Note - this configuration is applicable only for Active-Standby LogicalRouter. | For Active-Active LogicalRouter this configuration will be rejected. | [default to null]
**RedundantUplinkPortIds** | **[]string** | Identifiers of logical router uplink ports which are to be paired to provide | redundancy. Floating IP will be owned by one of these uplink ports (depending upon | which node is Active). | [default to null]
**Enabled** | **bool** | Flag to enable this ha vip config. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

