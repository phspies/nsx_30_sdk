# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptEULA**](SystemAdministrationSettingsApi.md#AcceptEULA) | **Post** /eula/accept | Accept end user license agreement 
[**AddCertificateImport**](SystemAdministrationSettingsApi.md#AddCertificateImport) | **Post** /trust-management/certificates?action&#x3D;import | Add a New Certificate
[**AddCrlImport**](SystemAdministrationSettingsApi.md#AddCrlImport) | **Post** /trust-management/crls?action&#x3D;import | Add a New Certificate Revocation List
[**AddNodeUserSshKeyAddSshKey**](SystemAdministrationSettingsApi.md#AddNodeUserSshKeyAddSshKey) | **Post** /node/users/{userid}/ssh-keys?action&#x3D;add_ssh_key | Add SSH public key to authorized_keys file for node user
[**AddOidcEndPoint**](SystemAdministrationSettingsApi.md#AddOidcEndPoint) | **Post** /trust-management/oidc-uris | Add an OpenID Connect end-point.
[**CollectSupportBundlesCollect**](SystemAdministrationSettingsApi.md#CollectSupportBundlesCollect) | **Post** /administration/support-bundles?action&#x3D;collect | Collect support bundles from registered cluster and fabric nodes
[**CreateCrlDistributionPoint**](SystemAdministrationSettingsApi.md#CreateCrlDistributionPoint) | **Post** /trust-management/crl-distribution-points | Create a Crl Distribution Point
[**CreateLicense**](SystemAdministrationSettingsApi.md#CreateLicense) | **Post** /licenses | Add a new license key
[**CreateOrUpdateLdapIdentitySource**](SystemAdministrationSettingsApi.md#CreateOrUpdateLdapIdentitySource) | **Put** /aaa/ldap-identity-sources/{ldap-identity-source-id} | Update an existing LDAP identity source
[**CreateRegistrationToken**](SystemAdministrationSettingsApi.md#CreateRegistrationToken) | **Post** /aaa/registration-token | Create registration access token
[**CreateRoleBinding**](SystemAdministrationSettingsApi.md#CreateRoleBinding) | **Post** /aaa/role-bindings | Assign roles to User or Group
[**DeleteAllStaleRoleBindingsDeleteStaleBindings**](SystemAdministrationSettingsApi.md#DeleteAllStaleRoleBindingsDeleteStaleBindings) | **Post** /aaa/role-bindings?action&#x3D;delete_stale_bindings | Delete all stale role assignments
[**DeleteCertificate**](SystemAdministrationSettingsApi.md#DeleteCertificate) | **Delete** /trust-management/certificates/{cert-id} | Delete Certificate for the Given Certificate ID
[**DeleteCrl**](SystemAdministrationSettingsApi.md#DeleteCrl) | **Delete** /trust-management/crls/{crl-id} | Delete a CRL
[**DeleteCrlDistributionPoint**](SystemAdministrationSettingsApi.md#DeleteCrlDistributionPoint) | **Delete** /trust-management/crl-distribution-points/{crl-distribution-point-id} | Delete a CrlDistributionPoint
[**DeleteCsr**](SystemAdministrationSettingsApi.md#DeleteCsr) | **Delete** /trust-management/csrs/{csr-id} | Delete a CSR
[**DeleteLdapIdentitySource**](SystemAdministrationSettingsApi.md#DeleteLdapIdentitySource) | **Delete** /aaa/ldap-identity-sources/{ldap-identity-source-id} | Delete an LDAP identity source
[**DeleteLicense**](SystemAdministrationSettingsApi.md#DeleteLicense) | **Delete** /licenses/{license-key} | Deprecated. Remove a license identified by the license-key
[**DeleteLicenseKeyDelete**](SystemAdministrationSettingsApi.md#DeleteLicenseKeyDelete) | **Post** /licenses?action&#x3D;delete | Remove a license
[**DeleteNodeUserSshKeyRemoveSshKey**](SystemAdministrationSettingsApi.md#DeleteNodeUserSshKeyRemoveSshKey) | **Post** /node/users/{userid}/ssh-keys?action&#x3D;remove_ssh_key | Remove SSH public key from authorized_keys file for node user
[**DeletePrincipalIdentity**](SystemAdministrationSettingsApi.md#DeletePrincipalIdentity) | **Delete** /trust-management/principal-identities/{principal-identity-id} | Delete a principal identity
[**DeleteRegistrationToken**](SystemAdministrationSettingsApi.md#DeleteRegistrationToken) | **Delete** /aaa/registration-token/{token} | Delete registration access token
[**DeleteRoleBinding**](SystemAdministrationSettingsApi.md#DeleteRoleBinding) | **Delete** /aaa/role-bindings/{binding-id} | Delete user/group&#x27;s roles assignment
[**DeleteSupportBundlesAsyncResponseDeleteAsyncResponse**](SystemAdministrationSettingsApi.md#DeleteSupportBundlesAsyncResponseDeleteAsyncResponse) | **Post** /administration/support-bundles?action&#x3D;delete_async_response | Delete existing support bundles waiting to be downloaded
[**DeleteTokenBasedPrincipalIdentity**](SystemAdministrationSettingsApi.md#DeleteTokenBasedPrincipalIdentity) | **Delete** /trust-management/token-principal-identities/{principal-identity-id} | Delete a token-based principal identity
[**FetchIdentitySourceLdapServerCertificateFetchCertificate**](SystemAdministrationSettingsApi.md#FetchIdentitySourceLdapServerCertificateFetchCertificate) | **Post** /aaa/ldap-identity-sources?action&#x3D;fetch_certificate | Fetch the server certificate of an LDAP server
[**GenerateCsr**](SystemAdministrationSettingsApi.md#GenerateCsr) | **Post** /trust-management/csrs | Generate a New Certificate Signing Request
[**GetAllRoleBindings**](SystemAdministrationSettingsApi.md#GetAllRoleBindings) | **Get** /aaa/role-bindings | Get all users and groups with their roles
[**GetAllRolesInfo**](SystemAdministrationSettingsApi.md#GetAllRolesInfo) | **Get** /aaa/roles | Get information about all roles
[**GetCertificate**](SystemAdministrationSettingsApi.md#GetCertificate) | **Get** /trust-management/certificates/{cert-id} | Show Certificate Data for the Given Certificate ID
[**GetCertificates**](SystemAdministrationSettingsApi.md#GetCertificates) | **Get** /trust-management/certificates | Return All the User-Facing Components&#x27; Certificates
[**GetCrl**](SystemAdministrationSettingsApi.md#GetCrl) | **Get** /trust-management/crls/{crl-id} | Show CRL Data for the Given CRL ID
[**GetCrlDistributionPoint**](SystemAdministrationSettingsApi.md#GetCrlDistributionPoint) | **Get** /trust-management/crl-distribution-points/{crl-distribution-point-id} | Return the CrlDistributionPoint with &lt;crl-distribution-point-id&gt;
[**GetCrlDistributionPointPem**](SystemAdministrationSettingsApi.md#GetCrlDistributionPointPem) | **Post** /trust-management/crl-distribution-points/pem-file | Return stored CRL in PEM format
[**GetCrlDistributionPointStatus**](SystemAdministrationSettingsApi.md#GetCrlDistributionPointStatus) | **Get** /trust-management/crl-distribution-points/{crl-distribution-point-id}/status | Return the status of the CrlDistributionPoint
[**GetCrls**](SystemAdministrationSettingsApi.md#GetCrls) | **Get** /trust-management/crls | Return All Added CRLs
[**GetCsr**](SystemAdministrationSettingsApi.md#GetCsr) | **Get** /trust-management/csrs/{csr-id} | Show CSR Data for the Given CSR ID
[**GetCsrPem**](SystemAdministrationSettingsApi.md#GetCsrPem) | **Get** /trust-management/csrs/{csr-id}/pem-file | Get CSR PEM File for the Given CSR ID
[**GetCsrs**](SystemAdministrationSettingsApi.md#GetCsrs) | **Get** /trust-management/csrs | Return All the Generated CSRs
[**GetCurrentUserInfo**](SystemAdministrationSettingsApi.md#GetCurrentUserInfo) | **Get** /aaa/user-info | Get information about logged-in user. The permissions parameter of the NsxRole has been deprecated.
[**GetEULAAcceptance**](SystemAdministrationSettingsApi.md#GetEULAAcceptance) | **Get** /eula/acceptance | Return the acceptance status of end user license agreement 
[**GetEULAContent**](SystemAdministrationSettingsApi.md#GetEULAContent) | **Get** /eula/content | Return the content of end user license agreement 
[**GetGroupVidmSearchResult**](SystemAdministrationSettingsApi.md#GetGroupVidmSearchResult) | **Get** /aaa/vidm/groups | Get all the User Groups where vIDM display name matches the search key case insensitively. The search key is checked to be a substring of display name. This is a non paginated API.
[**GetLicense**](SystemAdministrationSettingsApi.md#GetLicense) | **Get** /license | Deprecated. Return the Enterprise License 
[**GetLicenseByKey**](SystemAdministrationSettingsApi.md#GetLicenseByKey) | **Get** /licenses/{license-key} | Deprecated. Get license properties for license identified by the license-key
[**GetLicenseUsageReport**](SystemAdministrationSettingsApi.md#GetLicenseUsageReport) | **Get** /licenses/licenses-usage | Get usage report of all registered modules
[**GetLicenseUsageReportInCsvFormatCsv**](SystemAdministrationSettingsApi.md#GetLicenseUsageReportInCsvFormatCsv) | **Get** /licenses/licenses-usage?format&#x3D;csv | Get usage report of all registred modules in CSV format
[**GetLicenses**](SystemAdministrationSettingsApi.md#GetLicenses) | **Get** /licenses | Get all licenses
[**GetOidcEndPoint**](SystemAdministrationSettingsApi.md#GetOidcEndPoint) | **Get** /trust-management/oidc-uris/{id} | Get an OpenID Connect end-point.
[**GetPrincipalIdentities**](SystemAdministrationSettingsApi.md#GetPrincipalIdentities) | **Get** /trust-management/principal-identities | Return the list of principal identities
[**GetPrincipalIdentity**](SystemAdministrationSettingsApi.md#GetPrincipalIdentity) | **Get** /trust-management/principal-identities/{principal-identity-id} | Get a principal identity
[**GetProxyConfig**](SystemAdministrationSettingsApi.md#GetProxyConfig) | **Get** /proxy/config | Returns the proxy configuration
[**GetRegistrationToken**](SystemAdministrationSettingsApi.md#GetRegistrationToken) | **Get** /aaa/registration-token/{token} | Get registration access token
[**GetRoleBinding**](SystemAdministrationSettingsApi.md#GetRoleBinding) | **Get** /aaa/role-bindings/{binding-id} | Get user/group&#x27;s role information
[**GetRoleInfo**](SystemAdministrationSettingsApi.md#GetRoleInfo) | **Get** /aaa/roles/{role} | Get role information
[**GetTelemetryAgreement**](SystemAdministrationSettingsApi.md#GetTelemetryAgreement) | **Get** /telemetry/agreement | Returns telemetry agreement information
[**GetTelemetryConfig**](SystemAdministrationSettingsApi.md#GetTelemetryConfig) | **Get** /telemetry/config | Returns the telemetry configuration
[**GetTokenBasedPrincipalIdentity**](SystemAdministrationSettingsApi.md#GetTokenBasedPrincipalIdentity) | **Get** /trust-management/token-principal-identities/{principal-identity-id} | Get a token-based principal identity
[**GetTrustObjects**](SystemAdministrationSettingsApi.md#GetTrustObjects) | **Get** /trust-management | Return the Properties of a Trust Manager
[**GetUserVidmSearchResult**](SystemAdministrationSettingsApi.md#GetUserVidmSearchResult) | **Get** /aaa/vidm/users | Get all the users from vIDM whose userName, givenName or familyName matches the search key case insensitively. The search key is checked to be a substring of name or given name or family name. This is a non paginated API.
[**GetVidmSearchResult**](SystemAdministrationSettingsApi.md#GetVidmSearchResult) | **Post** /aaa/vidm/search | Get all the users and groups from vIDM matching the search key case insensitively. The search key is checked to be a substring of name or given name or family name of user and display name of group. This is a non paginated API.
[**ImportCertificateImport**](SystemAdministrationSettingsApi.md#ImportCertificateImport) | **Post** /trust-management/csrs/{csr-id}?action&#x3D;import | Import a Certificate Associated with an Approved CSR
[**ListCrlDistributionPoints**](SystemAdministrationSettingsApi.md#ListCrlDistributionPoints) | **Get** /trust-management/crl-distribution-points | Return the list of CrlDistributionPoints
[**ListLdapIdentitySources**](SystemAdministrationSettingsApi.md#ListLdapIdentitySources) | **Get** /aaa/ldap-identity-sources | List LDAP identity sources
[**ListNodeUserSshKeys**](SystemAdministrationSettingsApi.md#ListNodeUserSshKeys) | **Get** /node/users/{userid}/ssh-keys | List SSH keys from authorized_keys file for node user
[**ListNodeUsers**](SystemAdministrationSettingsApi.md#ListNodeUsers) | **Get** /node/users | List node users
[**ListOidcEndPoints**](SystemAdministrationSettingsApi.md#ListOidcEndPoints) | **Get** /trust-management/oidc-uris | Return the list of OpenID Connect end-points.
[**ListRolesInfo**](SystemAdministrationSettingsApi.md#ListRolesInfo) | **Get** /aaa/roles-with-feature-permissions | Get information about all roles with features and their permissions
[**ListTokenBasedPrincipalIdentities**](SystemAdministrationSettingsApi.md#ListTokenBasedPrincipalIdentities) | **Get** /trust-management/token-principal-identities | Return the list of token-based principal identities. | These don&#x27;t have certificate or role information.
[**ProbeConfiguredLdapIdentitySourceProbe**](SystemAdministrationSettingsApi.md#ProbeConfiguredLdapIdentitySourceProbe) | **Post** /aaa/ldap-identity-sources/{ldap-identity-source-id}?action&#x3D;probe | Test the configuration of an existing LDAP identity source
[**ProbeIdentitySourceLdapServerProbeLdapServer**](SystemAdministrationSettingsApi.md#ProbeIdentitySourceLdapServerProbeLdapServer) | **Post** /aaa/ldap-identity-sources?action&#x3D;probe_ldap_server | Test an LDAP server
[**ProbeUnconfiguredLdapIdentitySourceProbeIdentitySource**](SystemAdministrationSettingsApi.md#ProbeUnconfiguredLdapIdentitySourceProbeIdentitySource) | **Post** /aaa/ldap-identity-sources?action&#x3D;probe_identity_source | Probe an LDAP identity source
[**ReadAuthProviderVidm**](SystemAdministrationSettingsApi.md#ReadAuthProviderVidm) | **Get** /node/aaa/providers/vidm | Read AAA provider vIDM properties
[**ReadAuthProviderVidmStatus**](SystemAdministrationSettingsApi.md#ReadAuthProviderVidmStatus) | **Get** /node/aaa/providers/vidm/status | Read AAA provider vIDM status
[**ReadAuthenticationPolicyProperties**](SystemAdministrationSettingsApi.md#ReadAuthenticationPolicyProperties) | **Get** /node/aaa/auth-policy | Read node authentication policy configuration
[**ReadLdapIdentitySource**](SystemAdministrationSettingsApi.md#ReadLdapIdentitySource) | **Get** /aaa/ldap-identity-sources/{ldap-identity-source-id} | Read a single LDAP identity source
[**ReadNodeUser**](SystemAdministrationSettingsApi.md#ReadNodeUser) | **Get** /node/users/{userid} | Read node user
[**RegisterPrincipalIdentity**](SystemAdministrationSettingsApi.md#RegisterPrincipalIdentity) | **Post** /trust-management/principal-identities | Register a name-certificate combination.
[**RegisterPrincipalIdentityWithCertificate**](SystemAdministrationSettingsApi.md#RegisterPrincipalIdentityWithCertificate) | **Post** /trust-management/principal-identities/with-certificate | Register a name-certificate combination.
[**RegisterTokenBasedPrincipalIdentity**](SystemAdministrationSettingsApi.md#RegisterTokenBasedPrincipalIdentity) | **Post** /trust-management/token-principal-identities | Register a token-based principal identity.
[**SearchLdapIdentitySource**](SystemAdministrationSettingsApi.md#SearchLdapIdentitySource) | **Post** /aaa/ldap-identity-sources/{ldap-identity-source-id}/search | Search the LDAP identity source
[**SelfSignCertificateSelfSign**](SystemAdministrationSettingsApi.md#SelfSignCertificateSelfSign) | **Post** /trust-management/csrs/{csr-id}?action&#x3D;self_sign | Self-Sign the CSR
[**UpdateAuthProviderVidm**](SystemAdministrationSettingsApi.md#UpdateAuthProviderVidm) | **Put** /node/aaa/providers/vidm | Update AAA provider vIDM properties
[**UpdateAuthenticationPolicyProperties**](SystemAdministrationSettingsApi.md#UpdateAuthenticationPolicyProperties) | **Put** /node/aaa/auth-policy | Update node authentication policy configuration
[**UpdateCrl**](SystemAdministrationSettingsApi.md#UpdateCrl) | **Put** /trust-management/crls/{crl-id} | Update CRL for the Given CRL ID
[**UpdateCrlDistributionPoint**](SystemAdministrationSettingsApi.md#UpdateCrlDistributionPoint) | **Put** /trust-management/crl-distribution-points/{crl-distribution-point-id} | Update CrlDistributionPoint with &lt;crl-distribution-point-id&gt; This allows updating the ManagedResource fields. 
[**UpdateLicense**](SystemAdministrationSettingsApi.md#UpdateLicense) | **Put** /license | Deprecated. Assign an Updated Enterprise License Key 
[**UpdateNodeUser**](SystemAdministrationSettingsApi.md#UpdateNodeUser) | **Put** /node/users/{userid} | Update node user
[**UpdateOidcEndPointThumbprintUpdateThumbprint**](SystemAdministrationSettingsApi.md#UpdateOidcEndPointThumbprintUpdateThumbprint) | **Post** /trust-management/oidc-uris?action&#x3D;update_thumbprint | Update a OpenID Connect end-point&#x27;s thumbprint
[**UpdatePrincipalIdentityCertificateUpdateCertificate**](SystemAdministrationSettingsApi.md#UpdatePrincipalIdentityCertificateUpdateCertificate) | **Post** /trust-management/principal-identities?action&#x3D;update_certificate | Update a principal identity&#x27;s certificate
[**UpdateProxyConfig**](SystemAdministrationSettingsApi.md#UpdateProxyConfig) | **Put** /proxy/config | Creates or updates the proxy configuration
[**UpdateRoleBinding**](SystemAdministrationSettingsApi.md#UpdateRoleBinding) | **Put** /aaa/role-bindings/{binding-id} | Update User or Group&#x27;s roles
[**UpdateTelemetryAgreement**](SystemAdministrationSettingsApi.md#UpdateTelemetryAgreement) | **Put** /telemetry/agreement | Set telemetry agreement information
[**UpdateTelemetryConfig**](SystemAdministrationSettingsApi.md#UpdateTelemetryConfig) | **Put** /telemetry/config | Creates or updates the telemetry configuration

# **AcceptEULA**
> AcceptEULA(ctx, )
Accept end user license agreement 

Accept end user license agreement 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCertificateImport**
> CertificateList AddCertificateImport(ctx, body)
Add a New Certificate

Adds a new private-public certificate or a chain of certificates (CAs) and, optionally, a private key that can be applied to one of the user-facing components (appliance management or edge). The certificate and the key should be stored in PEM format. If no private key is provided, the certificate is used as a client certificate in the trust store. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TrustObjectData**](TrustObjectData.md)|  | 

### Return type

[**CertificateList**](CertificateList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCrlImport**
> CrlList AddCrlImport(ctx, body)
Add a New Certificate Revocation List

Adds a new certificate revocation list (CRL). The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CrlObjectData**](CrlObjectData.md)|  | 

### Return type

[**CrlList**](CrlList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNodeUserSshKeyAddSshKey**
> AddNodeUserSshKeyAddSshKey(ctx, body, userid)
Add SSH public key to authorized_keys file for node user

Add SSH public key to authorized_keys file for node user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SshKeyProperties**](SshKeyProperties.md)|  | 
  **userid** | **string**| User id of the user | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddOidcEndPoint**
> OidcEndPoint AddOidcEndPoint(ctx, body)
Add an OpenID Connect end-point.

This request also fetches the issuer and jwks_uri meta-data from the OIDC end-point and stores it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OidcEndPoint**](OidcEndPoint.md)|  | 

### Return type

[**OidcEndPoint**](OidcEndPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectSupportBundlesCollect**
> SupportBundleResult CollectSupportBundlesCollect(ctx, body, optional)
Collect support bundles from registered cluster and fabric nodes

Collect support bundles from registered cluster and fabric nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupportBundleRequest**](SupportBundleRequest.md)|  | 
 **optional** | ***SystemAdministrationSettingsApiCollectSupportBundlesCollectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiCollectSupportBundlesCollectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideAsyncResponse** | **optional.**| Override any existing support bundle async response | [default to false]
 **requireDeleteOrOverrideAsyncResponse** | **optional.**| Suppress auto-deletion of generated support bundle | [default to false]

### Return type

[**SupportBundleResult**](SupportBundleResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCrlDistributionPoint**
> CrlDistributionPoint CreateCrlDistributionPoint(ctx, body)
Create a Crl Distribution Point

Create an entity that will represent a Crl Distribution Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CrlDistributionPoint**](CrlDistributionPoint.md)|  | 

### Return type

[**CrlDistributionPoint**](CrlDistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLicense**
> License CreateLicense(ctx, body)
Add a new license key

This will add a license key to the system. The API supports adding only one license key for each license edition type - Standard, Advanced or Enterprise. If a new license key is tried to add for an edition for which the license key already exists, then this API will return an error. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**License**](License.md)|  | 

### Return type

[**License**](License.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateLdapIdentitySource**
> LdapIdentitySource CreateOrUpdateLdapIdentitySource(ctx, body, ldapIdentitySourceId)
Update an existing LDAP identity source

Update the configuration of an existing LDAP identity source. You may wish to verify the new configuration using the POST /aaa/ldap-identity-sources?action=probe API before changing the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LdapIdentitySource**](LdapIdentitySource.md)|  | 
  **ldapIdentitySourceId** | **string**|  | 

### Return type

[**LdapIdentitySource**](LdapIdentitySource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRegistrationToken**
> RegistrationToken CreateRegistrationToken(ctx, )
Create registration access token

The privileges of the registration token will be the same as the caller.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegistrationToken**](RegistrationToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRoleBinding**
> RoleBinding CreateRoleBinding(ctx, body)
Assign roles to User or Group

When assigning a user role, specify the user name with the same case as it appears in vIDM to access the NSX-T user interface. For example, if vIDM has the user name User1@example.com then the name attribute in the API call must be be User1@example.com and cannot be user1@example.com. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleBinding**](RoleBinding.md)|  | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllStaleRoleBindingsDeleteStaleBindings**
> DeleteAllStaleRoleBindingsDeleteStaleBindings(ctx, )
Delete all stale role assignments

Delete all stale role assignments

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCertificate**
> DeleteCertificate(ctx, certId)
Delete Certificate for the Given Certificate ID

Removes the specified certificate. The private key associated with the certificate is also deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certId** | **string**| ID of certificate to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCrl**
> DeleteCrl(ctx, crlId)
Delete a CRL

Deletes an existing CRL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlId** | **string**| ID of CRL to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCrlDistributionPoint**
> DeleteCrlDistributionPoint(ctx, crlDistributionPointId)
Delete a CrlDistributionPoint

Delete a CrlDistributionPoint. It does not delete the actual CRL. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlDistributionPointId** | **string**| Unique id of the CrlDistributionPoint to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCsr**
> DeleteCsr(ctx, csrId)
Delete a CSR

Removes a specified CSR. If a CSR is not used for verification, you can delete it. Note that the CSR import and upload POST actions automatically delete the associated CSR. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **csrId** | **string**| ID of CSR to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLdapIdentitySource**
> DeleteLdapIdentitySource(ctx, ldapIdentitySourceId)
Delete an LDAP identity source

Delete an LDAP identity source. Users defined in that source will no longer be able to access NSX.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapIdentitySourceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLicense**
> DeleteLicense(ctx, licenseKey)
Deprecated. Remove a license identified by the license-key

Deprecated. Use POST /licenses?action=delete API instead. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **licenseKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLicenseKeyDelete**
> DeleteLicenseKeyDelete(ctx, body)
Remove a license

This will delete the license key identified in the request body by \"license_key\" and its properties from the system. Attempting to delete the last license key will result in an error. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**License**](License.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNodeUserSshKeyRemoveSshKey**
> DeleteNodeUserSshKeyRemoveSshKey(ctx, body, userid)
Remove SSH public key from authorized_keys file for node user

Remove SSH public key from authorized_keys file for node user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SshKeyBaseProperties**](SshKeyBaseProperties.md)|  | 
  **userid** | **string**| User id of the user | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePrincipalIdentity**
> DeletePrincipalIdentity(ctx, principalIdentityId)
Delete a principal identity

Delete a principal identity. It does not delete the certificate. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **principalIdentityId** | **string**| Unique id of the principal identity to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRegistrationToken**
> DeleteRegistrationToken(ctx, token)
Delete registration access token

Delete registration access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Registration token | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleBinding**
> DeleteRoleBinding(ctx, bindingId)
Delete user/group's roles assignment

Delete user/group's roles assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**| User/Group&#x27;s id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSupportBundlesAsyncResponseDeleteAsyncResponse**
> DeleteSupportBundlesAsyncResponseDeleteAsyncResponse(ctx, )
Delete existing support bundles waiting to be downloaded

Delete existing support bundles waiting to be downloaded.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTokenBasedPrincipalIdentity**
> DeleteTokenBasedPrincipalIdentity(ctx, principalIdentityId)
Delete a token-based principal identity

Delete a token-based principal identity. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **principalIdentityId** | **string**| Unique id of the token-based principal identity to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchIdentitySourceLdapServerCertificateFetchCertificate**
> PeerCertificateChain FetchIdentitySourceLdapServerCertificateFetchCertificate(ctx, body)
Fetch the server certificate of an LDAP server

Attempt to connect to an LDAP server and retrieve the server certificate it presents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdentitySourceLdapServerEndpoint**](IdentitySourceLdapServerEndpoint.md)|  | 

### Return type

[**PeerCertificateChain**](PeerCertificateChain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateCsr**
> Csr GenerateCsr(ctx, body)
Generate a New Certificate Signing Request

Creates a new certificate signing request (CSR). A CSR is encrypted text that contains information about your organization (organization name, country, and so on) and your Web server's public key, which is a public certificate the is generated on the server that can be used to forward this request to a certificate authority (CA). A private key is also usually created at the same time as the CSR. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Csr**](Csr.md)|  | 

### Return type

[**Csr**](Csr.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRoleBindings**
> RoleBindingListResult GetAllRoleBindings(ctx, optional)
Get all users and groups with their roles

Get all users and groups with their roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationSettingsApiGetAllRoleBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetAllRoleBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **identitySourceId** | **optional.String**| Identity source ID | 
 **identitySourceType** | **optional.String**| Identity source type | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **name** | **optional.String**| User/Group name | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **role** | **optional.String**| Role ID | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Type | 

### Return type

[**RoleBindingListResult**](RoleBindingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRolesInfo**
> RoleListResult GetAllRolesInfo(ctx, )
Get information about all roles

Get information about all roles

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RoleListResult**](RoleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificate**
> Certificate GetCertificate(ctx, certId, optional)
Show Certificate Data for the Given Certificate ID

Returns information for the specified certificate ID, including the certificate's UUID; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certId** | **string**| ID of certificate to read | 
 **optional** | ***SystemAdministrationSettingsApiGetCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetCertificateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]

### Return type

[**Certificate**](Certificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificates**
> CertificateList GetCertificates(ctx, optional)
Return All the User-Facing Components' Certificates

Returns all certificate information viewable by the user, including each certificate's UUID; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationSettingsApiGetCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetCertificatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Type of certificate to return | 

### Return type

[**CertificateList**](CertificateList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrl**
> Crl GetCrl(ctx, crlId, optional)
Show CRL Data for the Given CRL ID

Returns information about the specified CRL. For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlId** | **string**| ID of CRL to read | 
 **optional** | ***SystemAdministrationSettingsApiGetCrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetCrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]

### Return type

[**Crl**](Crl.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrlDistributionPoint**
> CrlDistributionPoint GetCrlDistributionPoint(ctx, crlDistributionPointId)
Return the CrlDistributionPoint with <crl-distribution-point-id>

Return the CrlDistributionPoint with <crl-distribution-point-id>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlDistributionPointId** | **string**|  | 

### Return type

[**CrlDistributionPoint**](CrlDistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrlDistributionPointPem**
> string GetCrlDistributionPointPem(ctx, body)
Return stored CRL in PEM format

Return stored CRL in PEM format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CrlPemRequestType**](CrlPemRequestType.md)|  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrlDistributionPointStatus**
> CrlDistributionPointStatus GetCrlDistributionPointStatus(ctx, crlDistributionPointId)
Return the status of the CrlDistributionPoint

Return the status of the CrlDistributionPoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlDistributionPointId** | **string**|  | 

### Return type

[**CrlDistributionPointStatus**](CrlDistributionPointStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrls**
> CrlList GetCrls(ctx, optional)
Return All Added CRLs

Returns information about all CRLs. For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationSettingsApiGetCrlsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetCrlsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Type of certificate to return | 

### Return type

[**CrlList**](CrlList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCsr**
> Csr GetCsr(ctx, csrId)
Show CSR Data for the Given CSR ID

Returns information about the specified CSR.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **csrId** | **string**| ID of CSR to read | 

### Return type

[**Csr**](Csr.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCsrPem**
> string GetCsrPem(ctx, csrId)
Get CSR PEM File for the Given CSR ID

Downloads the CSR PEM file for a specified CSR. Clients must include an Accept: text/plain request header.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **csrId** | **string**| ID of CSR to read | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCsrs**
> CsrList GetCsrs(ctx, optional)
Return All the Generated CSRs

Returns information about all of the CSRs that have been created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationSettingsApiGetCsrsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetCsrsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**CsrList**](CsrList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUserInfo**
> UserInfo GetCurrentUserInfo(ctx, )
Get information about logged-in user. The permissions parameter of the NsxRole has been deprecated.

Get information about logged-in user. The permissions parameter of the NsxRole has been deprecated.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEULAAcceptance**
> EulaAcceptance GetEULAAcceptance(ctx, )
Return the acceptance status of end user license agreement 

Return the acceptance status of end user license agreement 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EulaAcceptance**](EULAAcceptance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEULAContent**
> EulaContent GetEULAContent(ctx, optional)
Return the content of end user license agreement 

Return the content of end user license agreement in the specified format. By default, it's pure string without line break 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationSettingsApiGetEULAContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetEULAContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **valueFormat** | **optional.String**| End User License Agreement content output format | 

### Return type

[**EulaContent**](EULAContent.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupVidmSearchResult**
> VidmInfoListResult GetGroupVidmSearchResult(ctx, searchString, optional)
Get all the User Groups where vIDM display name matches the search key case insensitively. The search key is checked to be a substring of display name. This is a non paginated API.

Get all the User Groups where vIDM display name matches the search key case insensitively. The search key is checked to be a substring of display name. This is a non paginated API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchString** | **string**| Search string to search for.  | 
 **optional** | ***SystemAdministrationSettingsApiGetGroupVidmSearchResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetGroupVidmSearchResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VidmInfoListResult**](VidmInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicense**
> License GetLicense(ctx, )
Deprecated. Return the Enterprise License 

Deprecated. Use the GET /licenses API instead. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**License**](License.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseByKey**
> License GetLicenseByKey(ctx, licenseKey)
Deprecated. Get license properties for license identified by the license-key

Deprecated. Use GET /licenses API instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **licenseKey** | **string**|  | 

### Return type

[**License**](License.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseUsageReport**
> FeatureUsageList GetLicenseUsageReport(ctx, )
Get usage report of all registered modules

Returns usage report of all registered modules 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FeatureUsageList**](FeatureUsageList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseUsageReportInCsvFormatCsv**
> FeatureUsageListInCsvFormat GetLicenseUsageReportInCsvFormatCsv(ctx, )
Get usage report of all registred modules in CSV format

Returns usage report of all registered modules in CSV format 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FeatureUsageListInCsvFormat**](FeatureUsageListInCsvFormat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenses**
> LicensesListResult GetLicenses(ctx, )
Get all licenses

Returns all licenses. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LicensesListResult**](LicensesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOidcEndPoint**
> OidcEndPoint GetOidcEndPoint(ctx, id, optional)
Get an OpenID Connect end-point.

When ?refresh=true is added to the request, the meta-data is newly fetched from the OIDC end-point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemAdministrationSettingsApiGetOidcEndPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetOidcEndPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **optional.Bool**| Refresh meta-data | [default to false]

### Return type

[**OidcEndPoint**](OidcEndPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrincipalIdentities**
> PrincipalIdentityList GetPrincipalIdentities(ctx, )
Return the list of principal identities

Returns the list of principals registered with a certificate.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PrincipalIdentityList**](PrincipalIdentityList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrincipalIdentity**
> PrincipalIdentity GetPrincipalIdentity(ctx, principalIdentityId)
Get a principal identity

Get a stored principal identity 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **principalIdentityId** | **string**| ID of the principal identity to get | 

### Return type

[**PrincipalIdentity**](PrincipalIdentity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxyConfig**
> Proxy GetProxyConfig(ctx, )
Returns the proxy configuration

Returns the proxy configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Proxy**](Proxy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegistrationToken**
> RegistrationToken GetRegistrationToken(ctx, token)
Get registration access token

Get registration access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Registration token | 

### Return type

[**RegistrationToken**](RegistrationToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleBinding**
> RoleBinding GetRoleBinding(ctx, bindingId)
Get user/group's role information

Get user/group's role information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**| User/Group&#x27;s id | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleInfo**
> RoleWithFeatures GetRoleInfo(ctx, role)
Get role information

Get role information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | **string**| Role id | 

### Return type

[**RoleWithFeatures**](RoleWithFeatures.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTelemetryAgreement**
> TelemetryAgreement GetTelemetryAgreement(ctx, )
Returns telemetry agreement information

Returns telemetry agreement information.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TelemetryAgreement**](TelemetryAgreement.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTelemetryConfig**
> TelemetryConfig GetTelemetryConfig(ctx, )
Returns the telemetry configuration

Returns the telemetry configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TelemetryConfig**](TelemetryConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenBasedPrincipalIdentity**
> TokenBasedPrincipalIdentity GetTokenBasedPrincipalIdentity(ctx, principalIdentityId)
Get a token-based principal identity

Get a stored token-based principal identity 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **principalIdentityId** | **string**| ID of the principal identity to get | 

### Return type

[**TokenBasedPrincipalIdentity**](TokenBasedPrincipalIdentity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrustObjects**
> TrustManagementData GetTrustObjects(ctx, )
Return the Properties of a Trust Manager

Returns information about the supported algorithms and key sizes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TrustManagementData**](TrustManagementData.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserVidmSearchResult**
> VidmInfoListResult GetUserVidmSearchResult(ctx, searchString, optional)
Get all the users from vIDM whose userName, givenName or familyName matches the search key case insensitively. The search key is checked to be a substring of name or given name or family name. This is a non paginated API.

Get all the users from vIDM whose userName, givenName or familyName matches the search key case insensitively. The search key is checked to be a substring of name or given name or family name. This is a non paginated API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchString** | **string**| Search string to search for.  | 
 **optional** | ***SystemAdministrationSettingsApiGetUserVidmSearchResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetUserVidmSearchResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VidmInfoListResult**](VidmInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVidmSearchResult**
> VidmInfoListResult GetVidmSearchResult(ctx, searchString, optional)
Get all the users and groups from vIDM matching the search key case insensitively. The search key is checked to be a substring of name or given name or family name of user and display name of group. This is a non paginated API.

Get all the users and groups from vIDM matching the search key case insensitively. The search key is checked to be a substring of name or given name or family name of user and display name of group. This is a non paginated API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchString** | **string**| Search string to search for.  | 
 **optional** | ***SystemAdministrationSettingsApiGetVidmSearchResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiGetVidmSearchResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VidmInfoListResult**](VidmInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportCertificateImport**
> CertificateList ImportCertificateImport(ctx, body, csrId)
Import a Certificate Associated with an Approved CSR

Imports a certificate authority (CA)-signed certificate for a CSR. This action links the certificate to the private key created by the CSR. The pem_encoded string in the request body is the signed certificate provided by your CA in response to the CSR that you provide to them. The import POST action automatically deletes the associated CSR. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TrustObjectData**](TrustObjectData.md)|  | 
  **csrId** | **string**| CSR this certificate is associated with | 

### Return type

[**CertificateList**](CertificateList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCrlDistributionPoints**
> CrlDistributionPointList ListCrlDistributionPoints(ctx, optional)
Return the list of CrlDistributionPoints

Return the list of CrlDistributionPoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationSettingsApiListCrlDistributionPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiListCrlDistributionPointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**CrlDistributionPointList**](CrlDistributionPointList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLdapIdentitySources**
> LdapIdentitySourceListResult ListLdapIdentitySources(ctx, optional)
List LDAP identity sources

Return a list of all configured LDAP identity sources.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationSettingsApiListLdapIdentitySourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiListLdapIdentitySourcesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LdapIdentitySourceListResult**](LdapIdentitySourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeUserSshKeys**
> SshKeyPropertiesListResult ListNodeUserSshKeys(ctx, userid)
List SSH keys from authorized_keys file for node user

Returns a list of all SSH keys from authorized_keys file for node user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**| User id of the user | 

### Return type

[**SshKeyPropertiesListResult**](SshKeyPropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeUsers**
> NodeUserPropertiesListResult ListNodeUsers(ctx, )
List node users

Returns the list of users configued to log in to the NSX appliance. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeUserPropertiesListResult**](NodeUserPropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOidcEndPoints**
> OidcEndPointListResult ListOidcEndPoints(ctx, )
Return the list of OpenID Connect end-points.

Return the list of OpenID Connect end-points.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OidcEndPointListResult**](OidcEndPointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRolesInfo**
> RoleWithFeaturesListResult ListRolesInfo(ctx, optional)
Get information about all roles with features and their permissions

Get information about all roles with features and their permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationSettingsApiListRolesInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiListRolesInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RoleWithFeaturesListResult**](RoleWithFeaturesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTokenBasedPrincipalIdentities**
> TokenBasedPrincipalIdentityListResult ListTokenBasedPrincipalIdentities(ctx, )
Return the list of token-based principal identities. | These don't have certificate or role information.

Return the list of token-based principal identities. | These don't have certificate or role information.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TokenBasedPrincipalIdentityListResult**](TokenBasedPrincipalIdentityListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProbeConfiguredLdapIdentitySourceProbe**
> LdapIdentitySourceProbeResults ProbeConfiguredLdapIdentitySourceProbe(ctx, ldapIdentitySourceId)
Test the configuration of an existing LDAP identity source

Attempt to connect to an existing LDAP identity source and report any errors encountered.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapIdentitySourceId** | **string**|  | 

### Return type

[**LdapIdentitySourceProbeResults**](LdapIdentitySourceProbeResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProbeIdentitySourceLdapServerProbeLdapServer**
> IdentitySourceLdapServerProbeResult ProbeIdentitySourceLdapServerProbeLdapServer(ctx, body)
Test an LDAP server

Attempt to connect to an LDAP server and ensure that the server can be contacted using the given URL and authentication credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdentitySourceLdapServer**](IdentitySourceLdapServer.md)|  | 

### Return type

[**IdentitySourceLdapServerProbeResult**](IdentitySourceLdapServerProbeResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProbeUnconfiguredLdapIdentitySourceProbeIdentitySource**
> LdapIdentitySourceProbeResults ProbeUnconfiguredLdapIdentitySourceProbeIdentitySource(ctx, body)
Probe an LDAP identity source

Verify that the configuration of an LDAP identity source is correct before actually creating the source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LdapIdentitySource**](LdapIdentitySource.md)|  | 

### Return type

[**LdapIdentitySourceProbeResults**](LdapIdentitySourceProbeResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAuthProviderVidm**
> NodeAuthProviderVidmProperties ReadAuthProviderVidm(ctx, )
Read AAA provider vIDM properties

Read AAA provider vIDM properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeAuthProviderVidmProperties**](NodeAuthProviderVidmProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAuthProviderVidmStatus**
> NodeAuthProviderVidmStatus ReadAuthProviderVidmStatus(ctx, )
Read AAA provider vIDM status

Read AAA provider vIDM status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeAuthProviderVidmStatus**](NodeAuthProviderVidmStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAuthenticationPolicyProperties**
> AuthenticationPolicyProperties ReadAuthenticationPolicyProperties(ctx, )
Read node authentication policy configuration

Returns information about the currently configured authentication policies on the node. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthenticationPolicyProperties**](AuthenticationPolicyProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLdapIdentitySource**
> LdapIdentitySource ReadLdapIdentitySource(ctx, ldapIdentitySourceId)
Read a single LDAP identity source

Return details about one LDAP identity source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapIdentitySourceId** | **string**|  | 

### Return type

[**LdapIdentitySource**](LdapIdentitySource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeUser**
> NodeUserProperties ReadNodeUser(ctx, userid)
Read node user

Returns information about a specified user who is configued to log in to the NSX appliance 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**| User id of the user | 

### Return type

[**NodeUserProperties**](NodeUserProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterPrincipalIdentity**
> PrincipalIdentity RegisterPrincipalIdentity(ctx, body)
Register a name-certificate combination.

Associates a principal's name with a certificate that is used to authenticate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities. Deprecated, use POST /trust-management/principal-identities/with-certificate instead. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PrincipalIdentity**](PrincipalIdentity.md)|  | 

### Return type

[**PrincipalIdentity**](PrincipalIdentity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterPrincipalIdentityWithCertificate**
> PrincipalIdentity RegisterPrincipalIdentityWithCertificate(ctx, body)
Register a name-certificate combination.

Create a principal identity with a new, unused, certificate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PrincipalIdentityWithCertificate**](PrincipalIdentityWithCertificate.md)|  | 

### Return type

[**PrincipalIdentity**](PrincipalIdentity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterTokenBasedPrincipalIdentity**
> TokenBasedPrincipalIdentity RegisterTokenBasedPrincipalIdentity(ctx, body)
Register a token-based principal identity.

Register a principal identity that is going to be authenticated through a token. The combination name and node_id needs to be unique across token-based and certificate-based principal identities. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TokenBasedPrincipalIdentity**](TokenBasedPrincipalIdentity.md)|  | 

### Return type

[**TokenBasedPrincipalIdentity**](TokenBasedPrincipalIdentity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLdapIdentitySource**
> LdapIdentitySourceSearchResultList SearchLdapIdentitySource(ctx, ldapIdentitySourceId, optional)
Search the LDAP identity source

Search the LDAP identity source for users and groups that match the given filter_value. In most cases, the LDAP source performs a case-insensitive search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapIdentitySourceId** | **string**|  | 
 **optional** | ***SystemAdministrationSettingsApiSearchLdapIdentitySourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationSettingsApiSearchLdapIdentitySourceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterValue** | **optional.String**| Search filter value | 

### Return type

[**LdapIdentitySourceSearchResultList**](LdapIdentitySourceSearchResultList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SelfSignCertificateSelfSign**
> Certificate SelfSignCertificateSelfSign(ctx, csrId, daysValid)
Self-Sign the CSR

Self-signs the previously generated CSR. This action is similar to the import certificate action, but instead of using a public certificate signed by a CA, the self_sign POST action uses a certificate that is signed with NSX's own private key. For validity, if a value greater than 825 days is provided, it will be set to 825 days. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **csrId** | **string**| CSR this certificate is associated with | 
  **daysValid** | **int64**| Number of days the certificate will be valid, default 825 days | [default to 825]

### Return type

[**Certificate**](Certificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthProviderVidm**
> NodeAuthProviderVidmProperties UpdateAuthProviderVidm(ctx, body)
Update AAA provider vIDM properties

Update AAA provider vIDM properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeAuthProviderVidmProperties**](NodeAuthProviderVidmProperties.md)|  | 

### Return type

[**NodeAuthProviderVidmProperties**](NodeAuthProviderVidmProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthenticationPolicyProperties**
> AuthenticationPolicyProperties UpdateAuthenticationPolicyProperties(ctx, body)
Update node authentication policy configuration

Update the currently configured authentication policy on the node. If any of api_max_auth_failures, api_failed_auth_reset_period, or api_failed_auth_lockout_period are modified, the http service is automatically restarted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthenticationPolicyProperties**](AuthenticationPolicyProperties.md)|  | 

### Return type

[**AuthenticationPolicyProperties**](AuthenticationPolicyProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCrl**
> Crl UpdateCrl(ctx, body, crlId)
Update CRL for the Given CRL ID

Updates an existing CRL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Crl**](Crl.md)|  | 
  **crlId** | **string**| ID of CRL to update | 

### Return type

[**Crl**](Crl.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCrlDistributionPoint**
> CrlDistributionPoint UpdateCrlDistributionPoint(ctx, body, crlDistributionPointId)
Update CrlDistributionPoint with <crl-distribution-point-id> This allows updating the ManagedResource fields. 

Update CrlDistributionPoint with <crl-distribution-point-id> This allows updating the ManagedResource fields. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CrlDistributionPoint**](CrlDistributionPoint.md)|  | 
  **crlDistributionPointId** | **string**|  | 

### Return type

[**CrlDistributionPoint**](CrlDistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLicense**
> License UpdateLicense(ctx, body)
Deprecated. Assign an Updated Enterprise License Key 

Deprecated. Use the POST /licenses API instead 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**License**](License.md)|  | 

### Return type

[**License**](License.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeUser**
> NodeUserProperties UpdateNodeUser(ctx, body, userid)
Update node user

Updates attributes of an existing NSX appliance user. This method cannot be used to add a new user. Modifiable attributes include the username, full name of the user, and password. If you specify a password in a PUT request, it is not returned in the response. Nor is it returned in a GET request. The specified password does not meet the following complexity requirements: - minimum 12 characters in length - minimum 1 uppercase character - minimum 1 lowercase character - minimum 1 numeric character - minimum 1 special character - minimum 5 unique characters - default password complexity rules as enforced by the Linux PAM module 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeUserProperties**](NodeUserProperties.md)|  | 
  **userid** | **string**| User id of the user | 

### Return type

[**NodeUserProperties**](NodeUserProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOidcEndPointThumbprintUpdateThumbprint**
> OidcEndPoint UpdateOidcEndPointThumbprintUpdateThumbprint(ctx, body)
Update a OpenID Connect end-point's thumbprint

Update a OpenID Connect end-point's thumbprint used to connect to the oidc_uri through SSL 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOidcEndPointThumbprintRequest**](UpdateOidcEndPointThumbprintRequest.md)|  | 

### Return type

[**OidcEndPoint**](OidcEndPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrincipalIdentityCertificateUpdateCertificate**
> PrincipalIdentity UpdatePrincipalIdentityCertificateUpdateCertificate(ctx, body)
Update a principal identity's certificate

Update a principal identity's certificate 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePrincipalIdentityCertificateRequest**](UpdatePrincipalIdentityCertificateRequest.md)|  | 

### Return type

[**PrincipalIdentity**](PrincipalIdentity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProxyConfig**
> Proxy UpdateProxyConfig(ctx, body)
Creates or updates the proxy configuration

Updates or creates the proxy configuration, and returns the new configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Proxy**](Proxy.md)|  | 

### Return type

[**Proxy**](Proxy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleBinding**
> RoleBinding UpdateRoleBinding(ctx, body, bindingId)
Update User or Group's roles

Update User or Group's roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleBinding**](RoleBinding.md)|  | 
  **bindingId** | **string**| User/Group&#x27;s id | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTelemetryAgreement**
> TelemetryAgreement UpdateTelemetryAgreement(ctx, body)
Set telemetry agreement information

Set telemetry agreement information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TelemetryAgreement**](TelemetryAgreement.md)|  | 

### Return type

[**TelemetryAgreement**](TelemetryAgreement.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTelemetryConfig**
> TelemetryConfig UpdateTelemetryConfig(ctx, body)
Creates or updates the telemetry configuration

Updates or creates the telemetry configuration, and returns the new configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TelemetryConfig**](TelemetryConfig.md)|  | 

### Return type

[**TelemetryConfig**](TelemetryConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

