# \DomainApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentAddDomain**](DomainApi.md#EnvironmentAddDomain) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/domain/{hostName} | 
[**EnvironmentRemoveDomain**](DomainApi.md#EnvironmentRemoveDomain) | **Delete** /account/{accountId}/application/{applicationId}/environment/{environmentName}/domain/{hostName} | 
[**GetConfiguration**](DomainApi.md#GetConfiguration) | **Get** /account/{accountId}/domain/{hostName}/https | 
[**PostConfiguration**](DomainApi.md#PostConfiguration) | **Post** /account/{accountId}/domain/{hostName}/https | 
[**RenewCertificate**](DomainApi.md#RenewCertificate) | **Post** /account/{accountId}/domain/{hostName}/renewCertificate | 
[**VerifyEngaged**](DomainApi.md#VerifyEngaged) | **Get** /account/{accountId}/domain/{hostName}/verifyEngaged | 


# **EnvironmentAddDomain**
> EnvironmentSummary EnvironmentAddDomain(ctx, accountId, applicationId, environmentName, hostName, optional)


Add a domain to an environment. If there is no certificate passed in the body, a Let's Encrypt certificate will be generated for this domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **hostName** | **string**| Host Name | 
 **optional** | ***DomainApiEnvironmentAddDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainApiEnvironmentAddDomainOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of Body2**](Body2.md)| Certificate payload | 

### Return type

[**EnvironmentSummary**](EnvironmentSummary.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentRemoveDomain**
> EnvironmentSummary EnvironmentRemoveDomain(ctx, accountId, applicationId, environmentName, hostName)


Remove a domain from an environment. Will also remove any certificate linked to the domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **hostName** | **string**| Host Name | 

### Return type

[**EnvironmentSummary**](EnvironmentSummary.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfiguration**
> HttpsConfiguration GetConfiguration(ctx, accountId, hostName)


Get https configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **hostName** | **string**| Host Name | 

### Return type

[**HttpsConfiguration**](HttpsConfiguration.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostConfiguration**
> HttpsConfiguration PostConfiguration(ctx, accountId, hostName, body)


Set https configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **hostName** | **string**| Host Name | 
  **body** | [**HttpsSetConfiguration**](HttpsSetConfiguration.md)| Https Set Configuration Payload | 

### Return type

[**HttpsConfiguration**](HttpsConfiguration.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewCertificate**
> interface{} RenewCertificate(ctx, accountId, hostName)


Renew certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **hostName** | **string**| Host Name | 

### Return type

**interface{}**

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyEngaged**
> DomainEngaged VerifyEngaged(ctx, accountId, hostName)


Checks if domain is configured to route traffic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **hostName** | **string**| Host Name | 

### Return type

[**DomainEngaged**](DomainEngaged.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

