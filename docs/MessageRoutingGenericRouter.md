# MessageRoutingGenericRouter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**MaxPendingBytes** | **int64** | The maximum number of bytes worth of pending messages that will be held while waiting for a connection to a peer to be created. Once reached, any additional messages to the peer will be flagged as undeliverable and returned to the originator. | [optional] [default to 23768]
**MaxPendingMessages** | **int64** | The maximum number of pending messages that will be held while waiting for a connection to a peer to be created. Once reached, any additional messages to the peer will be flagged as undeliverable and returned to the originator. | [optional] [default to 64]
**Partition** | **string** |  | [optional] [default to null]
**TrafficGroup** | **string** | Specifies the traffic group of the router. The default is inherited from the containing folder. | [optional] [default to null]
**UseLocalConnection** | **string** | If true, the router will route a message to an existing connection on the same TMM as the message was received on. If an existing connection is not found, it will route the message through an existing connection based on a deterministic algorithm that may be on another TMM. If a matching existing connection is not found, it will create a connection on the current TMM. Setting this flag may limit the number of connections that are created to a peer. | [optional] [default to null]
**IgnoreClientPort** | **string** | If set, the remote port on clientside connections (connections where the peer connected to the BIG IP) is ignored when searching for an existing connection. | [optional] [default to null]
**InheritedTrafficGroup** | **string** |  | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all of the settings and values from the specified parent profile. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


