# LtmMessageRoutingDiameterProfileRouter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**Description** | **string** | Specifies descriptive text that identifies the profile. | [optional] [default to null]
**MaxPendingBytes** | **int64** | Specifies the maximum number of bytes contained within pending messages that will be held while waiting for a connection to a peer to be created. If the specified value is reached, any additional messages to the peer will be undeliverable, and held messages are delivered to the peer. The default value is 0, which disables this setting. | [optional] [default to null]
**MaxPendingMessages** | **int64** | Specifies the maximum number of pending messages held while waiting for a connection to a peer to be created. If the specified value is reached, any additional messages to the peer will be undeliverable, and held messages are delivered to the peer. The default value is 0, which disables this setting. | [optional] [default to null]
**TransactionTimeout** | **int64** | Specifies the maximum period in seconds between a request and response. A provisional response restarts the timer. This setting might not affect all transactions. The default value is 10 seconds. | [optional] [default to 10]
**UseLocalConnection** | **string** | When selected (enabled), controls whether connections that are established by the ingress TMM are preferred to connections that are established by another TMM when selecting an egress connection to a destination peer. The default value is enabled. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile from which this profile inherits settings. The default is the system-supplied \&quot;diameterrouter\&quot; profile. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


