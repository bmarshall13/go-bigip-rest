# MessageRoutingGenericProtocol

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**DisableParser** | **string** | When set, the generic message profile parser will be disabled. It will ignore all incoming packets and not directly send message data. This mode supports iRule script protocol implementations that will generate messages from the incoming transport stream and send outgoing messages on the outgoing transport stream. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**MessageTerminator** | **string** | The string of characters used to terminate a message. If the message-terminator is empty, the generic message parser will not separate the input stream into messages. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**MaxMessageSize** | **int64** | Specifies the maximum size of a received message. If a message exceeds this size, the connection will be reset. | [optional] [default to 32768]
**NoResponse** | **string** | When set, matching of responses to requests is disabled. | [optional] [default to null]
**MaxEgressBuffer** | **int64** | Specifies the maximum size of the send buffer in bytes. If the number of bytes in the send buffer for a connection exceeds this value, the generic message protocol will stop receiving outgoing messages from the router until the size of the size of the buffer drops below this setting. | [optional] [default to 32768]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all of the settings and values from the specified parent profile. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


