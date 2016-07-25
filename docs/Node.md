# Node

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Monitor** | **string** | Specifies the name of the monitor or monitor rule that you want to associate with the node. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Logging** | **string** | Specifies whether the node&#39;s monitor(s) actions will be logged. Logs are stored in /var/log/monitors/ | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this node resides. | [optional] [default to null]
**Ephemeral** | **string** |  | [optional] [default to null]
**RateLimit** | **string** | Specifies the maximum number of connections per second allowed for a node or node address. The default value is &#39;disabled&#39;. | [optional] [default to null]
**State** | **string** | Marks the node up or down. The default value is user-up. | [optional] [default to null]
**Session** | **string** | Enables or disables the node for new sessions. The default value is user-enabled. | [optional] [default to null]
**DynamicRatio** | **int64** | Sets the dynamic ratio number for the node. Used for dynamic ratio load balancing. The ratio weights are based on continuous monitoring of the servers and are therefore continually changing. Dynamic Ratio load balancing may currently be implemented on RealNetworks RealServer platforms, on Windows platforms equipped with Windows Management Instrumentation (WMI), or on a server equipped with either the UC Davis SNMP agent or Windows 2000 Server SNMP agent. | [optional] [default to 1]
**ConnectionLimit** | **int64** | Specifies the maximum number of connections allowed for the node or node address. | [optional] [default to 0]
**Address** | **string** | IP address of the node. This is an optional field; if empty, the name needs to be of the form [ip address] | [optional] [default to null]
**Ratio** | **int64** | Specifies the fixed ratio value used for a node during ratio load balancing. | [optional] [default to 1]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


