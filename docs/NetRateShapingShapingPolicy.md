# NetRateShapingShapingPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**RatePercentage** | **float32** | Specifies the percentage of the maximum throughput rate specified for the parent in the rate class command that is available for this traffic flow. | [optional] [default to null]
**CeilingPercentage** | **float32** | Specifies the percentage of the ceiling rate specified for the parent in the rate class command that is available for this traffic flow. | [optional] [default to null]
**Queue** | **string** | Specifies the queue discipline for this traffic flow. The available pre-configured settings are pfifo (Priority First in, First out), sfq (Stochastic Fair Queuing), and none. The default value is none. You can create a customized queue discipline using the queue command. | [optional] [default to null]
**DropPolicy** | **string** | Specifies the drop policy for this traffic flow. The available pre-configured settings are tail (drops the end of the traffic stream), red (randomly drops packets), and fred (drops packets according to the type of traffic in the flow). The default value is none. You can create a customized drop policy using the drop-policy command. | [optional] [default to null]
**MaxBurst** | **int64** | Specifies the maximum number of bytes that traffic is allowed to burst beyond the base rate. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


