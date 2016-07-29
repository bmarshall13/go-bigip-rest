# SysCryptoClient

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryInterval** | **int64** | Specifies the interval in seconds between attempts to connect to the remote crypto server. The default value is 10 seconds. | [optional] [default to 10]
**ConnectionReset** | **bool** | Resets the connection to the remote crypto server. | [optional] [default to null]
**Addr** | **string** | Specifies the IP address of the remote crypto server. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**ReqTimeout** | **int64** | Specifies the timeout in milliseconds for crypto requests to complete. The default value is 5000 milliseconds. | [optional] [default to 5000]
**Heartbeat** | **int64** | Specifies the number of seconds to wait before sending a heartbeat request. A value of 0 disables the sending of heartbeat requests. The default value is 30 seconds. | [optional] [default to 30]
**MaxRetries** | **string** | Specifies the maximum number of times to retry connecting to the remote crypto server. If the maximum retries value is infinite, the crypto client retries connecting until a connection is made. The default value is infinite. | [optional] [default to null]
**Port** | **int64** | Specifies the port used by the remote crypto server. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


