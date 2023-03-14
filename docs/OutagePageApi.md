# \OutagePageApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutagePageDelete**](OutagePageApi.md#OutagePageDelete) | **Delete** /account/{accountId}/application/{applicationId}/environment/{environmentName}/outagepage | 
[**OutagePageGet**](OutagePageApi.md#OutagePageGet) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/outagepage | 
[**OutagePagePost**](OutagePageApi.md#OutagePagePost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/outagepage | 


# **OutagePageDelete**
> OutagePageDelete(ctx, accountId, applicationId, environmentName)


Disengage the outage page for the specified environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OutagePageGet**
> []OutagePage OutagePageGet(ctx, accountId, applicationId, environmentName)


Lists the available outage pages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

[**[]OutagePage**](OutagePage.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OutagePagePost**
> OutagePagePost(ctx, accountId, applicationId, environmentName, pageName)


Engage the specified outage page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **pageName** | **string**| Name of the outage page to engage | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

