# \ZoneApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteZoneRecord**](ZoneApi.md#DeleteZoneRecord) | **Delete** /account/{accountId}/zone/{zoneName}/{recordName}/{recordType} | 
[**EnableHostedDNS**](ZoneApi.md#EnableHostedDNS) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/enableHostedDNS | 
[**GetDomainZone**](ZoneApi.md#GetDomainZone) | **Get** /account/{accountId}/domain/{hostName}/zone | 
[**GetZoneDetail**](ZoneApi.md#GetZoneDetail) | **Get** /account/{accountId}/zone/{zoneName} | 
[**GetZones**](ZoneApi.md#GetZones) | **Get** /account/{accountId}/zone | 
[**PostZoneRecord**](ZoneApi.md#PostZoneRecord) | **Post** /account/{accountId}/zone/{zoneName} | 


# **DeleteZoneRecord**
> ZoneDetail DeleteZoneRecord(ctx, accountId, zoneName, recordName, recordType)


Delete a record from a zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **zoneName** | **string**| Zone Name | 
  **recordName** | **string**| Record Name | 
  **recordType** | **string**| Record Type | 

### Return type

[**ZoneDetail**](ZoneDetail.md)

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

# **GetDomainZone**
> ZoneCandidate GetDomainZone(ctx, accountId, hostName)


Find likely zone for domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **hostName** | **string**| Host Name | 

### Return type

[**ZoneCandidate**](ZoneCandidate.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZoneDetail**
> ZoneDetail GetZoneDetail(ctx, accountId, zoneName)


Get details of a zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **zoneName** | **string**| Zone Name | 

### Return type

[**ZoneDetail**](ZoneDetail.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZones**
> []ZoneSummary GetZones(ctx, accountId)


List all zones for account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 

### Return type

[**[]ZoneSummary**](ZoneSummary.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostZoneRecord**
> ZoneDetail PostZoneRecord(ctx, accountId, zoneName, body)


Add or update a record for a zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **zoneName** | **string**| Zone Name | 
  **body** | [**Body1**](Body1.md)|  | 

### Return type

[**ZoneDetail**](ZoneDetail.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

