# SysSflowReceiver

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this receiver resides. | [optional] [default to null]
**State** | **string** | Specifies the state of the receiver. The sFlow samples will be collected and sent to the receiver when enabled. The default value is disabled. | [optional] [default to null]
**MaxDatagramSize** | **int64** | Specifies the maximum size of sFlow datagram. The default value is 1400. | [optional] [default to 1400]
**Address** | **string** | Specifies the IP address of the receiver. | [default to null]
**Port** | **int64** | Specifies the IP port of the receiver. The default value is 6343. | [optional] [default to 6343]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


