# LtmProfileHttpEnforcement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnknownMethod** | **string** | Specifies whether to allow, reject or switch to pass-through mode when an unknown HTTP method is parsed. | [optional] [default to null]
**ExcessClientHeaders** | **string** | Specifies the behavior when too many client headers are received. If enabled, will switch to pass through mode instead of rejecting the connection. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**OversizeServerHeaders** | **string** | Specifies the behavior when too-large server headers are received. If enabled, will switch to pass through mode instead of rejecting the connection. | [optional] [default to null]
**ExcessServerHeaders** | **string** | Specifies the behavior when too many server headers are received. If enabled, will switch to pass through mode instead of rejecting the connection. | [optional] [default to null]
**OversizeClientHeaders** | **string** | Specifies the behavior when too-large client headers are received. If enabled, will switch to pass through mode instead of rejecting the connection. | [optional] [default to null]
**MaxRequests** | **int64** | Specifies the number of requests that the system accepts on a per-connection basis. The default value is 0 (zero), which means the system does not limit the number of requests per connection. | [optional] [default to null]
**Pipeline** | **string** | Enables HTTP/1.1 pipelining. This allows clients to make requests even when prior requests have not received a response. In order for this to succeed, however, destination servers must include support for pipelining. If set to pass-through, pipelined data will cause the BigIP to immediately switch to pass-through mode and disable the HTTP filter. | [optional] [default to null]
**TruncatedRedirects** | **string** | Specifies what happens if a truncated redirect is seen from a server. If enabled, the redirect will be forwarded to the client, otherwise the malformed HTTP will be silently ignored. | [optional] [default to null]
**KnownMethods** | **string** | Specifies which HTTP methods count as being known.  Removing RFC-defined methods from this list will cause the HTTP filter to not recognize them. | [optional] [default to null]
**MaxHeaderSize** | **int64** | Specifies the maximum header size. | [optional] [default to 32768]
**MaxHeaderCount** | **int64** | Specifies the maximum number of headers allowed in HTTP request/response. The default is 64 headers. | [optional] [default to 64]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


