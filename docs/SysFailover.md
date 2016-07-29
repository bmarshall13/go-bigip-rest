# SysFailover

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**NoPersist** | **bool** | Does not persist the change in status of a unit or cluster. Valid only with offline status. | [optional] [default to null]
**Standby** | **bool** | Specifies that the active unit or cluster fails over to a Standby state, causing the standby unit or cluster to become Active. | [optional] [default to null]
**Online** | **bool** | Changes the status of a unit or cluster from Forced Offline to either Active or Standby, depending upon the status of the other unit or cluster in a redundant pair. | [optional] [default to null]
**TrafficGroup** | **string** | Specifies the name of the traffic-group that the standby command refers to. | [optional] [default to null]
**Persist** | **bool** | Persists the change in status of a unit or cluster. Valid only with offline status. | [optional] [default to null]
**Device** | **string** | Specifies the name of the device that should become the active device for the traffic group or for all traffic groups. | [optional] [default to null]
**Offline** | **bool** | Changes the status of a unit or cluster to Forced Offline. If persist or no-persist are not specified, the change in status will be persisted in-between system restarts. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


