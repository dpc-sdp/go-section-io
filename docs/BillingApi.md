# \BillingApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountBillingGet**](BillingApi.md#AccountBillingGet) | **Get** /account/{accountId}/billingsummary | 
[**AccountBillingHistoryGet**](BillingApi.md#AccountBillingHistoryGet) | **Get** /account/{accountId}/billinghistory | 


# **AccountBillingGet**
> BillingSummary AccountBillingGet(ctx, accountId)


Get basic billing information for account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 

### Return type

[**BillingSummary**](BillingSummary.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountBillingHistoryGet**
> []BillingHistory AccountBillingHistoryGet(ctx, accountId)


Get billing history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 

### Return type

[**[]BillingHistory**](BillingHistory.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

