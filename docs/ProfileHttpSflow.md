# ProfileHttpSflow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamplingRate** | **int64** | Specifies the ratio of packets observed to the samples generated. For example, a sampling rate of 2000 specifies that 1 sample will be randomly generated for every 2000 packets observed. To enable this setting, you must also set the sampling-rate-global setting to no. | [optional] [default to 0]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**SamplingRateGlobal** | **string** | Specifies whether the global HTTP sampling-rate setting overrides the object-level sampling-rate setting. The default value is yes. | [optional] [default to null]
**PollInterval** | **int64** | Specifies the maximum interval in seconds between two pollings. To enable this setting, you must also set the poll-interval-global setting to no. | [optional] [default to 0]
**PollIntervalGlobal** | **string** | Specifies whether the global HTTP poll-interval setting overrides the object-level poll-interval setting. The default value is yes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


