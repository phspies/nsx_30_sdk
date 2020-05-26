# MpaConfigProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**RmqClientType** | [***interface{}**](interface{}.md) | The nodes client type. | [optional] [default to null]
**RmqBrokerCluster** | [**[]BrokerProperties**](BrokerProperties.md) | The list of messaging brokers this controller is configured with. | [optional] [default to null]
**SharedSecret** | **string** | The shared secret to use when autnenticating to the management plane&#x27;s message bus. Not returned in REST responses. | [optional] [default to null]
**AccountName** | [***interface{}**](interface{}.md) | The account name to use when authenticating to the management plane&#x27;s message bus. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

