# LtmProfileFasthttp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**ConnpoolMaxReuse** | **int64** | Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero). | [optional] [default to 0]
**MaxRequests** | **int64** | Specifies the maximum number of requests that the system can receive on a client-side connection, before the system closes the connection. A setting of 0 specifies that requests are not limited. The default value is 0 (zero). | [optional] [default to 0]
**HeaderInsert** | **string** | Specifies a string that the system inserts as a header in an HTTP request. If the header exists already, the system does not replace it. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ConnpoolReplenish** | **string** | The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option. | [optional] [default to null]
**ConnpoolMinSize** | **int64** | Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10. | [optional] [default to 4]
**InsertXforwardedFor** | **string** | Specifies whether the system inserts the XForwarded For header in an HTTP request with the client IP address, to use with connection pooling. Enabling this specifies that the system inserts the XForwarded For header with the client IP address. Disabling this specifies that the system does not insert the XForwarded For header. | [optional] [default to null]
**ConnpoolStep** | **int64** | Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4. | [optional] [default to 4]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**ConnpoolIdleTimeoutOverride** | **int64** | Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting. | [optional] [default to 0]
**ServerSack** | **string** | Specifies whether to support server sack option in cookie response by default. The default is disabled. | [optional] [default to null]
**ClientCloseTimeout** | **int64** | Specifies the number of seconds after which the system closes a client connection, when the system either receives a client FIN packet or sends a FIN packet. This setting overrides the idle timeout setting. The default value is 5. | [optional] [default to 5]
**ServerTimestamp** | **string** | Specifies whether to support server timestamp option in cookie response by default. The default is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**ForceHttp10Response** | **string** | Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**ConnpoolMaxSize** | **int64** | Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048. | [optional] [default to 2048]
**HardwareSynCookie** | **string** | Specifies whether or not to use hardware SYN Cookie when cross system limit. The default is disabled. | [optional] [default to null]
**UncleanShutdown** | **string** | Specifies how the system handles closing a connection. The default value is enabled, which allows unclean shutdown of a client connection. Use disabled to prevent unclean shutdown of a client connection. Fast specifies that the system sends a RESET packet to close the connection only if the client attempts to send further data after the response has completed. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**IdleTimeout** | **int64** | Specifies the number of seconds after which a connection is eligible for deletion, when the connection has no traffic. The default value is 300 seconds. | [optional] [default to 300]
**Http11CloseWorkarounds** | **string** | Enables or disables HTTP 1.1 close workarounds. The default value is disabled. | [optional] [default to null]
**Layer7** | **string** | When enabled, the system parses HTTP data in the stream. Disable this setting if you want to use the performance HTTP profile to shield against denial-of-service attacks against non-HTTP protocols. The default value is enabled. | [optional] [default to null]
**ResetOnTimeout** | **string** | When enabled, the system sends a TCP RESET packet when a connection times out, and deletes the connection. The default value is enabled. | [optional] [default to null]
**ReceiveWindowSize** | **int64** | Specifies the size of the receive window, in bytes. The minimum and default value is 65535 bytes without scale. | [optional] [default to null]
**ServerCloseTimeout** | **int64** | Specifies the number of seconds after which the system closes a client connection, when the system either receives a client FIN packet or send a FIN packet. This setting overrides the idle timeout setting. The default value is 5. | [optional] [default to 5]
**MssOverride** | **int64** | Specifies a maximum segment size (MSS) override for server-side connections. The default value is 0 (zero), which corresponds to an MSS of 1460. You can specify any integer between 536 and 1460. | [optional] [default to 0]
**MaxHeaderSize** | **int64** | Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768. | [optional] [default to 32768]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


