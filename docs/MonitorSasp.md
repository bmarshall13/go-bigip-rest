# MonitorSasp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Destination** | **string** |  | [optional] [default to null]
**Protocol** | **string** | Specifies the protocol that the monitor uses to communicate with the target. The options are tcp and udp. The default value is tcp. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**MonInterval** | **int64** |  | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing monitor from which the system imports settings for the new monitor. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Service** | **string** | Specifies the port through which the SASP monitor communicates with the Group Workload Manager. The default port is 3000. | [optional] [default to null]
**Interval** | **string** | Specifies the frequency at which the system queries the Group Workload Manager (GWM). The default value is auto. You can specify a range between 10 and 600 seconds. | [optional] [default to null]
**PrimaryAddress** | **string** | Specifies the IP address of the Primary Group Workload Manager. | [optional] [default to null]
**SecondaryAddress** | **string** | Specifies the IP address of the Secondary Group Workload Manager. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**TimeUntilUp** | **int64** | Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0. | [optional] [default to null]
**Mode** | **string** | Specifies whether the load balancer should send Get Weight Request messages (pull) or receive Send Weight messages (push) from the GWM. The default mode is pull. | [optional] [default to null]
**Timeout** | **int64** | Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 100 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


