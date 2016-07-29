# SysApplicationCustomStat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Keyspace** | **string** | The type of object to which this derivation applies.  e.g. \&quot;ltm.pool\&quot; or \&quot;net.interface\&quot;. | [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Measure** | **string** | The output measure name. | [default to null]
**Formula** | **string** | Quoted string indicating the derivation.  e.g. \&quot;rate counter source_measure 60\&quot; computes the 60-second rate of the source_measure counter. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


