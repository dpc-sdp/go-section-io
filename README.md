# Go API client for sectionio

Get edgey with the Section API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./sectionio"
```

## Documentation for API Endpoints

All URIs are relative to *https://aperture.section.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**AccountBillingGet**](docs/AccountApi.md#accountbillingget) | **Get** /account/{accountId}/billingsummary | 
*AccountApi* | [**AccountBillingHistoryGet**](docs/AccountApi.md#accountbillinghistoryget) | **Get** /account/{accountId}/billinghistory | 
*AccountApi* | [**AccountCreate**](docs/AccountApi.md#accountcreate) | **Post** /account/create | 
*AccountApi* | [**AccountDomainList**](docs/AccountApi.md#accountdomainlist) | **Get** /account/{accountId}/domains | 
*AccountApi* | [**AccountGet**](docs/AccountApi.md#accountget) | **Get** /account/{accountId} | 
*AccountApi* | [**AccountGraph**](docs/AccountApi.md#accountgraph) | **Get** /account/graph | 
*AccountApi* | [**AccountList**](docs/AccountApi.md#accountlist) | **Get** /account | 
*AccountApi* | [**AccountUserUpdate**](docs/AccountApi.md#accountuserupdate) | **Post** /account/{accountId}/user/{userId} | 
*ApplicationApi* | [**AccountUserGet**](docs/ApplicationApi.md#accountuserget) | **Get** /account/{accountId}/user/{userId} | 
*ApplicationApi* | [**AccountUserList**](docs/ApplicationApi.md#accountuserlist) | **Get** /account/{accountId}/user | 
*ApplicationApi* | [**ApplicationClone**](docs/ApplicationApi.md#applicationclone) | **Post** /account/{accountId}/application/{applicationId}/clone | 
*ApplicationApi* | [**ApplicationCreate**](docs/ApplicationApi.md#applicationcreate) | **Post** /account/{accountId}/application/create | 
*ApplicationApi* | [**ApplicationDelete**](docs/ApplicationApi.md#applicationdelete) | **Delete** /account/{accountId}/application/{applicationId} | 
*ApplicationApi* | [**ApplicationGet**](docs/ApplicationApi.md#applicationget) | **Get** /account/{accountId}/application/{applicationId} | 
*ApplicationApi* | [**ApplicationList**](docs/ApplicationApi.md#applicationlist) | **Get** /account/{accountId}/application/ | 
*ApplicationApi* | [**ApplicationSplit**](docs/ApplicationApi.md#applicationsplit) | **Post** /account/{accountId}/application/{applicationId}/split | 
*ApplicationApi* | [**ApplicationStatePost**](docs/ApplicationApi.md#applicationstatepost) | **Post** /account/{accountId}/application/{applicationId}/state | 
*ApplicationApi* | [**ApplicationVerifyEngaged**](docs/ApplicationApi.md#applicationverifyengaged) | **Get** /account/{accountId}/application/{applicationId}/verifyEngaged | 
*ApplicationApi* | [**ProxyTemplateInitialState**](docs/ApplicationApi.md#proxytemplateinitialstate) | **Get** /proxytemplate/initialstate | 
*ApplicationApi* | [**ProxyTemplateList**](docs/ApplicationApi.md#proxytemplatelist) | **Get** /proxytemplate | 
*ApplicationApi* | [**ResolveGet**](docs/ApplicationApi.md#resolveget) | **Get** /origin | 
*ApplicationApi* | [**StackList**](docs/ApplicationApi.md#stacklist) | **Get** /stack | 
*BillingApi* | [**AccountBillingGet**](docs/BillingApi.md#accountbillingget) | **Get** /account/{accountId}/billingsummary | 
*BillingApi* | [**AccountBillingHistoryGet**](docs/BillingApi.md#accountbillinghistoryget) | **Get** /account/{accountId}/billinghistory | 
*DefaultApi* | [**CacheClear**](docs/DefaultApi.md#cacheclear) | **Post** /user/{userId}/cache | 
*DefaultApi* | [**LinkExternal**](docs/DefaultApi.md#linkexternal) | **Post** /user/linkexternal | 
*DomainApi* | [**EnvironmentAddDomain**](docs/DomainApi.md#environmentadddomain) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/domain/{hostName} | 
*DomainApi* | [**EnvironmentRemoveDomain**](docs/DomainApi.md#environmentremovedomain) | **Delete** /account/{accountId}/application/{applicationId}/environment/{environmentName}/domain/{hostName} | 
*DomainApi* | [**GetConfiguration**](docs/DomainApi.md#getconfiguration) | **Get** /account/{accountId}/domain/{hostName}/https | 
*DomainApi* | [**PostConfiguration**](docs/DomainApi.md#postconfiguration) | **Post** /account/{accountId}/domain/{hostName}/https | 
*DomainApi* | [**RenewCertificate**](docs/DomainApi.md#renewcertificate) | **Post** /account/{accountId}/domain/{hostName}/renewCertificate | 
*DomainApi* | [**VerifyEngaged**](docs/DomainApi.md#verifyengaged) | **Get** /account/{accountId}/domain/{hostName}/verifyEngaged | 
*EnvironmentApi* | [**EgressGet**](docs/EnvironmentApi.md#egressget) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/egress | 
*EnvironmentApi* | [**EgressPost**](docs/EnvironmentApi.md#egresspost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/egress | 
*EnvironmentApi* | [**EnableHostedDNS**](docs/EnvironmentApi.md#enablehosteddns) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/enableHostedDNS | 
*EnvironmentApi* | [**EnvironmentAddDomain**](docs/EnvironmentApi.md#environmentadddomain) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/domain/{hostName} | 
*EnvironmentApi* | [**EnvironmentCreate**](docs/EnvironmentApi.md#environmentcreate) | **Post** /account/{accountId}/application/{applicationId}/environment/create | 
*EnvironmentApi* | [**EnvironmentDnsBypassDelete**](docs/EnvironmentApi.md#environmentdnsbypassdelete) | **Delete** /account/{accountId}/application/{applicationId}/environment/{environmentName}/dnsbypass | 
*EnvironmentApi* | [**EnvironmentDnsBypassPost**](docs/EnvironmentApi.md#environmentdnsbypasspost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/dnsbypass | 
*EnvironmentApi* | [**EnvironmentFetchEngaged**](docs/EnvironmentApi.md#environmentfetchengaged) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/engaged | 
*EnvironmentApi* | [**EnvironmentGet**](docs/EnvironmentApi.md#environmentget) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName} | 
*EnvironmentApi* | [**EnvironmentIpRestrictionsGet**](docs/EnvironmentApi.md#environmentiprestrictionsget) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/ipRestrictions | 
*EnvironmentApi* | [**EnvironmentIpRestrictionsPost**](docs/EnvironmentApi.md#environmentiprestrictionspost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/ipRestrictions | 
*EnvironmentApi* | [**EnvironmentList**](docs/EnvironmentApi.md#environmentlist) | **Get** /account/{accountId}/application/{applicationId}/environment | 
*EnvironmentApi* | [**EnvironmentRemoveDomain**](docs/EnvironmentApi.md#environmentremovedomain) | **Delete** /account/{accountId}/application/{applicationId}/environment/{environmentName}/domain/{hostName} | 
*EnvironmentApi* | [**EnvironmentStackGet**](docs/EnvironmentApi.md#environmentstackget) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/stack | 
*EnvironmentApi* | [**EnvironmentVerifyEngaged**](docs/EnvironmentApi.md#environmentverifyengaged) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/verifyEngaged | 
*EnvironmentApi* | [**GetConfiguration**](docs/EnvironmentApi.md#getconfiguration) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/configuration | 
*EnvironmentApi* | [**PatchConfiguration**](docs/EnvironmentApi.md#patchconfiguration) | **Patch** /account/{accountId}/application/{applicationId}/environment/{environmentName}/configuration | 
*HelpApi* | [**HelpGet**](docs/HelpApi.md#helpget) | **Get** /help/{helpId} | 
*OutagePageApi* | [**OutagePageDelete**](docs/OutagePageApi.md#outagepagedelete) | **Delete** /account/{accountId}/application/{applicationId}/environment/{environmentName}/outagepage | 
*OutagePageApi* | [**OutagePageGet**](docs/OutagePageApi.md#outagepageget) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/outagepage | 
*OutagePageApi* | [**OutagePagePost**](docs/OutagePageApi.md#outagepagepost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/outagepage | 
*ProxyApi* | [**ConfigurationGet**](docs/ProxyApi.md#configurationget) | **Get** /account/{accountId}/application/{applicationId}/environment/{environmentName}/proxy/{proxyName}/configuration | 
*ProxyApi* | [**ConfigurationPost**](docs/ProxyApi.md#configurationpost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/proxy/{proxyName}/configuration | 
*ProxyApi* | [**PatchFile**](docs/ProxyApi.md#patchfile) | **Patch** /account/{accountId}/application/{applicationId}/environment/{environmentName}/update | 
*ProxyApi* | [**ProxyOperationAction**](docs/ProxyApi.md#proxyoperationaction) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/proxy/{proxyName}/{operationName} | 
*ProxyApi* | [**ProxyStatePost**](docs/ProxyApi.md#proxystatepost) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/proxy/{proxyName}/state | 
*UserApi* | [**AccountUserGet**](docs/UserApi.md#accountuserget) | **Get** /account/{accountId}/user/{userId} | 
*UserApi* | [**AccountUserInvitePost**](docs/UserApi.md#accountuserinvitepost) | **Post** /account/{accountId}/user | 
*UserApi* | [**AccountUserList**](docs/UserApi.md#accountuserlist) | **Get** /account/{accountId}/user | 
*UserApi* | [**AccountUserRemove**](docs/UserApi.md#accountuserremove) | **Delete** /account/{accountId}/user/{userId} | 
*UserApi* | [**AccountUserUpdate**](docs/UserApi.md#accountuserupdate) | **Post** /account/{accountId}/user/{userId} | 
*UserApi* | [**CurrentUserGet**](docs/UserApi.md#currentuserget) | **Get** /user | 
*UserApi* | [**CurrentUserPost**](docs/UserApi.md#currentuserpost) | **Post** /user | 
*UserApi* | [**CurrentUserTotpPost**](docs/UserApi.md#currentusertotppost) | **Post** /user/totp | 
*UserApi* | [**CurrentUserVerificationPost**](docs/UserApi.md#currentuserverificationpost) | **Post** /user/resendverification | 
*ZoneApi* | [**DeleteZoneRecord**](docs/ZoneApi.md#deletezonerecord) | **Delete** /account/{accountId}/zone/{zoneName}/{recordName}/{recordType} | 
*ZoneApi* | [**EnableHostedDNS**](docs/ZoneApi.md#enablehosteddns) | **Post** /account/{accountId}/application/{applicationId}/environment/{environmentName}/enableHostedDNS | 
*ZoneApi* | [**GetDomainZone**](docs/ZoneApi.md#getdomainzone) | **Get** /account/{accountId}/domain/{hostName}/zone | 
*ZoneApi* | [**GetZoneDetail**](docs/ZoneApi.md#getzonedetail) | **Get** /account/{accountId}/zone/{zoneName} | 
*ZoneApi* | [**GetZones**](docs/ZoneApi.md#getzones) | **Get** /account/{accountId}/zone | 
*ZoneApi* | [**PostZoneRecord**](docs/ZoneApi.md#postzonerecord) | **Post** /account/{accountId}/zone/{zoneName} | 


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountCreateRequest](docs/AccountCreateRequest.md)
 - [AccountGraph](docs/AccountGraph.md)
 - [AccountGraphApplications](docs/AccountGraphApplications.md)
 - [AccountGraphEnvironments](docs/AccountGraphEnvironments.md)
 - [AccountUser](docs/AccountUser.md)
 - [AccountUserActionResult](docs/AccountUserActionResult.md)
 - [AccountUserInviteRequest](docs/AccountUserInviteRequest.md)
 - [AccountUserUpdateParams](docs/AccountUserUpdateParams.md)
 - [AccountUserUpdateResult](docs/AccountUserUpdateResult.md)
 - [Application](docs/Application.md)
 - [ApplicationCloneRequest](docs/ApplicationCloneRequest.md)
 - [ApplicationCreateRequest](docs/ApplicationCreateRequest.md)
 - [ApplicationSplitRequest](docs/ApplicationSplitRequest.md)
 - [ApplicationStateUpdateRequest](docs/ApplicationStateUpdateRequest.md)
 - [ApplicationStateUpdateResult](docs/ApplicationStateUpdateResult.md)
 - [ApplicationSummary](docs/ApplicationSummary.md)
 - [BillingHistory](docs/BillingHistory.md)
 - [BillingSummary](docs/BillingSummary.md)
 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [Body2](docs/Body2.md)
 - [Body3](docs/Body3.md)
 - [CreateAccountResponse](docs/CreateAccountResponse.md)
 - [CreateAccountResponseApplications](docs/CreateAccountResponseApplications.md)
 - [DomainEngaged](docs/DomainEngaged.md)
 - [DomainList](docs/DomainList.md)
 - [DomainListInner](docs/DomainListInner.md)
 - [DomainSummary](docs/DomainSummary.md)
 - [Egress](docs/Egress.md)
 - [EgressOrigin](docs/EgressOrigin.md)
 - [EgressOriginName](docs/EgressOriginName.md)
 - [EgressOrigins](docs/EgressOrigins.md)
 - [EnvironmentConfiguration](docs/EnvironmentConfiguration.md)
 - [EnvironmentCreateRequest](docs/EnvironmentCreateRequest.md)
 - [EnvironmentEngaged](docs/EnvironmentEngaged.md)
 - [EnvironmentSummary](docs/EnvironmentSummary.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [HelpContent](docs/HelpContent.md)
 - [HostAddressDns](docs/HostAddressDns.md)
 - [HostAddressDnsOrIp](docs/HostAddressDnsOrIp.md)
 - [HttpHeaderFieldName](docs/HttpHeaderFieldName.md)
 - [HttpsConfiguration](docs/HttpsConfiguration.md)
 - [HttpsSetConfiguration](docs/HttpsSetConfiguration.md)
 - [IpRestrictions](docs/IpRestrictions.md)
 - [MultiErrorModel](docs/MultiErrorModel.md)
 - [Origin](docs/Origin.md)
 - [OutagePage](docs/OutagePage.md)
 - [PatchRequest](docs/PatchRequest.md)
 - [Proxy](docs/Proxy.md)
 - [ProxyConfiguration](docs/ProxyConfiguration.md)
 - [ProxyTemplate](docs/ProxyTemplate.md)
 - [ProxyType](docs/ProxyType.md)
 - [Rfc6902Operation](docs/Rfc6902Operation.md)
 - [Stack](docs/Stack.md)
 - [Totp](docs/Totp.md)
 - [User](docs/User.md)
 - [UserResendVerificationResult](docs/UserResendVerificationResult.md)
 - [ZoneCandidate](docs/ZoneCandidate.md)
 - [ZoneDetail](docs/ZoneDetail.md)
 - [ZoneRecord](docs/ZoneRecord.md)
 - [ZoneSummary](docs/ZoneSummary.md)


## Documentation For Authorization

## basic
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```
## sectionToken
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author



