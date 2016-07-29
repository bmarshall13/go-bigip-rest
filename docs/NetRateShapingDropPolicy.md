# NetRateShapingDropPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**MinThreshold** | **int64** | Specifies the queue length above which packets are not dropped. The default value is 0. | [optional] [default to 0]
**AveragePacketSize** | **int64** | Specifies the average MTU (maximum transmission unit) size in the range of 0 to 10000 bytes. The default value is 0. | [optional] [default to 0]
**FredMaxDrop** | **int64** | Specifies the hard drop limit in the range of 0 to 400. The default value is 0. Setting this to a small value does not change the hard drop limit, but a higher number increases the limit. | [optional] [default to 0]
**MaxProbability** | **int64** | Specifies the maximum percentage probability in the range of 0 to 100 according to which packets are dropped when the average queue length is between the minimum and maximum thresholds. The default value is 0. | [optional] [default to 0]
**RedHardLimit** | **int64** | Specifies the maximum queue size in kilobytes or megabytes. Additional packets are dropped. The default value is 0. This option applies only to the red type. | [optional] [default to 0]
**FredMaxActive** | **int64** | Specifies the maximum number of flows that can be active for each queue. The range is 0 to 10000. The default value is 0, which disables active flow limitation. | [optional] [default to 0]
**InverseWeight** | **int64** | Specifies the weight used to calculate the average queue length. Valid values are 0, 64, 128, 256, 512, and 1024. The default value is 0. | [optional] [default to 0]
**MaxThreshold** | **int64** | Specifies the queue length below which packets are not dropped. The default value is 0. | [optional] [default to 0]
**Type_** | **string** | Specifies the type of drop policy. The available settings are tail (drops the end of the traffic stream), red (randomly drops packets), and fred (drops packets according to the type of traffic in the flow). The default value is red. Although you could create a drop policy based on tail, that is already the default value for drop policy in both the shaping policy and rate class commands. | [optional] [default to null]
**FredMinDrop** | **int64** | Specifies the hard no drop limit in the range of 0 to 100. The default value is 0. Setting this to a large value prevents packets from being dropped. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


