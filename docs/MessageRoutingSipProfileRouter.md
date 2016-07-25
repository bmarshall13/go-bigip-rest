# MessageRoutingSipProfileRouter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**MaxPendingBytes** | **int64** | Specifies the maximum number of bytes contained within pending messages that will be held while waiting for a connection to a peer to be created. | [optional] [default to 23768]
**MaxPendingMessages** | **int64** | Specifies the maximum number of pending messages that will be held while waiting for a connection to a peer to be created. | [optional] [default to 64]
**Partition** | **string** |  | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**UseLocalConnection** | **string** | Enables or disables to specify whether the system, when selecting an outgoing connection to a destination peer, prefers the connection established by the incoming request, rather than other connections. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all of the settings and values from the specified parent profile. | [optional] [default to null]
**OperationMode** | **string** | Specifies the behavior of the routing instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


