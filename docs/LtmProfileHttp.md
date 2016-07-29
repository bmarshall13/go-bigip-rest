# LtmProfileHttp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**ProxyType** | **string** | Specifies the type of HTTP proxy. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**HeaderInsert** | **string** | Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none. The HTTP header being inserted can include a client IP address. Including a client IP address in an HTTP header is useful when a connection goes through a secure network address translation (SNAT) and you need to preserve the original client IP address. When you assign the configured HTTP profile to a virtual server, the system then inserts the header specified by the profile into any HTTP request that the system sends to a pool or pool member. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ViaRequest** | **string** | Specifies whether to append, remove, or preserve a Via header in an HTTP request. | [optional] [default to null]
**HeaderErase** | **string** | Specifies the header string that you want to erase from an HTTP request. You can also specify none. | [optional] [default to null]
**BasicAuthRealm** | **string** | Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none. | [optional] [default to null]
**FallbackStatusCodes** | **string** | Specifies one or more three-digit status codes that can be returned by an HTTP server. | [optional] [default to null]
**InsertXforwardedFor** | **string** | When using connection pooling, which allows clients to make use of other client requests&#39; server-side connections, you can insert the X-Forwarded-For header and specify a client IP address. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**OneconnectTransformations** | **string** | Enables the system to perform HTTP header transformations for the purpose of keeping server-side connections open. This feature requires configuration of a OneConnect profile. | [optional] [default to null]
**EncryptCookieSecret** | **string** | Specifies a passphrase for the cookie encryption. | [optional] [default to null]
**AcceptXff** | **string** | Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request&#39;s XFF (X-forwarded-for) headers, if they exist. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**ServerAgentName** | **string** | Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is \&quot;BigIP\&quot;. If no string is specified, then no Server header will be added to such responses. | [optional] [default to null]
**LwsWidth** | **int64** | Specifies the maximum number of columns allowed for a header that is inserted into an HTTP request. | [optional] [default to 80]
**ViaHostName** | **string** | Specifies the hostname to include into Via header. | [optional] [default to null]
**RedirectRewrite** | **string** | Specifies which of the application HTTP redirects the system rewrites to HTTPS. Use this feature when the application is generating HTTP redirects that send the client to HTTP (a non-secure channel) when you want the client to continue accessing the application using HTTPS (a secure channel). This is a common occurrence when using client-side SSL processing on a BIG-IP system. | [optional] [default to null]
**ResponseChunking** | **string** | Specifies how to handle chunked and unchunked responses. | [optional] [default to null]
**RequestChunking** | **string** | Specifies how to handle chunked and unchunked requests. | [optional] [default to null]
**XffAlternativeNames** | **string** | Specifies alternative XFF headers instead of the default X-forwarded-for header. | [optional] [default to null]
**EncryptCookies** | **string** | Encrypts specified cookies that the BIG-IP system sends to a client system. | [optional] [default to null]
**LwsSeparator** | **string** | Specifies the linear white space separator that the system should use between HTTP headers when a header exceeds the maximum width specified by the lws width setting. | [optional] [default to null]
**ResponseHeadersPermitted** | **string** | Specifies headers that the BIG-IP system allows in an HTTP response. | [optional] [default to null]
**FallbackHost** | **string** | Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number, or URI path. For example, if all members of the targeted pool are unavailable (that is, the members are disabled, marked as down, or have exceeded their connection limit), the system can redirect the HTTP request to the fallback host, with the HTTP reply Status Code 302 Found. | [optional] [default to null]
**ViaResponse** | **string** | Specifies whether to append, remove, or preserve a Via header in an HTTP response. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


