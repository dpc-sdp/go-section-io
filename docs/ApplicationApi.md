# \ApplicationApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountUserGet**](ApplicationApi.md#AccountUserGet) | **Get** /account/{accountId}/user/{userId} | 
[**AccountUserList**](ApplicationApi.md#AccountUserList) | **Get** /account/{accountId}/user | 
[**ApplicationClone**](ApplicationApi.md#ApplicationClone) | **Post** /account/{accountId}/application/{applicationId}/clone | 
[**ApplicationCreate**](ApplicationApi.md#ApplicationCreate) | **Post** /account/{accountId}/application/create | 
[**ApplicationDelete**](ApplicationApi.md#ApplicationDelete) | **Delete** /account/{accountId}/application/{applicationId} | 
[**ApplicationGet**](ApplicationApi.md#ApplicationGet) | **Get** /account/{accountId}/application/{applicationId} | 
[**ApplicationList**](ApplicationApi.md#ApplicationList) | **Get** /account/{accountId}/application/ | 
[**ApplicationSplit**](ApplicationApi.md#ApplicationSplit) | **Post** /account/{accountId}/application/{applicationId}/split | 
[**ApplicationStatePost**](ApplicationApi.md#ApplicationStatePost) | **Post** /account/{accountId}/application/{applicationId}/state | 
[**ApplicationVerifyEngaged**](ApplicationApi.md#ApplicationVerifyEngaged) | **Get** /account/{accountId}/application/{applicationId}/verifyEngaged | 
[**ProxyTemplateInitialState**](ApplicationApi.md#ProxyTemplateInitialState) | **Get** /proxytemplate/initialstate | 
[**ProxyTemplateList**](ApplicationApi.md#ProxyTemplateList) | **Get** /proxytemplate | 
[**ResolveGet**](ApplicationApi.md#ResolveGet) | **Get** /origin | 
[**StackList**](ApplicationApi.md#StackList) | **Get** /stack | 


# **AccountUserGet**
> AccountUser AccountUserGet(ctx, accountId, userId)


Gets list of users with permissions to account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **userId** | **int64**| The user identifier number | 

### Return type

[**AccountUser**](AccountUser.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountUserList**
> []AccountUser AccountUserList(ctx, accountId)


Gets list of users with permissions to account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 

### Return type

[**[]AccountUser**](AccountUser.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationClone**
> Application ApplicationClone(ctx, accountId, applicationId, body)


Clone the configuration of an existing application into a new application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **body** | [**ApplicationCloneRequest**](ApplicationCloneRequest.md)| Application Clone Payload | 

### Return type

[**Application**](Application.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationCreate**
> Application ApplicationCreate(ctx, accountId, body)


Create a new application for the account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **body** | [**ApplicationCreateRequest**](ApplicationCreateRequest.md)| Application Create Payload | 

### Return type

[**Application**](Application.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationDelete**
> ApplicationDelete(ctx, accountId, applicationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationGet**
> Application ApplicationGet(ctx, accountId, applicationId)


Get basic information about an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 

### Return type

[**Application**](Application.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationList**
> []ApplicationSummary ApplicationList(ctx, accountId)


Get the account's applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 

### Return type

[**[]ApplicationSummary**](ApplicationSummary.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationSplit**
> Application ApplicationSplit(ctx, accountId, applicationId, body)


Split the url space of an existing application into a new application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **body** | [**ApplicationSplitRequest**](ApplicationSplitRequest.md)| Application Split Payload | 

### Return type

[**Application**](Application.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationStatePost**
> ApplicationStateUpdateResult ApplicationStatePost(ctx, accountId, applicationId, body)


Apply a cache ban to the default hosted environment for this application.  Deprecated, please use the proxy state method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **body** | [**ApplicationStateUpdateRequest**](ApplicationStateUpdateRequest.md)| State update description | 

### Return type

[**ApplicationStateUpdateResult**](ApplicationStateUpdateResult.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationVerifyEngaged**
> DomainEngaged ApplicationVerifyEngaged(ctx, accountId, applicationId)


Checks if any environment for this application is configured to route through Section

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 

### Return type

[**DomainEngaged**](DomainEngaged.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxyTemplateInitialState**
> *os.File ProxyTemplateInitialState(ctx, proxyTemplateName)


Returns a tar.gz containing the inital state of the proxy repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyTemplateName** | **string**| Proxy template identifier | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxyTemplateList**
> []ProxyTemplate ProxyTemplateList(ctx, )


Returns an array of proxy templates that can be included in a proxy stack

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProxyTemplate**](ProxyTemplate.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResolveGet**
> Origin ResolveGet(ctx, hostName)


Gets origin endpoint from hostname

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostName** | **string**| Host Name | 

### Return type

[**Origin**](Origin.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackList**
> []Stack StackList(ctx, )


Returns an array of stacks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Stack**](Stack.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

