# IdsEventsBySignature

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Count** | **int64** | Number of times this particular signature was detected. | [optional] [default to null]
**FirstOccurence** | **int64** | First occurence of the intrusion, in epoch milliseconds. | [optional] [default to null]
**Severity** | **string** | Severity of the threat covered by the signature, can be Critical, High, Medium, or Low. | [optional] [default to null]
**SignatureName** | **string** | Name of the signature pertaining to the detected intrusion. | [optional] [default to null]
**IsOngoing** | **bool** | Flag indicating an ongoing intrusion. | [optional] [default to null]
**SignatureId** | **int64** | Signature ID pertaining to the detected intrusion. | [optional] [default to null]
**ResourceType** | **string** | IDSEvent resource type. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

