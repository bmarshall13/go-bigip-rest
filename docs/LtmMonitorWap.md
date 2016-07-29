# LtmMonitorWap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Destination** | **string** | Specifies the IP address and service port of the resource that is the destination of this monitor. Possible values are:  *:* (Specifies to perform a health check on the IP address and port supplied by a pool member), *:port (Specifies to perform a health check on the server with the IP address supplied by the pool member and the port you specify.), and  IP : port  (Specifies to mark a pool member up or down based on the response of the server at the IP address and port you specify.). | [optional] [default to null]
**AccountingNode** | **string** | Specifies the RADIUS server that provides authentication for the WAP target. This is an optional setting. Note that if you configure the accounting-port option, but you do not configure the accounting-node option, the system assumes that the RADIUS server and the WAP server are the same system. | [optional] [default to null]
**UpInterval** | **int64** | Specifies how often in seconds that the system issues the monitor check when the node is up. The default value is the same as the (down) interval. | [optional] [default to null]
**ManualResume** | **string** | Specifies whether the system automatically changes the status of a resource to enabled at the next successful monitor check. The default value is disabled. If you set this option to enabled, you must manually re-enable the resource before the system can use it for load balancing connections. | [optional] [default to null]
**AccountingPort** | **string** | Specifies the port that the monitor uses for RADIUS accounting. The default value is none. A value of 0 (zero) disables RADIUS accounting. | [optional] [default to null]
**Interval** | **int64** | Specifies the frequency at which the system issues the monitor check. The default value is 10 seconds. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing monitor from which the system imports settings for the new monitor. | [optional] [default to null]
**Recv** | **string** | Specifies the text string that the monitor looks for in the returned resource. The most common receive expressions contain a text string that is included in an HTML file on your site. The text string can be regular text, HTML tags, or image names. If you do not specify both a Send String and a Receive String, the monitor performs a simple service check and connect only. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**CallId** | **string** | Specifies the 11-digit phone number for the RADIUS server. This is an optional setting. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**SessionId** | **string** | Specifies the RADIUS session identification number when configuring a RADIUS server. This is an optional setting. | [optional] [default to null]
**Send** | **string** | Specifies the text string that the monitor sends to the target object. The default setting is GET /, which retrieves a default HTML file for a web site. To retrieve a specific page from a web site, specify a fully-qualified path name, for example, GET /www/company/index.html. Since the string may have special characters, the system may require that the string be enclosed with single quotation marks. If this value is null, then a valid connection suffices to determine that the service is up. In this case, the system does not need the recv, recv-row, and recv-column options and ignores the options even if not null. | [optional] [default to null]
**TimeUntilUp** | **int64** | Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0. | [optional] [default to null]
**Secret** | **string** | Specifies the password the monitor needs to access the resource. | [optional] [default to null]
**Timeout** | **int64** | Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 31 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire. | [optional] [default to null]
**Debug** | **string** | Specifies whether the monitor sends error messages and additional information to a log file created and labeled specifically for this monitor. The default setting is no. You can use the log information to help diagnose and troubleshoot unsuccessful health checks. The options are &#39;no&#39; (specifies that the system does not redirect error messages and additional information related to this monitor to the log file) and &#39;yes&#39; (specifies that the system redirects error messages and additional information to the /var/log/monitors/ monitor_name - node_name - port .log file). | [optional] [default to null]
**ServerId** | **string** | Specifies the RADIUS NAS-ID for this system when configuring a RADIUS server. | [optional] [default to null]
**FramedAddress** | **string** | Specifies the RADIUS framed IP address. This is an optional setting. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


