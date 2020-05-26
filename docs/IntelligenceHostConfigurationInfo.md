# IntelligenceHostConfigurationInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | Unique identifier of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**ContextDataCollectionInterval** | **int64** | Interval in minute of reporting VM guest context data to NSX-Intelligence. Recommend to keep this value the same as flow_data_collection_interval.  | [optional] [default to null]
**BrokerTruststore** | **string** | A truststore to establish the trust between NSX and NSX-Intelligence brokers.  | [optional] [default to null]
**BrokerCertificate** | **string** | A broker certificate to verify the identity of brokers.  | [optional] [default to null]
**EnableContextDataCollection** | **bool** | Enable NSX-Intelligence context data collection in host nodes.  | [optional] [default to null]
**ContextUserUids** | **[]string** | List of linux user uid to collect context data. Empty implies all users.  | [optional] [default to null]
**EnableFlowDataCollection** | **bool** | Enable NSX-Intelligence flow data collection in host nodes.  | [optional] [default to null]
**ContextProcessHashes** | **[]string** | List of hashes of processes to collect context data. Empty implies all processes.  | [optional] [default to null]
**EnableDataCollection** | **bool** | Enable NSX-Intelligence data collection in host nodes.  This property has been deprecated. To enable flow data collection, use property enable_flow_data_collection instead. To enable context data collection, use property enable_context_data_collection instead.  When this property is set to false, no data collection is performed even if enable_flow_data_collection or enable_context_data_collection is set to true.  When this property is set to true, property enable_flow_data_collection and enable_context_data_collection control whether to collect flow data and context data separately.  | [optional] [default to null]
**ContextProcessNames** | **[]string** | List of processes to collect context data. Empty implies all processes.  | [optional] [default to null]
**PrivateIpPrefix** | [**[]IntelligenceFlowPrivateIpPrefixInfo**](IntelligenceFlowPrivateIpPrefixInfo.md) | List of private IP prefix that NSX-Intelligence network flow is collected from.  | [optional] [default to null]
**BrokerBootstrapServers** | [**[]IntelligenceBrokerEndpointInfo**](IntelligenceBrokerEndpointInfo.md) | List of NSX-Intelligence broker endpoints that host nodes contact initially.  | [optional] [default to null]
**MaxInactiveFlowCount** | **int64** | Maximum inactive network flow to collect in collection interval.  | [optional] [default to null]
**ContextUserSids** | **[]string** | List of windows user sid to collect context data. Empty implies all users.  | [optional] [default to null]
**FlowDataCollectionInterval** | **int64** | Interval in minute of reporting network flow data to NSX-Intelligence. Recommend to keep this value the same as context_data_collection_interval.  | [optional] [default to null]
**MaxActiveFlowCount** | **int64** | Maximum active network flow to collect in collection interval.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

