# \EnvironmentApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EgressGet**](EnvironmentApi.md#EgressGet) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/egress | 
[**EgressPost**](EnvironmentApi.md#EgressPost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/egress | 
[**EnableHostedDNS**](EnvironmentApi.md#EnableHostedDNS) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/enableHostedDNS | 
[**EnvironmentAddDomain**](EnvironmentApi.md#EnvironmentAddDomain) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/domain/{hostName} | 
[**EnvironmentCreate**](EnvironmentApi.md#EnvironmentCreate) | **Post** /account/{accountId}/application/{applicationId}/environment/create | 
[**EnvironmentDnsBypassDelete**](EnvironmentApi.md#EnvironmentDnsBypassDelete) | **Delete** /account/{accountId}/application/{applicationId}/environment/{environmentName}/dnsbypass | 
[**EnvironmentDnsBypassPost**](EnvironmentApi.md#EnvironmentDnsBypassPost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/dnsbypass | 
[**EnvironmentFetchEngaged**](EnvironmentApi.md#EnvironmentFetchEngaged) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/engaged | 
[**EnvironmentGet**](EnvironmentApi.md#EnvironmentGet) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName} | 
[**EnvironmentIpRestrictionsGet**](EnvironmentApi.md#EnvironmentIpRestrictionsGet) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/ipRestrictions | 
[**EnvironmentIpRestrictionsPost**](EnvironmentApi.md#EnvironmentIpRestrictionsPost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/ipRestrictions | 
[**EnvironmentList**](EnvironmentApi.md#EnvironmentList) | **Get** /account/{accountId}/application/{applicationId}/environment | 
[**EnvironmentRemoveDomain**](EnvironmentApi.md#EnvironmentRemoveDomain) | **Delete** /account/{accountId}/application/{applicationId}/environment/{environmentName}/domain/{hostName} | 
[**EnvironmentStackGet**](EnvironmentApi.md#EnvironmentStackGet) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/stack | 
[**EnvironmentVerifyEngaged**](EnvironmentApi.md#EnvironmentVerifyEngaged) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/verifyEngaged | 
[**GetConfiguration**](EnvironmentApi.md#GetConfiguration) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/configuration | 
[**PatchConfiguration**](EnvironmentApi.md#PatchConfiguration) | **Patch** /account/{accountId}/application/{applicationId}/environment/{environmentName}/configuration | 


# **EgressGet**
> Egress EgressGet(ctx, accountId, applicationId, environmentName)


Get the environment's egress configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

[**Egress**](Egress.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EgressPost**
> Egress EgressPost(ctx, accountId, applicationId, environmentName, body)


Get the environment's egress configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **body** | [**Egress**](Egress.md)| Environment Egress Payload | 

### Return type

[**Egress**](Egress.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableHostedDNS**
> ZoneSummary EnableHostedDNS(ctx, accountId, applicationId, environmentName, body)


Enable Section Hosted DNS for the environment's domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **body** | [**Body3**](Body3.md)|  | 

### Return type

[**ZoneSummary**](ZoneSummary.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 **optional** | ***EnvironmentApiEnvironmentAddDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentApiEnvironmentAddDomainOpts struct

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

# **EnvironmentCreate**
> EnvironmentSummary EnvironmentCreate(ctx, accountId, applicationId, body)


Create an application environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **body** | [**EnvironmentCreateRequest**](EnvironmentCreateRequest.md)| Environment Create Payload | 

### Return type

[**EnvironmentSummary**](EnvironmentSummary.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentDnsBypassDelete**
> EnvironmentDnsBypassDelete(ctx, accountId, applicationId, environmentName)


Remove the bypass of Section for this environment. All traffic will be directed to Section's servers.

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

# **EnvironmentDnsBypassPost**
> EnvironmentDnsBypassPost(ctx, accountId, applicationId, environmentName)


Request the bypassing of Section for this environment. All traffic will be directed to the environment's origin address.

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

# **EnvironmentFetchEngaged**
> EnvironmentEngaged EnvironmentFetchEngaged(ctx, accountId, applicationId, environmentName)


Returns the last known state of whether this environment is configured to route through Section without re-checking.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

[**EnvironmentEngaged**](EnvironmentEngaged.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentGet**
> EnvironmentSummary EnvironmentGet(ctx, accountId, applicationId, environmentName)


Get the environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

[**EnvironmentSummary**](EnvironmentSummary.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentIpRestrictionsGet**
> IpRestrictions EnvironmentIpRestrictionsGet(ctx, accountId, applicationId, environmentName)


Returns the list of IP addresses and CIDR blocks that are restricted from accessing this environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

[**IpRestrictions**](IpRestrictions.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentIpRestrictionsPost**
> IpRestrictions EnvironmentIpRestrictionsPost(ctx, accountId, applicationId, environmentName, body)


Update the list of IP addresses and CIDR blocks that are restricted from accessing this environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **body** | [**IpRestrictions**](IpRestrictions.md)| IP restrictions list | 

### Return type

[**IpRestrictions**](IpRestrictions.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentList**
> []EnvironmentSummary EnvironmentList(ctx, accountId, applicationId)


Get the application's environments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 

### Return type

[**[]EnvironmentSummary**](EnvironmentSummary.md)

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

# **EnvironmentStackGet**
> []Proxy EnvironmentStackGet(ctx, accountId, applicationId, environmentName)


Get the proxy stack list for an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

[**[]Proxy**](Proxy.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentVerifyEngaged**
> EnvironmentEngaged EnvironmentVerifyEngaged(ctx, accountId, applicationId, environmentName)


Checks if this environment is configured to route through Section

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

[**EnvironmentEngaged**](EnvironmentEngaged.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfiguration**
> EnvironmentConfiguration GetConfiguration(ctx, accountId, applicationId, environmentName)


Get configuration for the environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 

### Return type

[**EnvironmentConfiguration**](EnvironmentConfiguration.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchConfiguration**
> EnvironmentConfiguration PatchConfiguration(ctx, accountId, applicationId, environmentName, body)


Patch configuration for the environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **applicationId** | **int64**| The application identifier number | 
  **environmentName** | **string**| The name of the environment | 
  **body** | [**PatchRequest**](PatchRequest.md)|  | 

### Return type

[**EnvironmentConfiguration**](EnvironmentConfiguration.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

