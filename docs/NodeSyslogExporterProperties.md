# NodeSyslogExporterProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**TlsCaPem** | **string** | CA certificate PEM of TLS server to export to | [optional] [default to null]
**Protocol** | **string** | Export protocol | [default to null]
**ExporterName** | **string** | Syslog exporter name | [default to null]
**Level** | **string** | Logging level to export | [default to null]
**TlsClientCaPem** | **string** | CA certificate PEM of the rsyslog client | [optional] [default to null]
**TlsCertPem** | **string** | Certificate PEM of the rsyslog client | [optional] [default to null]
**Server** | **string** | IP address or hostname of server to export to | [default to null]
**Facilities** | **[]string** | Facilities to export | [optional] [default to null]
**Msgids** | **[]string** | MSGIDs to export | [optional] [default to null]
**StructuredData** | **[]string** | Structured data to export | [optional] [default to null]
**Port** | **int64** | Port to export to | [optional] [default to 514]
**TlsKeyPem** | **string** | Private key PEM of the rsyslog client | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

