# ContainerApplicationInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**Status** | **string** | Status of the container application instance. | [optional] [default to null]
**NetworkStatus** | **string** | Network status of container application instance. | [optional] [default to null]
**ContainerClusterId** | **string** | Identifier of the container cluster this application instance belongs to. | [optional] [default to null]
**ClusterNodeId** | **string** | Cluster node id where application instance is running. | [optional] [default to null]
**ExternalId** | **string** | Identifier of the container application instance on container cluster. | [default to null]
**OriginProperties** | [**[]KeyValuePair**](KeyValuePair.md) | Array of additional specific properties of container application instance in key-value format.  | [optional] [default to null]
**ContainerApplicationIds** | **[]string** | List of identifiers of the container application. | [optional] [default to null]
**ContainerProjectId** | **string** | Identifier of the container project which this container application instance belongs to.  | [optional] [default to null]
**NetworkErrors** | [**[]NetworkError**](NetworkError.md) | List of network errors related to container application instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

