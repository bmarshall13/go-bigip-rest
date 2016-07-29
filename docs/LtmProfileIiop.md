# LtmProfileIiop

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**AbortOnTimeout** | **string** | Specifies whether the system aborts the connection if there is no response received within the time specified in the timeout option. The default value is disabled. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**Timeout** | **int64** | Specifies the request timeout. The system uses this value when the abort-on-timeout option is enabled. The default value is 30 seconds. | [optional] [default to 30]
**PersistObjectKey** | **string** | Specifies whether to persist connections based on the object key in the IIOP request. The default value is disabled. | [optional] [default to null]
**PersistRequestId** | **string** | Specifies whether to persist connections based on the request ID in the IIOP request. The default value is enabled. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


