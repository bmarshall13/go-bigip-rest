# SysHttpd

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslProtocol** | **string** | The list of SSL Protocols to accept on the management console. | [optional] [default to null]
**LogLevel** | **string** | Specifies the minimum httpd message level to include in the system log. The default value is warn. | [optional] [default to null]
**RequestBodyMinRate** | **int64** | Specifies, in bytes per second, the minimum average rate at which the request body must be received. The default value is 500. | [optional] [default to null]
**SslOcspResponseTimeSkew** | **int64** | Specifies the maximum allowable time skew in seconds for OCSP response validation. The default is 300 seconds. | [optional] [default to null]
**MaxClients** | **int64** | Specifies the maximum number of concurrent connections to the GUI. The default value is 10. | [optional] [default to null]
**SslCaCertFile** | **string** | Specifies the name of the file that contains the Certificate Authority (CA) certificate file. The default id none. | [optional] [default to null]
**SslPort** | **int64** |  | [optional] [default to null]
**RequestHeaderMaxTimeout** | **int64** | Specifies, in seconds, the maximum time allowed to receive all of the request headers, if the request-header-min-rate option is used, in which case the timeout is extended as more data arrives. Ignored if request-header-min-rate is not used. A value of 0 means no limit. The default value is 40. | [optional] [default to null]
**SslCiphersuite** | **string** | Specifies the ciphers that the system uses. | [optional] [default to null]
**SslVerifyClient** | **string** | Specifies if the client certificate needs to be verified for SSL session establishment. The default is none. | [optional] [default to null]
**FastcgiTimeout** | **int64** | Specifies the seconds before FastCGI timeout. | [optional] [default to null]
**SslOcspEnable** | **string** | Specifies OCSP validation of the client certificate chain. The default is off. | [optional] [default to null]
**SslOcspDefaultResponder** | **string** | Specifies the default responder URI for OCSP validation. The default is http://localhost.localdomain. The default responder value should always be preceded with http:// | [optional] [default to null]
**Include** | **string** | Warning: Do not use this parameter without assistance from the F5 Technical Support team. The system does not validate the commands issued using the include parameter. If you use this parameter incorrectly, you put the functionality of the system at risk. | [optional] [default to null]
**SslOcspOverrideResponder** | **string** | Specifies the force use of default responder URI for OCSP validation. The default is off. | [optional] [default to null]
**AuthPamValidateIp** | **string** | Specifies whether the check for consistent inbound IP for the entire web session is enforced or not. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**SslOcspResponderTimeout** | **int64** | Specifies the maximum allowable time in seconds for OCSP responses. The default value is 300 seconds. | [optional] [default to null]
**RequestHeaderTimeout** | **int64** | Specifies, in seconds, the time allowed to receive all of the request headers. This time includes completion of the SSL handshake. A value of 0 means no limit. If you use the request-header-min-rate option, this represents the initial value for the timeout, which will be extended as more data arrives. The default value is 20. | [optional] [default to null]
**SslCertchainfile** | **string** | Specifies the name of the file that contains the SSL certificate chain. The default is none. | [optional] [default to null]
**SslOcspResponseMaxAge** | **int64** | Specifies the maximum allowable age in seconds for OCSP responses. A value of -1 specifies that a maximum age is not enforced. The default value is -1. | [optional] [default to null]
**RedirectHttpToHttps** | **string** | Specifies whether the system should redirect HTTP requests targeted at the configuration utility to HTTPS. The default value is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**RequestBodyTimeout** | **int64** | Specifies, in seconds, the time allowed to receive all of the request body. A value of 0 means no limit. If you use the request-body-min-rate option, this represents the initial value for the timeout, which will be extended as more data arrives. The default value is 60. | [optional] [default to null]
**RequestHeaderMinRate** | **int64** | Specifies, in bytes per second, the minimum average rate at which the request headers must be received. The default value is 500. | [optional] [default to null]
**HostnameLookup** | **string** | Specifies whether to lookup hostname or not. The default value is off. | [optional] [default to null]
**RequestBodyMaxTimeout** | **int64** | Specifies, in seconds, the maximum time allowed to receive all of the request body, if the request-body-min-rate option is used, in which case the timeout is extended as more data arrives. Ignored if request-body-min-rate is not used. A value of 0 means no limit. The default value is 0. | [optional] [default to null]
**AuthPamDashboardTimeout** | **string** | Specifies whether browser session timeout occurs when the dashboard is running. The default value is disabled. | [optional] [default to null]
**AuthPamIdleTimeout** | **int64** | Specifies the seconds before GUI session timeout. | [optional] [default to null]
**Allow** | **string** | Adds or deletes IP addresses, partial IP addresses, and IP address ranges, hostnames, partial hostnames, domain names, partial domain names, and network and netmask pairs for the HTTP clients from which the httpd daemon accepts requests. The default value is all. Warning Using the value none resets the httpd daemon to allow all HTTP clients access to the system. F5 recommends that you do not use the value none with the httpd command. | [optional] [default to null]
**AuthName** | **string** | Specifies the name for the authentication realm. The default value is BIG-IP. | [optional] [default to null]
**SslInclude** | **string** | Warning: Do not use this parameter without assistance from the F5 Technical Support team. The system does not validate the commands issued using the include parameter. If you use this parameter incorrectly, you put the functionality of the system at risk. | [optional] [default to null]
**SslCertfile** | **string** | Specifies the name of the file that contains the SSL certificate. The default value is /etc/httpd/conf/ssl.crt/server.crt. Note that the path to the file must start with /etc/httpd/conf/ssl.crt/ or /config/httpd/conf/ssl.crt/ unless the path is a relative path.  If the path is a relative path, then it must start with conf/ssl.crt/. | [optional] [default to null]
**SslVerifyDepth** | **int64** | Specifies maximum depth of CA certificates in client certificate verification. The default is 10. | [optional] [default to null]
**SslCertkeyfile** | **string** | Specifies the name of the file that contains the SSL certificate key. The default value is /etc/httpd/conf/ssl.key/server.key. Note that the path to the file must start with /etc/httpd/conf/ssl.key/ or /config/httpd/conf/ssl.key/ unless the path is a relative path. If the path is a relative path, then it must start with conf/ssl.key/. When you change the key file, you must also change the certificate file. In other words, the following command does not work to change the key: bigpipe httpd sslcertkeyfile  string . Instead, you must use this command: { bigpipe httpd sslcertfile  string  sslcerkeyfile  string  }. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


