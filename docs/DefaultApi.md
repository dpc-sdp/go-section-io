# \DefaultApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CacheClear**](DefaultApi.md#CacheClear) | **Post** /user/{userId}/cache | 
[**LinkExternal**](DefaultApi.md#LinkExternal) | **Post** /user/linkexternal | 


# **CacheClear**
> CacheClear(ctx, userId)


Clear internal cache for user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| The user identifier number | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkExternal**
> LinkExternal(ctx, body)


Associate your user account with an external IDP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)| Account create | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

