# DnsCacheRecordsNameserver

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | **int64** | Chassis slot to query. | [optional] [default to 0]
**HasEdns** | **string** | Include or exclude nameservers which are EDNS lame. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Cache** | **string** | Name of DNS Cache configuration object to inspect. | [optional] [default to null]
**TtlRange** | **string** | Filter results by TTL, low:high. | [optional] [default to null]
**Address** | **string** | Filter results by nameserver address. | [optional] [default to null]
**RttRange** | **string** | Filter results by RTT, low:high. | [optional] [default to null]
**HasLame** | **string** | Include or exclude records which have lame zones. | [optional] [default to null]
**Tmm** | **int64** | TMM to query (typically unset and used for debugging). | [optional] [default to -1]
**ZoneName** | **string** | Filter results by zone name. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


