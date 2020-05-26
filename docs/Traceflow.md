# Traceflow

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
**Id** | **string** | The id of the traceflow round | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**OperationState** | **string** | Represents the traceflow operation state | [optional] [default to null]
**LogicalCounters** | [***TraceflowObservationCounters**](TraceflowObservationCounters.md) |  | [optional] [default to null]
**Timeout** | **int64** | Maximum time (in ms) the management plane will be waiting for this traceflow round. | [optional] [default to null]
**ResultOverflowed** | **bool** | A flag, when set true, indicates some observations were deleted from the result set. | [optional] [default to null]
**LportId** | **string** | id of the source logical port used for injecting the traceflow packet | [optional] [default to null]
**Counters** | [***TraceflowObservationCounters**](TraceflowObservationCounters.md) |  | [optional] [default to null]
**RequestStatus** | **string** | The status of the traceflow RPC request. SUCCESS - The traceflow request is sent successfully. TIMEOUT - The traceflow request gets timeout. SOURCE_PORT_NOT_FOUND - The source port of the request cannot be found. DATA_PATH_NOT_READY - The datapath component cannot be ready to receive request. CONNECTION_ERROR - There is connection error on datapath component. UNKNOWN - The status of traceflow request cannot be determined. | [optional] [default to null]
**Analysis** | **[]string** | Traceflow result analysis notes | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

