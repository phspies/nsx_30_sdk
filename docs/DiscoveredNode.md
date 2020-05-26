# DiscoveredNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**Stateless** | **bool** | The stateless property describes whether host persists its state across reboot or not. If state persists, value is set as false otherwise true. | [optional] [default to null]
**ParentComputeCollection** | **string** | External id of the compute collection to which this node belongs | [optional] [default to null]
**Certificate** | **string** | Certificate of the discovered node | [optional] [default to null]
**OriginId** | **string** | Id of the compute manager from where this node was discovered | [optional] [default to null]
**IpAddresses** | **[]string** | IP Addresses of the the discovered node. | [optional] [default to null]
**HardwareId** | **string** | Hardware Id is generated using system hardware info. It is used to retrieve fabric node of the esx. | [optional] [default to null]
**OsVersion** | **string** | OS version of the discovered node | [optional] [default to null]
**NodeType** | **string** | Discovered Node type like Host | [optional] [default to null]
**OsType** | **string** | OS type of the discovered node | [optional] [default to null]
**OriginProperties** | [**[]KeyValuePair**](KeyValuePair.md) | Key-Value map of additional specific properties of discovered node in the Compute Manager  | [optional] [default to null]
**ExternalId** | **string** | External id of the discovered node, ex. a mo-ref from VC | [optional] [default to null]
**CmLocalId** | **string** | Local Id of the discovered node in the Compute Manager | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

