# LtmMonitorSnmpDca

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**CpuThreshold** | **string** | Specifies the maximum acceptable CPU usage on the target server. The default value is 80 percent. | [optional] [default to null]
**Community** | **string** | Specifies the community name that the BIG-IP system must use to authenticate with the host server through SNMP. The default value is public. | [optional] [default to null]
**Interval** | **int64** | Specifies the frequency at which the system issues the monitor check. The default value is 10 seconds. | [optional] [default to null]
**MemoryThreshold** | **string** | Specifies the maximum acceptable memory usage on the target server. The default value is 70 percent. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing monitor from which the system imports settings for the new monitor. | [optional] [default to null]
**DiskThreshold** | **string** | Specifies the maximum acceptable disk usage on the target server. The default value is 90 percent. | [optional] [default to null]
**DiskCoefficient** | **string** | Specifies the coefficient that the system uses to calculate the weight of the disk threshold in the dynamic ratio load balancing algorithm. The default value is 2.0. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**UserDefined** | **string** | Specifies user defined variables for the monitor. The name and value of the variable must be specified. If \&quot;none\&quot; is specified as the value then the variable is removed. | [optional] [default to null]
**AgentType** | **string** | Specifies the type of agent. The options are ucd, win2000, generic, and other. | [optional] [default to null]
**Destination** | **string** | Specifies the IP address of the resource that is the destination of this monitor. The default value is *. Possible values are: * (Specifies to perform a health check on the IP address of the node.) or   IP  (Specifies to perform a health check on the IP address that you specify, route the check through the IP address of the associated node, and mark the IP address of the associated node up or down accordingly.). | [optional] [default to null]
**CpuCoefficient** | **string** | Specifies the coefficient that the system uses to calculate the weight of the CPU threshold in the dynamic ratio load balancing algorithm. The default value is 1.5. | [optional] [default to null]
**TimeUntilUp** | **int64** | Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0. | [optional] [default to null]
**Version** | **string** | Specifies the version of SNMP that the host server uses. | [optional] [default to null]
**MemoryCoefficient** | **string** | Specifies the coefficient that the system uses to calculate the weight of the memory threshold in the dynamic ratio load balancing algorithm. The default value is 1.0. | [optional] [default to null]
**Timeout** | **int64** | Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 30 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


