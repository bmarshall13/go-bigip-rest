# ProfileWebAcceleration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**CacheMaxEntries** | **int64** | Specifies the maximum number of entries that can be in the cache. The default value is 0 (zero), which means that the system does not limit the maximum entries. | [optional] [default to null]
**CacheObjectMaxSize** | **int64** | Specifies the largest object that the system considers eligible for caching. The default value is 50000 bytes. | [optional] [default to null]
**CacheAgingRate** | **int64** | Specifies how quickly the system ages a cache entry. The aging rate ranges from 0 (slowest aging) to 10 (fastest aging). The default value is 9. | [optional] [default to null]
**MetadataCacheMaxSize** | **int64** | Specifies the maximum size, in megabytes, of the metadata cache. This does not include the content cache, and only applies when a WebAccelerator application is on the profile. | [optional] [default to 25]
**CacheMaxAge** | **int64** | Specifies how long the system considers the cached content to be valid. The default value is 3600 seconds. | [optional] [default to null]
**CacheUriExclude** | **string** | Configures a list of URIs to exclude from the cache. The default value of none specifies no URIs are excluded. | [optional] [default to null]
**Applications** | **string** | Configures a list of applications assigned to this profile. Assigning at least one application enables WA functionality. The default value of none specifies that WA is not enabled. | [optional] [default to null]
**CacheUriPinned** | **string** | Configures a list of URIs to keep in the cache. The pinning process keeps URIs in cache when they would normally be evicted to make room for more active URIs. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**CacheClientCacheControlMode** | **string** | Specifies which cache disabling headers sent by clients the system ignores. The default value is all. | [optional] [default to null]
**CacheObjectMinSize** | **int64** | Specifies the smallest object that the system considers eligible for caching. The default value is 500 bytes. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**CacheUriIncludeOverride** | **string** | Configures a list of URIs to include in the cache even if they would normally be excluded due to factors like object size or HTTP request type. The default value of none specifies no URIs are to be forced into the cache. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the profile resides. | [optional] [default to null]
**CacheInsertAgeHeader** | **string** | Inserts Age and Date headers in the response. The default value is enabled. | [optional] [default to null]
**CacheUriInclude** | **string** | Configures a list of URIs to include in the cache. The default value of .* specifies that all URIs are cacheable. | [optional] [default to null]
**CacheSize** | **int64** | Specifies the maximum size for the cache. When the cache reaches the maximum size, the system starts removing the oldest entries. The default value is 100 megabytes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


