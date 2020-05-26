# PrincipalIdentityWithCertificate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsProtected** | **bool** | Indicator whether the entities created by this principal should be protected. | [optional] [default to null]
**Role** | **string** | Role | [optional] [default to null]
**Name** | **string** | Name of the principal. | [default to null]
**PermissionGroup** | **string** | Use the &#x27;role&#x27; field instead and pass in &#x27;auditor&#x27; for read_only_api_users or &#x27;enterprise_admin&#x27; for the others. | [optional] [default to null]
**CertificateId** | **string** | Id of the stored certificate. When used with the deprecated POST /trust-management/principal-identities API this field is required. | [optional] [default to null]
**NodeId** | **string** | Unique node-id of a principal. | [default to null]
**CertificatePem** | **string** | PEM encoding of the new certificate. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

