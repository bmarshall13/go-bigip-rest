# MonitorInband

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Destination** | **string** |  | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**FailureInterval** | **int64** | Specifies an interval, in seconds. If the number of failures specified in the failures option occurs within this interval, the system marks the pool member as being unavailable. | [optional] [default to null]
**ResponseTime** | **int64** | Specifies an amount of time, in seconds. If the pool member does not respond with data after the specified amount of time has passed, the number of failures in this interval increments by 1. Specifying a value of 0 (zero) disables this option. | [optional] [default to null]
**RetryTime** | **int64** | Specifies the amount of time in seconds after the pool member has been marked unavailable before the system retries to connect to the pool member. Specifying a value of 0 (zero) disables this option. | [optional] [default to null]
**Failures** | **int64** | Specifies the number of times within a given time period that the system tries to connect to a pool member before marking that server as being unavailable. The default value is 3. Specifying a value of 0 (zero) disables this option. A failure can be either a failure to connect or a failure of the pool member to respond within the time specified in the response-time option. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the name of the monitor from which you want your custom monitor to inherit settings. The default value is inband. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


