# \SharedApi

All URIs are relative to *https://localhost/mgmt*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](SharedApi.md#Login) | **Post** /shared/authn/login | 


# **Login**
> LoginResp Login($loginBody)



Login to generate an API token


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginBody** | [**LoginBody**](LoginBody.md)| User to authenticate as | 

### Return type

[**LoginResp**](loginResp.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

