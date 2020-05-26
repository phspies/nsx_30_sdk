# ServiceManager

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
**Port** | **int64** | Integer port value to specify a standard/non-standard HTTPS port. | [default to null]
**ServiceIds** | [**[]ResourceReference**](ResourceReference.md) | The IDs of services, provided by partner. | [default to null]
**AuthenticationScheme** | [***CallbackAuthenticationScheme**](CallbackAuthenticationScheme.md) |  | [default to null]
**Thumbprint** | **string** | Thumbprint (SHA-256 hash represented in lower case hex) for the certificate on the partner console. This will be required to establish secure communication with the console and to avoid man-in-the-middle attacks. | [optional] [default to null]
**VendorId** | **string** | Id which is unique to a vendor or partner for which the service is created. | [optional] [default to null]
**Uri** | **string** | URI on which notification requests should be made on the specified server. | [default to null]
**Server** | **string** | IP address or fully qualified domain name of the partner REST server. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

