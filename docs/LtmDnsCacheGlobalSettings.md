# LtmDnsCacheGlobalSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheMaximumTtl** | **int64** | Force resolver to re-query for resource records at cache-maximum-ttl seconds if its original ttl is greater than cache-maximum-ttl. | [optional] [default to 864000]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**ResolverEdnsBufferSize** | **int64** | Buffer size resolver advertises in udp queries. | [optional] [default to 4096]
**CacheMinimumTtl** | **int64** | Cache resource records for at least cache-minimum-ttl seconds, i.e. longer than owner intended. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


