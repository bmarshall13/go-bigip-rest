# SysStateMirroring

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | **string** | Specifies the primary self-IP address on this unit to which the peer unit in this redundant pair mirrors its connections.  This is a required setting. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**PeerAddr** | **string** | Specifies the primary self-IP address on the peer unit to which this unit mirrors its connections. This is a required setting. | [optional] [default to null]
**SecondaryAddr** | **string** | Specifies another self-IP address on this unit to which the peer unit mirrors its connections when the primary address is unavail able. | [optional] [default to null]
**SecondaryPeerAddr** | **string** | Specifies another self-IP address on the peer unit to which this unit mirrors its connections when the primary peer address is unavailable. | [optional] [default to null]
**State** | **string** | Enables or disables connection mirroring. The default is enabled. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


