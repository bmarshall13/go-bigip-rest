# ProfileHttp2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ActivationModes** | **string** | Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used. | [optional] [default to null]
**ConcurrentStreamsPerConnection** | **int64** | Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection. | [optional] [default to 10]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**EnforceTlsRequirements** | **string** |  | [optional] [default to null]
**HeaderTableSize** | **int64** | Specifies what table size will be used for the compression of headers (unused). | [optional] [default to 4096]
**FrameSize** | **int64** | The size in bytes of the data frames that will be produced by HTTP/2. | [optional] [default to 2048]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**ReceiveWindow** | **int64** | Specifies in KB the size of the receive window for HTTP/2 flow-control. | [optional] [default to 32]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**InsertHeader** | **string** | Specifies whether an HTTP header should be added to the HTTP request to show the request was received via HTTP/2. | [optional] [default to null]
**InsertHeaderName** | **string** | Specifies the name of the header that is added to the HTTP request when insert-header is enabled. | [optional] [default to null]
**WriteSize** | **int64** | The size in bytes of the SSL records that will be produced by HTTP/2.handled. | [optional] [default to 16384]
**ConnectionIdleTimeout** | **int64** | Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. | [optional] [default to 300]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


