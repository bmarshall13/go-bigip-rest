# LtmMonitorWmi

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Username** | **string** | Specifies the user name if the monitored target requires authentication. | [optional] [default to null]
**Interval** | **int64** | Specifies the frequency at which the system issues the monitor check. The default value is 5 seconds. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Agent** | **string** | Displays the agent for the monitor. The default agent is Mozilla/4.0 (compatible: MSIE 5.0; Windows NT). You cannot modify the agent. | [optional] [default to null]
**Metrics** | **string** | Specifies the performance metrics that the commands collect from the target. The default value is LoadPercentage, DiskUsage, PhysicalMemoryUsage:1.5, VirtualMemoryUsage:2.0 | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing monitor from which the system imports settings for the new monitor. | [optional] [default to null]
**Post** | **string** | Specifies the mechanism that the monitor uses for posting. The default value is RespFormat&#x3D;HTML. You cannot change the post format for WMI monitors. | [optional] [default to null]
**Password** | **string** | Specifies the password if the monitored target requires authentication. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Url** | **string** | Specifies the URL that the monitor uses. The default value is /scripts/f5Isapi.dll. | [optional] [default to null]
**TmCommand** | **string** | Specifies the command that the system uses to obtain the metrics from the resource. See the documentation for this resource for information on available commands. | [optional] [default to null]
**Destination** | **string** | Specifies the IP address of the resource that is the destination of this monitor. Possible values are:  * (Specifies to perform a health check on the IP address of the node.), and  IP  (Specifies to perform a health check on the IP address that you specify, route the check through the IP address of the associated node, and mark the IP address of the associated node up or down accordingly.). | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**TimeUntilUp** | **int64** | Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0. | [optional] [default to null]
**Timeout** | **int64** | Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 16 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire. | [optional] [default to null]
**Method** | **string** | Displays the GET method. You cannot modify the method. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


