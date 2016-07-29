# NetTunnelsFec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**DecodeIdleTimeout** | **int64** | Specifies the maximum waiting time for packets in decoding queues. | [optional] [default to 1500]
**Description** | **string** | User defined description. | [optional] [default to null]
**SourceAdaptive** | **string** | Controls the possibility of using the adaptive FEC source packets parameter. | [optional] [default to null]
**DecodeQueues** | **int64** | Specifies the number of decoding queues. | [optional] [default to 32]
**KeepaliveInterval** | **int64** | Specifies the interval between keepalive (statistical data) packets. | [optional] [default to 5]
**SourcePackets** | **int64** | Specifies that the source packets parameter is used to divide the aggregated payload. | [optional] [default to 15]
**Lzo** | **string** | Controls use of the LZO compression algorithm for compressing data packets. | [optional] [default to null]
**EncodeMaxDelay** | **int64** | Specifies the maximum time for packet aggregation. | [optional] [default to 500]
**DefaultsFrom** | **string** | The profile from which to inherit settings. The default value is fec. | [optional] [default to null]
**RepairPackets** | **int64** | Specifies that the repair packets parameter is used for adding redundant packets. | [optional] [default to 15]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**UdpPort** | **int64** | Specifies the local port for receiving FEC packets. | [optional] [default to 8288]
**Name** | **string** | Name of entity | [optional] [default to null]
**Partition** | **string** | Displays the admin-partition within which this component resides. | [optional] [default to null]
**DecodeMaxPackets** | **int64** | Specifies the maximum number of waiting packets in decoding queues. | [optional] [default to 512]
**RepairAdaptive** | **string** | Controls the possibility of using the FEC adaptive algorithm to modify the repaired packets parameter to real network conditions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


