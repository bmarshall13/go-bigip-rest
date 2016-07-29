# LtmProfileSctp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeartbeatInterval** | **int64** | Specifies the number of seconds to wait before sending a heartbeat chunk. The default value is 30 seconds. | [optional] [default to 30]
**InStreams** | **int64** | Specifies the number of inbound streams. The default value is 2. | [optional] [default to 2]
**InitMaxRetries** | **int64** | Specifies the maximum number of retries to establish a connection. The default value is 4. | [optional] [default to 8]
**CookieExpiration** | **int64** | Specifies how many seconds the cookie is valid. The default value is 60 seconds. | [optional] [default to 60]
**LinkQos** | **string** | Specifies the link quality of service set in sent packets. The default value is 0. | [optional] [default to null]
**ReceiveChunks** | **int64** | Specifies the size (in chunks) of the rx_chunk buffer. The default value is 256. | [optional] [default to 256]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**TransmitChunks** | **int64** | Specifies the size (in chunks) of the tx_chunk buffer. The default value is 256. | [optional] [default to 256]
**ReceiveOrdered** | **string** | When enabled, the system delivers messages to the application layer in order. The default value is enabled. | [optional] [default to null]
**SendPartial** | **string** | When enabled, the default, the system accepts partial application data. | [optional] [default to null]
**Secret** | **string** | Specifies the internal secret string that the system uses for HTTP Message Authenticated Code (HMAC) cookies. | [optional] [default to null]
**HeartbeatMaxBurst** | **int64** | Specifies the maximum number of heartbeat packets sent in a single burst.The default value is 1. | [optional] [default to 1]
**MaxPathRetransmitLimit** | **int64** |  | [optional] [default to 5]
**MaxBurst** | **int64** | Specifies the maximum number of data packets sent in a single burst.The default value is 4. | [optional] [default to 4]
**ReceiveWindowSize** | **int64** | Specifies the size (in bytes) of the receive window. Prorate this value to the Receive Chunks value. The default value is 65536. | [optional] [default to null]
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**ClientSideMultihoming** | **string** |  | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**SendBufferSize** | **int64** | Specifies the size in bytes of the buffer. The default value is 65536. | [optional] [default to null]
**MaxCommunicationPaths** | **int64** |  | [optional] [default to 16]
**TcpShutdown** | **string** | When enabled, the default, the system emulates the closing of a TCP connection. | [optional] [default to null]
**ProxyBufferLow** | **int64** | Specifies the proxy buffer level after which the system opens the receive window. The default value is 4096. | [optional] [default to 4096]
**Name** | **string** | Name of entity | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**ServerSideMultihoming** | **string** |  | [optional] [default to null]
**IpTos** | **string** | Specifies the type of IP service set in packets sent to peer. The default value is 0. | [optional] [default to null]
**ProxyBufferHigh** | **int64** | Specifies the proxy buffer level after which the system closes the receive window. The default value is 16384. | [optional] [default to 16384]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**IdleTimeout** | **int64** | Specifies the number of seconds without traffic before a connection is eligible for deletion. The default value is 300 seconds. | [optional] [default to 300]
**ResetOnTimeout** | **string** | When enabled, the default, the system resets a connection when the connection times out. | [optional] [default to null]
**OutStreams** | **int64** | Specifies the number of outbound streams. The default value is 2. | [optional] [default to 2]
**SendMaxRetries** | **int64** | Specifies the maximum number of times the system tries again to send data. The default value is 8. | [optional] [default to 10]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


