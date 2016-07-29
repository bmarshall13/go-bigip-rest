# LtmPersistenceGlobalSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestAddrLimitMode** | **string** | Specifies that the persistence session is limited by either the number of seconds before the persistence entry times out, or by a maximum number of requests to the destination address. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**ProxyGroup** | **string** | Specifies a group of servers that are configured to process all of the requests from a single source address during a persistence session. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**DestAddrMax** | **int64** | Specifies the maximum number of entries that can be in the persistence table at any one time when using the destination address affinity mode and when the option dest addr limit is set to max-count. The default value is 2048 entries. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


