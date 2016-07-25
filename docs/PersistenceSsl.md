# PersistenceSsl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**MatchAcrossPools** | **string** | Specifies, when enabled, that the system can use any pool that contains this persistence record. The default value is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**MatchAcrossServices** | **string** | Specifies, when enabled, that all persistent connections from a client IP address, which go to the same virtual IP address, also go to the same node. The default value is disabled. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**OverrideConnectionLimit** | **string** | Specifies, when enabled, that the pool member connection limits are not enforced for persisted clients. Per-virtual connection limits remain hard limits and are not disabled. The default value is disabled. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the profile resides. | [optional] [default to null]
**MatchAcrossVirtuals** | **string** | Specifies, when enabled, that all persistent connections from the same client IP address go to the same node. The default value is disabled. | [optional] [default to null]
**Mode** | **string** |  | [optional] [default to null]
**Timeout** | **string** | Specifies the duration of the persistence entries. The default value is 180 seconds. | [optional] [default to null]
**Mirror** | **string** | Specifies whether the system mirrors persistence records to the high-availability peer. The default value is disabled. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing profile from which the system imports settings for the new profile. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


