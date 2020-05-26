# SourceNsxApiEndpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcPort** | **int32** | VC port that will be used to fetch details. | [optional] [default to 443]
**VcUsername** | **string** | Username for connecting to VC. | [default to null]
**VcIp** | **string** | IP address or host name of VC. | [default to null]
**Ip** | **string** | IP address or hostname of a source NSX API endpoint. This field is not applicable in case of vSphere network migration. | [optional] [default to null]
**AuthToken** | **string** | Auth token used to make REST calls to source NSX API endpoint. This field is not applicable in case of vSphere network migration. | [optional] [default to null]
**NsxSyncrole** | **string** | Signifies Universal Sync role status (STANDALONE, PRIMARY, SECONDARY) of a source NSX API endpoint. | [optional] [default to null]
**VcVersion** | **string** | Build version of VC. | [optional] [default to null]
**NsxUsername** | **string** | Username for connecting to NSX manager. This field is not applicable in case of vSphere network migration. | [optional] [default to null]
**NsxVersion** | **string** | Build version (major, minor, patch) of a source NSX API endpoint. | [optional] [default to null]
**NsxPassword** | **string** | Password for connecting to NSX manager. This field is not applicable in case of vSphere network migration. | [optional] [default to null]
**VcPassword** | **string** | Password for connecting to VC. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

