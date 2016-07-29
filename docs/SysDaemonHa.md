# SysDaemonHa

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeartbeatAction** | **string** | Specifies the action to take if the daemon does not maintain its heartbeat. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Running** | **string** | Specifies whether the running-timeout and non-running-action properties are enabled. | [optional] [default to null]
**NotRunningAction** | **string** | Specifies that action that will be taken if the daemon is not running. This is a read only property. | [optional] [default to null]
**RunningTimeout** | **int64** | Specifies the amount of time (in seconds) that must elapse before the daemon is considered to be not running. This is a read only property. | [optional] [default to 10]
**Heartbeat** | **string** | Specifies whether heartbeat monitoring is enabled for the daemon. If monitoring is enabled and the daemon does not maintain its heartbeat the action specified by heartbeat-action will be taken. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


