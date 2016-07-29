# SysCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinUpMembers** | **int64** | Specifies the minimum number of cluster members that must be up for the cluster to remain Active. The default value is 0 (zero). | [optional] [default to 0]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**MinUpMembersEnabled** | **string** | When enabled, specifies that when the number of active cluster members is below the value of the min up members option, the cluster fails over to its peer. Enable this parameter when you configure a redundant pair. The default value is disabled. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Address** | **string** | Specifies the IP address of the cluster. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


