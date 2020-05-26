# NodeStatusProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**LoadAverage** | **[]float64** | One, five, and fifteen minute load averages for the system | [optional] [default to null]
**SwapUsed** | **int64** | Amount of swap disk in use, in kilobytes | [optional] [default to null]
**CpuUsage** | [***CpuUsage**](CpuUsage.md) |  | [optional] [default to null]
**NonDpdkCpuCores** | **int64** | Number of non-DPDK cores on Edge Node. | [optional] [default to null]
**SwapTotal** | **int64** | Amount of disk available for swap, in kilobytes | [optional] [default to null]
**SystemTime** | **int64** | Current time expressed in milliseconds since epoch | [optional] [default to null]
**CpuCores** | **int64** | Number of CPU cores on the system | [optional] [default to null]
**Uptime** | **int64** | Milliseconds since system start | [optional] [default to null]
**MemCache** | **int64** | Amount of RAM on the system that can be flushed out to disk, in kilobytes | [optional] [default to null]
**MemTotal** | **int64** | Amount of RAM allocated to the system, in kilobytes | [optional] [default to null]
**MemUsed** | **int64** | Amount of RAM in use on the system, in kilobytes | [optional] [default to null]
**DpdkCpuCores** | **int64** | Number of DPDK cores on Edge Node which are used for packet IO processing. | [optional] [default to null]
**FileSystems** | [**[]NodeFileSystemProperties**](NodeFileSystemProperties.md) | File systems configured on the system | [optional] [default to null]
**Source** | **string** | Source of status data. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

