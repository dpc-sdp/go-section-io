# \UserApi

All URIs are relative to *https://aperture.section.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountUserGet**](UserApi.md#AccountUserGet) | **Get** /account/{accountId}/user/{userId} | 
[**AccountUserInvitePost**](UserApi.md#AccountUserInvitePost) | **Post** /account/{accountId}/user | 
[**AccountUserList**](UserApi.md#AccountUserList) | **Get** /account/{accountId}/user | 
[**AccountUserRemove**](UserApi.md#AccountUserRemove) | **Delete** /account/{accountId}/user/{userId} | 
[**AccountUserUpdate**](UserApi.md#AccountUserUpdate) | **Post** /account/{accountId}/user/{userId} | 
[**CurrentUserGet**](UserApi.md#CurrentUserGet) | **Get** /user | 
[**CurrentUserPost**](UserApi.md#CurrentUserPost) | **Post** /user | 
[**CurrentUserTotpPost**](UserApi.md#CurrentUserTotpPost) | **Post** /user/totp | 
[**CurrentUserVerificationPost**](UserApi.md#CurrentUserVerificationPost) | **Post** /user/resendverification | 


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

# **AccountUserInvitePost**
> AccountUserActionResult AccountUserInvitePost(ctx, accountId, body)


Invite user to account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **body** | [**AccountUserInviteRequest**](AccountUserInviteRequest.md)| Invite user to account | 

### Return type

[**AccountUserActionResult**](AccountUserActionResult.md)

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

# **AccountUserRemove**
> AccountUserActionResult AccountUserRemove(ctx, accountId, userId)


Invite user to account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| The account identifier number | 
  **userId** | **int64**| The user identifier number | 

### Return type

[**AccountUserActionResult**](AccountUserActionResult.md)

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

# **CurrentUserGet**
> User CurrentUserGet(ctx, )


Get current user details

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrentUserPost**
> User CurrentUserPost(ctx, body)


Update current user details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)| User details | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrentUserTotpPost**
> Totp CurrentUserTotpPost(ctx, )


Generate and return a new 2FA (TOTP) token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Totp**](Totp.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrentUserVerificationPost**
> UserResendVerificationResult CurrentUserVerificationPost(ctx, )


Resend email verification

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserResendVerificationResult**](UserResendVerificationResult.md)

### Authorization

[basic](../README.md#basic), [sectionToken](../README.md#sectionToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/tar+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

