# \ProxyApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurationGet**](ProxyApi.md#ConfigurationGet) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/proxy/{proxyName}/configuration | 
[**ConfigurationPost**](ProxyApi.md#ConfigurationPost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/proxy/{proxyName}/configuration | 
[**PatchFile**](ProxyApi.md#PatchFile) | **Patch** /account/{accountId}/application/{applicationId}/environment/{environmentName}/update | 
[**ProxyOperationAction**](ProxyApi.md#ProxyOperationAction) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/proxy/{proxyName}/{operationName} | 
[**ProxyStatePost**](ProxyApi.md#ProxyStatePost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/proxy/{proxyName}/state | 


# **ConfigurationGet**
> ProxyConfiguration ConfigurationGet(ctx, accountId, applicationId, environmentName, proxyName)


Get the proxy's configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **proxyName** | **string**| The proxy identifier | 

### Return type

[**ProxyConfiguration**](ProxyConfiguration.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationPost**
> ErrorModel ConfigurationPost(ctx, accountId, applicationId, environmentName, proxyName, body)


Set the proxy's configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **proxyName** | **string**| The proxy identifier | 
  **body** | [**ProxyConfiguration**](ProxyConfiguration.md)| Proxy configuration update details | 

### Return type

[**ErrorModel**](ErrorModel.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFile**
> interface{} PatchFile(ctx, accountId, applicationId, environmentName, body, filepath)


Patch the proxy with a JSON Patch. See - RFC6902 https://tools.ietf.org/html/rfc6902 - http://jsonpatch.com/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **body** | [**PatchRequest**](PatchRequest.md)|  | 
  **filepath** | **string**| the path to the file you want to modify. Case-sensitive. | 

### Return type

**interface{}**

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxyOperationAction**
> ProxyOperationAction(ctx, accountId, applicationId, environmentName, proxyName, operationName, optional)


Send a message to a specified proxy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **proxyName** | **string**| The proxy identifier | 
  **operationName** | **string**|  | 
 **optional** | ***ProxyApiProxyOperationActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiProxyOperationActionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **operationParameter** | **optional.String**| Parmeter to send with the message to a proxy | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxyStatePost**
> ProxyStatePost(ctx, accountId, applicationId, environmentName, proxyName, optional)


Applies a varnish ban expression to the specified proxy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **proxyName** | **string**| The proxy identifier | 
 **optional** | ***ProxyApiProxyStatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiProxyStatePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **banExpression** | **optional.String**| The varnish ban expression to apply | [default to req.url ~ /]
 **async** | **optional.Bool**| If true the call will return immediately, not waiting for all delivery nodes to complete the ban | [default to false]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

