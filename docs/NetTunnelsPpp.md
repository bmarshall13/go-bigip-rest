# NetTunnelsPpp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**LcpEchoFailure** | **int64** | Number of consecutive PPP LCP echo messages that must go unanswered for the server to drop PPP connection. For example, if the server sends  number  of consecutive PPP LCP Echo Request messages that go unanswered (by Echo Reply), it will close the PPP connection. | [optional] [default to 4]
**Partition** | **string** | Displays the admin-partition within which this component resides. | [optional] [default to null]
**Ipcp** | **string** |  | [optional] [default to null]
**DefaultsFrom** | **string** | The profile from which to inherit settings. The default value is ppp. | [optional] [default to null]
**Ipv6cp** | **string** |  | [optional] [default to null]
**Vj** | **string** | VJ is a data compression protocol described in RFC 1144, specifically designed by Van Jacobson to improve TCP/IP performance over slow serial links. Van Jacobson Header Compression (also known as VJ compression, or just Header Compression) is an option in most versions of PPP. | [optional] [default to null]
**LcpEchoInterval** | **int64** | Specifies interval in seconds between PPP LCP Echo Request messages that the server sends to the peer (client). | [optional] [default to 30]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


