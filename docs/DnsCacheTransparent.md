# DnsCacheTransparent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Idx** | **float32** |  | [optional] [default to null]
**MsgCacheSize** | **int64** | Number of bytes allocated for the message cache. The default value is 1m | [optional] [default to 1048576]
**Partition** | **string** |  | [optional] [default to null]
**AnswerDefaultZones** | **string** | Answer queries for default zones: localhost, reverse 127.0.0.1 and ::1, and AS112 zones. The default value is no. | [optional] [default to null]
**LocalZones** | **string** | Zones and associated resource records for which the cache will provide Authoritative answers. | [optional] [default to null]
**RrsetCacheSize** | **int64** | Number of bytes allocated for the resource record set cache. The default value is 10m. | [optional] [default to 10485760]
**Type_** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


