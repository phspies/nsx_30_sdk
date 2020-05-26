# NodeProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**SystemTime** | **int64** | Current time expressed in milliseconds since epoch | [optional] [default to null]
**NodeUuid** | **string** | Node Unique Identifier | [optional] [default to null]
**Motd** | [***interface{}**](interface{}.md) | Message of the day to display when users login to node using the NSX CLI | [optional] [default to null]
**CliTimeout** | **int64** | NSX CLI inactivity timeout, set to 0 to configure no timeout | [optional] [default to null]
**KernelVersion** | **string** | Kernel version | [optional] [default to null]
**ExportType** | **string** | Export restrictions in effect, if any | [optional] [default to null]
**Hostname** | **string** | Host name or fully qualified domain name of node | [optional] [default to null]
**ProductVersion** | **string** | Product version | [optional] [default to null]
**NodeVersion** | **string** | Node version | [optional] [default to null]
**SystemDatetime** | **string** | System date time in UTC | [optional] [default to null]
**FullyQualifiedDomainName** | **string** | Fully qualified domain name | [optional] [default to null]
**Timezone** | **string** | Timezone | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

