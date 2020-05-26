# NodeProcessProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**MemUsed** | **int64** | Virtual memory used by process in bytes | [optional] [default to null]
**CpuTime** | **int64** | CPU time (user and system) consumed by process in milliseconds | [optional] [default to null]
**Ppid** | **int64** | Parent process id | [optional] [default to null]
**StartTime** | **int64** | Process start time expressed in milliseconds since epoch | [optional] [default to null]
**ProcessName** | **string** | Process name | [optional] [default to null]
**Pid** | **int64** | Process id | [optional] [default to null]
**Uptime** | **int64** | Milliseconds since process started | [optional] [default to null]
**MemResident** | **int64** | Resident set size of process in bytes | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

