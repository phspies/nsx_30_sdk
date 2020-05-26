# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetInterSiteAphCertificateSetApplianceProxyCertificateForInterSiteCommunication**](DefaultApi.md#SetInterSiteAphCertificateSetApplianceProxyCertificateForInterSiteCommunication) | **Post** /trust-management/certificates?action&#x3D;set_appliance_proxy_certificate_for_inter_site_communication | Set a certificate as the Appliance Proxy certificate to be used in inter-site communication
[**SetPrincipalIdentityCertificateForFederationSetPiCertificateForFederation**](DefaultApi.md#SetPrincipalIdentityCertificateForFederationSetPiCertificateForFederation) | **Post** /trust-management/certificates?action&#x3D;set_pi_certificate_for_federation | Set a certificate as a GM or LM Principal Identity certificate
[**ValidateCertificateValidate**](DefaultApi.md#ValidateCertificateValidate) | **Get** /trust-management/certificates/{cert-id}?action&#x3D;validate | Validate a certificate

# **SetInterSiteAphCertificateSetApplianceProxyCertificateForInterSiteCommunication**
> SetInterSiteAphCertificateSetApplianceProxyCertificateForInterSiteCommunication(ctx, body)
Set a certificate as the Appliance Proxy certificate to be used in inter-site communication

Set a certificate that has been imported to be the Appliance Proxy certificate used for communicating with Appliance Proxies on other sites. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetInterSiteAphCertificateRequest**](SetInterSiteAphCertificateRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPrincipalIdentityCertificateForFederationSetPiCertificateForFederation**
> SetPrincipalIdentityCertificateForFederationSetPiCertificateForFederation(ctx, body)
Set a certificate as a GM or LM Principal Identity certificate

Set a certificate that has been imported to be either the principal identity certificate for the local cluster with either GM or LM service type. Currently, the service type specified must match the current service type of the local cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetPrincipalIdentityCertificateForFederationRequest**](SetPrincipalIdentityCertificateForFederationRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateCertificateValidate**
> CertificateCheckingStatus ValidateCertificateValidate(ctx, certId)
Validate a certificate

Checks whether certificate is valid. When the certificate contains a chain, the full chain is validated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certId** | **string**| ID of certificate to validate | 

### Return type

[**CertificateCheckingStatus**](CertificateCheckingStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

