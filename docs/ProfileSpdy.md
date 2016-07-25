# ProfileSpdy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**ProtocolVersions** | **string** | Specifies which protocol versions should enabled. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**CompressionWindowSize** | **int64** | Specifies in KB the size of the compression window. | [optional] [default to 8]
**ActivationMode** | **string** | Specifies what will cause an incoming connection to be handled as a SPDY connection. The default value &#39;tls&#39; specifies that both the TLS next-protocol-negotiation and the application-layer-protocol-negotiationextensions will be used. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**PriorityHandling** | **string** | Specifies how concurrent streams of different priorities should be handled. | [optional] [default to null]
**ConcurrentStreamsPerConnection** | **int64** | Specifies how many concurrent requests are allowed to be outstanding on a single SPDY connection. | [optional] [default to 10]
**CompressionLevel** | **int64** | Specifies the amount of compression of headers to do by default. Higher numbers correspond to more compression. | [optional] [default to 5]
**ConnectionIdleTimeout** | **int64** | Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. | [optional] [default to 300]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**InsertHeader** | **string** | Specifies whether an HTTP header should be added to the HTTP request to show the request was received via SPDY. | [optional] [default to null]
**InsertHeaderName** | **string** | Specifies the name of the header that is added to the HTTP request when insert-header is enabled. | [optional] [default to null]
**FrameSize** | **int64** | The size in bytes of the data frames that will be produced by SPDY. | [optional] [default to 2048]
**WriteSize** | **int64** | The size in bytes of the SSL records that will be produced by SPDY. | [optional] [default to 16384]
**ReceiveWindow** | **int64** | Specifies in KB the size of the receive window for SPDY flow-control. | [optional] [default to 32]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


