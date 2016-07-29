# LtmMonitorSip

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Destination** | **string** | Specifies the IP address and service port of the resource that is the destination of this monitor. Possible values are:  *:* (Specifies to perform a health check on the IP address and port supplied by a pool member), *:port (Specifies to perform a health check on the server with the IP address supplied by the pool member and the port you specify.), and  IP : port  (Specifies to mark a pool member up or down based on the response of the server at the IP address and port you specify.). | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**UpInterval** | **int64** | Specifies how often in seconds that the system issues the monitor check when the node is up. The default value is the same as the (down) interval. | [optional] [default to null]
**FilterNeg** | **string** | Specifies the SIP status codes that the target can return to be considered down. By default the system always accepts status codes according to sip-monitor.filter. After checking that, the status code is checked against this key. The options are none (Specifies that the monitor does not specifically reject any other SIP status codes.), any (Specifies that the monitor rejects all SIP status codes that are not in the sip-monitor.filter property.), and status (Specifies one or more status codes that you want to add to the monitor.)  If a code is also in sip-monitor.filter, the node will be marked up. | [optional] [default to null]
**ManualResume** | **string** | Specifies whether the system automatically changes the status of a resource to enabled at the next successful monitor check. If you set this option to yes, you must manually re-enable the resource before the system can use it for load balancing connections. | [optional] [default to null]
**Key** | **string** | Specifies the key if the monitored target requires authentication. | [optional] [default to null]
**Interval** | **int64** | Specifies the frequency at which the system issues the monitor check. The default value is 5 seconds. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing monitor from which the system imports settings for the new monitor. | [optional] [default to null]
**Compatibility** | **string** | Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Headers** | **string** | Specifies the set of SIP headers in the SIP message that is sent to the target. Each header should be separated by a newline. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**Request** | **string** | Specifies the SIP request line in the SIP message that is sent to the target. | [optional] [default to null]
**Filter** | **string** | Specifies the SIP status codes that the target can return to be considered up. By default the system always accepts status codes whose value is in the 100, 200 or 300s. The options are none (Specifies that the monitor does not accept any other SIP status codes.), any (Specifies that the monitor accepts any SIP status codes.), and status (Specifies one or more status codes that you want to add to the monitor.) | [optional] [default to null]
**TimeUntilUp** | **int64** | Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0. | [optional] [default to null]
**Cert** | **string** | Specifies a fully-qualified path for a client certificate that the monitor sends to the target SSL server. | [optional] [default to null]
**Mode** | **string** | Specifies the protocol that the monitor uses to communicate with the target. The options are tcp (the monitor uses TCP to communicate with the target) tls (the monitor uses TLS to communicate with the target) sips (the monitor uses TLS to communicate with the target, and the SIP URI is SIPS) and udp (the monitor uses UDP to communicate with the target) | [optional] [default to null]
**Timeout** | **int64** | Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 16 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire. | [optional] [default to null]
**Cipherlist** | **string** | Specifies the list of ciphers for this monitor. The default value is \&quot;DEFAULT:+SHA:+3DES:+kEDH\&quot;. | [optional] [default to null]
**Debug** | **string** | Specifies whether the monitor sends error messages and additional information to a log file created and labeled specifically for this monitor. The default setting is no. You can use the log information to help diagnose and troubleshoot unsuccessful health checks. The options are: no (Specifies that the system does not redirect error messages and additional information related to this monitor.) and yes (Specifies that the system redirects error messages and additional information to the /var/log/monitors/ monitor_name - node_name - port .log file.). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


