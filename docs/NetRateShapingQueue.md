# NetRateShapingQueue

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**SfqPerturbation** | **int64** | Specifies the interval in seconds at which the system reconfigures the SFQ hash function. This option applies only to the sfq type. | [optional] [default to 10]
**Description** | **string** | User defined description. | [optional] [default to null]
**SfqBucketSize** | **int64** | Specifies the bucket size in kilobytes or megabytes for the sfq type. The default value is 0 (zero). | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]
**PfifoMaxSize** | **int64** | Specifies the size in kilobytes or megabytes of the largest queue for the pfifo type only. The default value is 0 (zero). | [optional] [default to 0]
**Type_** | **string** | Specifies the queue discipline this custom queue uses. The available values are none, sfq, and pfifo. See the help page for a more detailed description. | [optional] [default to null]
**PfifoMinSize** | **int64** | Specifies the size in kilobytes or megabytes of the smallest queue for the pfifo type only. The default value is 0 (zero). | [optional] [default to 0]
**SfqBucketCount** | **int64** | Specifies the number of buckets in kilobytes or megabytes into which the queue is divided when you are configuring the sfq type. Valid values are 0, 16, 32, 64, 128, 256, 512, and 1024. The default value is 0 (zero). | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


