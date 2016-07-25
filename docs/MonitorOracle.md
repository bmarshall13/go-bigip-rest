# MonitorOracle

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Username** | **string** | Specifies the user name if the monitored target requires authentication. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**UpInterval** | **int64** | Specifies how often in seconds that the system issues the monitor check when the node is up. The default value is the same as the (down) interval. | [optional] [default to null]
**ManualResume** | **string** | Specifies whether the system automatically changes the status of a resource to enabled at the next successful monitor check. The default value is disabled. If you set this option to enabled, you must manually re-enable the resource before the system can use it for load balancing connections. | [optional] [default to null]
**RecvRow** | **string** | Specifies the row in the database where the specified Receive String should be located. This is an optional setting, and is applicable only if you configure the Send String and the Receive String settings. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing monitor from which the system imports settings for the new monitor. | [optional] [default to null]
**Destination** | **string** | Specifies the IP address and service port of the resource that is the destination of this monitor. Possible values are:  *:* (Specifies to perform a health check on the IP address and port supplied by a pool member), *:port (Specifies to perform a health check on the server with the IP address supplied by the pool member and the port you specify.), and  IP : port  (Specifies to mark a pool member up or down based on the response of the server at the IP address and port you specify.). | [optional] [default to null]
**Password** | **string** | Specifies the password if the monitored target requires authentication. | [optional] [default to null]
**Recv** | **string** | Specifies the text string that the monitor looks for in the returned resource. The most common receive expressions contain a text string that is included in a field in your database. If you do not specify both a Send String and a Receive String, the monitor performs a simple service check and connect only. | [optional] [default to null]
**Count** | **string** | Specifies the number of monitor probes after which the connection to the database will be terminated. Count value of zero indicates that the connection will never be terminated. The default value is zero. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Database** | **string** | Specifies the name of the database that the monitor tries to access. The proper format for name is  node_ip : node_port : database_name . | [optional] [default to null]
**Interval** | **int64** | Specifies the frequency at which the system issues the monitor check. The default value is 30 seconds. | [optional] [default to null]
**RecvColumn** | **string** | Specifies the column in the database where the specified Receive String should be located. This is an optional setting, and is applicable only if you configure the Send String and the Receive String settings. | [optional] [default to null]
**Send** | **string** | Specifies the SQL query that the monitor sends to the target object. For example, SELECT count(*) FROM mytable returns the number of rows in mytable. Since the string may have special characters, the system may require that the string be enclosed with single quotation marks. If this value is null, then a valid connection suffices to determine that the service is up. In this case, the system does not need the recv, recv-row, and recv-column options and ignores the options even if not null. | [optional] [default to null]
**TimeUntilUp** | **int64** | Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0. | [optional] [default to null]
**Timeout** | **int64** | Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 91 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire. | [optional] [default to null]
**Debug** | **string** | Specifies whether the monitor sends error messages and additional information to a log file created and labeled specifically for this monitor. The default setting is no. You can use the log information to help diagnose and troubleshoot unsuccessful health checks. The options are &#39;no&#39; (specifies that the system does not redirect error messages and additional information related to this monitor to the log file) and &#39;yes&#39; (specifies that the system redirects error messages and additional information to the /var/log/monitors/ monitor_name - node_name - port .log file). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


