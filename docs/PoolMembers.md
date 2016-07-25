# PoolMembers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Session** | **string** | Enables or disables the pool member for new sessions. The default value is user-enabled. | [optional] [default to null]
**DynamicRatio** | **int64** | Specifies a range of numbers that you want the system to use in conjunction with the ratio load balancing method. The default value is 1. | [optional] [default to 1]
**ConnectionLimit** | **int64** | Specifies the maximum number of concurrent connections allowed for a pool member. The default value is 0 (zero). | [optional] [default to 0]
**Address** | **string** | IP address of a pool member if a node by the given name does not already exist. | [optional] [default to null]
**Ratio** | **int64** | Specifies the ratio weight that you want to assign to the pool member. The default value is 1. | [optional] [default to 1]
**Monitor** | **string** | Displays the health monitors that are configured to monitor the pool member, and the status of each monitor. The default value is default. | [optional] [default to null]
**InheritProfile** | **string** | Specifies whether the pool member inherits the encapsulation profile from the parent pool. The default value is enabled. If you disable inheritance, no encapsulation takes place, unless you specify another encapsulation profile for the pool member using the profiles attribute. | [optional] [default to null]
**Logging** | **string** | Specifies whether the pool member&#39;s monitor(s) actions will be logged. Logs are stored in /var/log/monitors/ | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Ephemeral** | **string** |  | [optional] [default to null]
**State** | **string** | user-down forces the pool member offline, overriding monitors. user-up reverts the user-down. When user-up, this displays the monitor state. | [optional] [default to null]
**RateLimit** | **string** | Specifies the maximum number of connections per second allowed for a pool member. The default value is &#39;disabled&#39;. | [optional] [default to null]
**PriorityGroup** | **int64** | Specifies the priority group within the pool for this pool member. The priority group number specifies that traffic is directed to that member before being directed to a member of a lower priority. The default value is 0. | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


