# \HelpApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HelpGet**](HelpApi.md#HelpGet) | **Get** /help/{helpId} | 


# **HelpGet**
> HelpContent HelpGet(ctx, helpId)


Gets contextual help content by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **helpId** | **string**| Identifier for contextual help | 

### Return type

[**HelpContent**](HelpContent.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

