# PersistenceHash

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**HashAlgorithm** | **string** | Specifies the algorithm the system uses for hash persistence load balancing. The hash result is the input for the algorithm. The default value is default. carp specifies to use the Cache Array Routing Protocol (CARP) to obtain the hash result for the input to the algorithm. default specifies to use the index of pool members to obtain the hash result for the input to the algorithm. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**MatchAcrossServices** | **string** | Specifies, when enabled, that all persistent connections from a client IP address, which go to the same virtual IP address, also go to the same node. The default value is disabled. | [optional] [default to null]
**HashBufferLimit** | **int64** | Specifies the maximum buffer length the system collects to locate the hashing pattern for hash persistence load balancing. The default value is 0. | [optional] [default to 0]
**DefaultsFrom** | **string** | Specifies the existing profile from which the system imports settings for the new profile. | [optional] [default to null]
**Mirror** | **string** | Specifies whether the system mirrors persistence records to the high-availability peer. The default value is disabled. | [optional] [default to null]
**OverrideConnectionLimit** | **string** | Specifies, when enabled, that the pool member connection limits are not enforced for persisted clients. Per-virtual connection limits remain hard limits and are not disabled. The default value is disabled. | [optional] [default to null]
**MatchAcrossPools** | **string** | Specifies, when enabled, that the system can use any pool that contains this persistence record. The default value is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the profile resides. | [optional] [default to null]
**Rule** | **string** | Specifies a rule name if you are using a rule for universal persistence. | [optional] [default to null]
**HashLength** | **int64** | Specifies the length of data within the packet in bytes that the system uses to calculate the hash value when performing hash persistence load balancing. The default value is 0 bytes. | [optional] [default to 0]
**MatchAcrossVirtuals** | **string** | Specifies, when enabled, that all persistent connections from the same client IP address go to the same node. The default value is disabled. | [optional] [default to null]
**Mode** | **string** |  | [optional] [default to null]
**Timeout** | **string** | Specifies the duration of the persistence entries. The default value is 180 seconds. | [optional] [default to null]
**HashStartPattern** | **string** | Specifies the string that describes the start location of the hash pattern that the system uses to perform hash persistence load balancing. The default value is none. | [optional] [default to null]
**HashOffset** | **int64** | Specifies the start offset within the packet from which the system begins the hash when performing hash persistence load balancing. The default value is 0. | [optional] [default to 0]
**HashEndPattern** | **string** | Specifies the string that describes the ending location of the hash pattern that the system uses to perform hash persistence load balancing. The default value is none. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


