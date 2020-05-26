# IdsEventsSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**FirstOccurence** | **int64** | First occurence of the intrusion, in epoch milliseconds. | [optional] [default to null]
**LatestOccurence** | **int64** | Latest occurence of the intrusion, in epoch milliseconds. | [optional] [default to null]
**TotalCount** | **int64** | Number of times this particular signature was detected. | [optional] [default to null]
**UserDetails** | [***interface{}**](interface{}.md) | List of users logged into VMs on which a particular signature was detected. | [optional] [default to null]
**VmDetails** | [***interface{}**](interface{}.md) | List of VMs on which a particular signature was detected with the count. | [optional] [default to null]
**IsRuleValid** | **bool** | Indicates if the rule id is valid or not. | [optional] [default to null]
**SignatureMetadata** | [***interface{}**](interface{}.md) | Metadata about the detected signature including name, id, severity, product affected, protocol etc. | [optional] [default to null]
**IdsflowDetails** | [***interface{}**](interface{}.md) | IDS event flow data specific to each IDS event. The data includes source ip, source port, destination ip, destination port, and protocol. | [optional] [default to null]
**SignatureId** | **int64** | Signature ID pertaining to the detected intrusion. | [optional] [default to null]
**RuleId** | **int64** | The IDS Rule id that detected this particular intrusion. | [optional] [default to null]
**IsOngoing** | **bool** | Flag indicating an ongoing intrusion. | [optional] [default to null]
**AffectedVmCount** | **int64** | Count of VMs on which a particular signature was detected. | [optional] [default to null]
**ResourceType** | **string** | IDSEvent resource type. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

