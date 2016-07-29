# LtmMessageRoutingGenericPeer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Ratio** | **int64** | Specifies the ratio to be used for selection of a peer within a list of peers in a ltm genericmsg-route. | [optional] [default to 1]
**Description** | **string** | User defined description. | [optional] [default to null]
**ConnectionMode** | **string** | Specifies how the number of connections per host are to be limited. Note a host (specified in the referred pool) may also exist in other peer objects and those other peer objects may have different settings for connection-mode and number_connections. Thus these settings specify how messages routed through this peer are distributed between a set of connections, not the maximum number of connections to a specified host. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**TransportConfig** | **string** | The name of the ltm virtual or ltm transport-config to use for creating an outgoing connection. | [optional] [default to null]
**NumberConnections** | **int64** | Specifies the distribution of connections between the BIG-IP and a remote host. | [optional] [default to 1]
**Pool** | **string** | Specifies the name of the pool that messages will be routed towards. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


