# LtmProfileHttpExplicitProxy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**DnsErrorMessage** | **string** | Specifies the error message that will be returned to the browser when a proxy request can&#39;t be completed because of a failure to resolve the hostname in the request. | [optional] [default to null]
**BadRequestMessage** | **string** | Specifies the error message that will be returned to the browser when a proxy request can&#39;t be completed because the request was malformed. | [optional] [default to null]
**BadResponseMessage** | **string** | Specifies the error message that will be returned to the browser when a proxy request can&#39;t be completed because the response was malformed. | [optional] [default to null]
**DnsResolver** | **string** | Specifies the dns-resolver object that will be used to resolve hostnames in proxy requests. | [optional] [default to null]
**TunnelName** | **string** | Specifies the tunnel that will be used for outbound proxy requests.  This enables other virtual servers to receive connections initiated by the proxy service. | [optional] [default to null]
**HostNames** | **string** | Specifies the which host names are to be treated as local. Proxy requests made for those hosts will be treated as regular HTTP requests and will be sent to the configured default pool. | [optional] [default to null]
**DefaultConnectHandling** | **string** | Specifies the behavior of the proxy service for CONNECT requests. If set to &#39;deny&#39;, CONNECT requests will only be honored if there is another virtual server listening for the requested outbound connection. If set to &#39;allow&#39; outbound connections will be made regardless of other virtual servers. | [optional] [default to null]
**RouteDomain** | **string** | Specifies the route-domain that will be used for outbound proxy requests. | [optional] [default to null]
**ConnectErrorMessage** | **string** | Specifies the error message that will be returned to the browser when a proxy request can&#39;t be completed because of a failure to establish the outbound connection. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


