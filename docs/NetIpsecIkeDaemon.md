# NetIpsecIkeDaemon

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**LogLevel** | **string** | Specifies the logging level of the IKE daemon. The log file is located at /var/log/racoon.log. | [optional] [default to null]
**LogPublisher** | **string** | Specifies the logging publisher. A new log-publisher object can be created via TMSH command tmsh create sys log-config publisher. | [optional] [default to null]
**IsakmpPort** | **int64** | Specifies the port for the IKE daemon to accept ISAKMP messages. Only 500 is currently supported. | [optional] [default to 500]
**IsakmpNattPort** | **int64** | Specifies the port for the IKE daemon to accept ISAKMP messages when NAT-Traversal is detected. This is also the port number to accept UDP-encapsulated ESP traffic for NAT-Traversal. Only 4500 is currently supported. | [optional] [default to 4500]
**NattKeepAlive** | **int64** | Specifies the interval between sending NAT-Traversal keep-alive packets. The default value is 20 seconds. Set to 0 to disable keep-alive packets. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


