# ProfileRtsp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**MulticastRedirect** | **string** | When enabled, the client can select the destination to which to stream data. The default value is disabled. | [optional] [default to null]
**IdleTimeout** | **string** | Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds. | [optional] [default to null]
**RtcpPort** | **int64** | The Real Time Control Protocol (RTCP) allows monitoring of the real-time data delivery. Specifies the number of the port to use for the RTCP service. | [optional] [default to 0]
**CheckSource** | **string** | When enabled the system uses the source attribute in the transport header to establish the target address of the RTP stream, and before the response is forwarded to the client, updates the value of the source attribute to be the virtual address of the BIG-IP system. When disabled the system does not change the source attribute. The default value is enabled. | [optional] [default to null]
**Proxy** | **string** | Specifies whether the RTSP filter is associated with an RTSP proxy configuration. The default value is none. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**RtpPort** | **int64** | The Real Time Protocol (RTP) provides data transport functions suitable for applications transmitting real-time data. Specifies the number of the port to use for the RTP service. | [optional] [default to 0]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**UnicastRedirect** | **string** | When enabled, the client can select the destination to which to stream data. The default value is disabled. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**LogPublisher** | **string** | Configures the log publisher that handles events logging for this profile. | [optional] [default to null]
**SessionReconnect** | **string** | When enabled, the RTSP filter persists the control connection, which is being resumed, to the correct server. The default value is disabled. | [optional] [default to null]
**ProxyHeader** | **string** | When the proxy option is set, specifies the name of the header in the RTSP proxy configuration that is passed from the client-side virtual server to the server-side virtual server. Note that the name of the header must begin with X-. | [optional] [default to null]
**MaxQueuedData** | **int64** | Specifies the maximum amount of data that the RTSP filter buffers before dropping the connection. The default value is 32768 bytes. | [optional] [default to 32768]
**LogProfile** | **string** | Configures the ALG log profile that controls logging. | [optional] [default to null]
**RealHttpPersistence** | **string** | When enabled, the RTSP filter automatically persists Real Networks RTSP over HTTP using the RTSP port. The default value is enabled. If you disable this parameter, you can override the default behavior with an iRule. | [optional] [default to null]
**MaxHeaderSize** | **int64** | Specifies the maximum size of an RTSP request or response header that the RTSP filter allows before dropping the connection. The default value is 4096 bytes. | [optional] [default to 4096]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


