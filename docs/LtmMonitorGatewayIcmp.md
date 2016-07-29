# LtmMonitorGatewayIcmp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**AdaptiveSamplingTimespan** | **int64** | Specifies the size of the sliding window, in seconds, which records probe history.  For example, if this value is 300, then a sliding window of the last five minutes&#39; probe history will be used for calculating probe mean latency. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**UpInterval** | **int64** | Specifies how often in seconds that the system issues the monitor check when the node is up. The default value is the same as the (down) interval. | [optional] [default to null]
**ManualResume** | **string** | Specifies whether the system automatically changes the status of a resource to enabled at the next successful monitor check. The default value is disabled. If you set this option to enabled, you must manually re-enable the resource before the system can use it for load balancing connections. | [optional] [default to null]
**AdaptiveDivergenceValue** | **int64** | Specifies how far from mean latency each monitor probe is allowed to be.  If adaptive-divergence-type is relative, this value is a percentage deviation from mean (e.g. 50 would indicate the probe is allowed to exceed the mean latency by 50%.)  If adaptive-divergence-type is absolute, this value is an offset from mean in milliseconds (e.g. 250 would indicate the probe is allowed allowed to exceed the mean latency by 250 ms.) | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing monitor from which the system imports settings for the new monitor. | [optional] [default to null]
**Destination** | **string** | Specifies the IP address and service port of the resource that is the destination of this monitor. Possible values are:  *:* (Specifies to perform a health check on the IP address and port supplied by a pool member), *:port (Specifies to perform a health check on the server with the IP address supplied by the pool member and the port you specify.),   IP : port  (Specifies to mark a pool member up or down based on the response of the server at the IP address and port you specify.),   IP : port   with the transparent option enabled (Specifies to perform a health check on the server at the IP address and port specified in the monitor, routing the check through the IP address and port supplied by the pool member.  The pool member (the gateway) is marked up or down accordingly.). | [optional] [default to null]
**AdaptiveLimit** | **int64** | Specifies the hard limit, in milliseconds, which the probe is not allowed to exceed, regardless of the divergence value.  For example, if this value is 500, then the probe latency may not exceed 500 ms even if that would still fall within the divergence value. | [optional] [default to null]
**Adaptive** | **string** | Specifies whether the adaptive feature is enabled for this monitor.  Not all monitors support the adaptive feature. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Interval** | **int64** | Specifies the frequency at which the system issues the monitor check. The default value is 5 seconds. | [optional] [default to null]
**Transparent** | **string** | Specifies whether the monitor operates in transparent mode. Monitors in transparent mode can monitor pool members through firewalls. The default value is disabled. | [optional] [default to null]
**TimeUntilUp** | **int64** | Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0. | [optional] [default to null]
**Timeout** | **int64** | Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 16 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire. | [optional] [default to null]
**AdaptiveDivergenceType** | **string** | Specifies whether the adaptive-divergence-value is relative or absolute. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


