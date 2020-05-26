# IpMirrorDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Resource types of mirror destination | [default to null]
**DestinationIps** | **[]string** | The destination IPs of the mirror packet will be sent to. | [default to null]
**EncapsulationType** | **string** | You can choose GRE, ERSPAN II or ERSPAN III. | [default to ENCAPSULATION_TYPE.GRE]
**ErspanId** | **int32** | Used by physical switch for the mirror traffic forwarding. Must be provided and only effective when encapsulation type is ERSPAN type II or type III.  | [optional] [default to null]
**GreKey** | **int32** | User-configurable 32-bit key only for GRE | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

