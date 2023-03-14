# \AccountApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountBillingGet**](AccountApi.md#AccountBillingGet) | **Get** /account/{accountId}/billingsummary | 
[**AccountBillingHistoryGet**](AccountApi.md#AccountBillingHistoryGet) | **Get** /account/{accountId}/billinghistory | 
[**AccountCreate**](AccountApi.md#AccountCreate) | **Post** /account/create | 
[**AccountDomainList**](AccountApi.md#AccountDomainList) | **Get** /account/{accountId}/domains | 
[**AccountGet**](AccountApi.md#AccountGet) | **Get** /account/{accountId} | 
[**AccountGraph**](AccountApi.md#AccountGraph) | **Get** /account/graph | 
[**AccountList**](AccountApi.md#AccountList) | **Get** /account | 
[**AccountUserUpdate**](AccountApi.md#AccountUserUpdate) | **Post** /account/{accountId}/user/{userId} | 


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

# **AccountCreate**
> CreateAccountResponse AccountCreate(ctx, body)


Create a new account and application for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountCreateRequest**](AccountCreateRequest.md)| Account create | 

### Return type

[**CreateAccountResponse**](CreateAccountResponse.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountDomainList**
> DomainList AccountDomainList(ctx, accountId, optional)


Get list of account's domains.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
 **optional** | ***AccountApiAccountDomainListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiAccountDomainListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engaged** | **optional.Bool**| The flag to return engaged or disengaged domains. | 

### Return type

[**DomainList**](DomainList.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountGet**
> Account AccountGet(ctx, accountId)


Get basic information about an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 

### Return type

[**Account**](Account.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountGraph**
> []AccountGraph AccountGraph(ctx, )


Gets a graph of accounts/applications & environments

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccountGraph**](AccountGraph.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountList**
> []Account AccountList(ctx, )


Get the list of accounts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Account**](Account.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountUserUpdate**
> AccountUserUpdateResult AccountUserUpdate(ctx, accountId, userId, body)


Update Account User for permissions changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **userId** | **int64**| The user identifier number | 
  **body** | [**AccountUserUpdateParams**](AccountUserUpdateParams.md)| Updates the accountCapability for a specific user within a specific account | 

### Return type

[**AccountUserUpdateResult**](AccountUserUpdateResult.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

