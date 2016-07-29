# SysDatastor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**LowWaterMark** | **int64** | Specifies the percentage of full cache below which pruning stops. | [optional] [default to 80]
**Description** | **string** | User defined description. | [optional] [default to null]
**HighWaterMark** | **int64** | Specifies the percentage of full cache above which pruning starts. | [optional] [default to 90]
**StoreSize** | **int64** | Displays the amount of space for each disk path specified. | [optional] [default to 0]
**WebCacheWeight** | **int64** | Specifies the relative weight of the web cache. | [optional] [default to 10]
**CacheSize** | **int64** | Displays the size of the data storage in megabytes (MB). | [optional] [default to 0]
**Disk** | **string** | Enables or disables the use of the disk for data storage. | [optional] [default to null]
**DedupCacheWeight** | **int64** | Specifies the relative weight of the dedup cache. | [optional] [default to 10]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


