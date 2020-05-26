# SwitchSecuritySwitchingProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredCapabilities** | **[]string** |  | [optional] [default to null]
**ResourceType** | **string** |  | [default to null]
**BpduFilter** | [***BpduFilter**](BpduFilter.md) |  | [optional] [default to null]
**RateLimits** | [***RateLimits**](RateLimits.md) |  | [optional] [default to null]
**RaGuardEnabled** | **bool** | RA Guard when enabled blocks unauthorized/rogue Router Advertisement (RA) packets. | [optional] [default to true]
**DhcpFilter** | [***DhcpFilter**](DhcpFilter.md) |  | [optional] [default to null]
**BlockNonIpTraffic** | **bool** | A flag to block all traffic except IP/(G)ARP/BPDU | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

