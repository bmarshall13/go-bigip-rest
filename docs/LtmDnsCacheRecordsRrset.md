# LtmDnsCacheRecordsRrset

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | **int64** | Chassis slot to query. | [optional] [default to 0]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Cache** | **string** | Name of DNS Cache configuration object to inspect. | [optional] [default to null]
**TmClass** | **string** | Record class to filter on. | [optional] [default to null]
**TtlRange** | **string** | Range of TTLs to filter on, low:high. | [optional] [default to null]
**Owner** | **string** | DNS owner name for filtering RRs. | [optional] [default to null]
**Type_** | **string** | Record type to filter on. | [optional] [default to null]
**Tmm** | **int64** | TMM to query (typically unset and used for debugging). | [optional] [default to -1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


