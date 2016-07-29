# LtmEvictionPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**HighWater** | **int64** | The target maximum load as a percentage (of the full capacity available to the context assigned). | [optional] [default to 95]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**LowWater** | **int64** | The percentage (of the full capacity available to the context assigned) at which aggressive killing of flows is enabled or disabled. | [optional] [default to 85]
**Partition** | **string** |  | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


