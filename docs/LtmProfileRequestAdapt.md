# LtmProfileRequestAdapt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**PreviewSize** | **int64** | Specifies the maximum size of the preview buffer. The preview buffer is used to hold a copy of the HTTP request header and data sent to the internal virtual server in case the adaptation server reports that it does not need to adapt the HTTP request. Setting the preview-size to zero, disables buffering the request and should only be done if the adaptation server will always return with a modified HTTP request or the original HTTP request. The default value is 1024. | [optional] [default to 1024]
**Name** | **string** | Name of entity | [optional] [default to null]
**ServiceDownAction** | **string** | Specifies the action to take if the internal virtual server does not exist or returns an error. The default value is ignore. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**Enabled** | **string** | Enables adaptation of HTTP requests. If set to yes, HTTP requests will be forwarded to the specified internal virtual server for adaptation. The default value is yes. | [optional] [default to null]
**AllowHttp10** | **string** | Specifies whether to forward HTTP version 1.0 requests for adaptation. By default only HTTP version 1.1 requests are forwarded. Version 1.0 is not supported. While it should work in most cases, it might be necessary to restrict adaptation on a site-specific basis. The default value is no. | [optional] [default to null]
**Timeout** | **int64** | Specifies a timeout in milliseconds. If the internal virtual server does not return a result within the specified time, a timeout error will occur. A 0 value disables the timeout. The default value is 0. | [optional] [default to 0]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**InternalVirtual** | **string** | Specifies the name of the internal virtual server to use for adapting the HTTP request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


