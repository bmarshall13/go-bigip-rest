# NetVlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**CustomerTag** | **string** | Specifies a number that the system adds into the header of any frame passing through the QinQ VLAN. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**DagTunnel** | **string** | Specifies whether the ip tunnel traffic on the VLAN should be disaggregated based on the inner ip header or outer ip header. The default value is outer. | [optional] [default to null]
**CmpHash** | **string** | Specifies the cmp hash applied to the traffic on the vlan. | [optional] [default to null]
**Tag** | **int64** | Specifies a number that the system adds into the header of any frame passing through the VLAN. | [optional] [default to 0]
**Learning** | **string** | Specifies whether switch ports placed in the VLAN are configured for switch learning, forwarding only, or dropped. The default value is enable. | [optional] [default to null]
**IfIndex** | **int64** | Displays the index assigned to this VLAN. It is a unique identifier assigned for all objects displayed in the SNMP IF-MIB. | [optional] [default to 0]
**Mtu** | **int64** | Sets a specific maximum transmission unit (MTU) for the VLAN. The default value is 1500. | [optional] [default to 1500]
**DagRoundRobin** | **string** | Specifies whether some of the stateless traffic on the VLAN should be disaggregated in a round-robin order. | [optional] [default to null]
**SourceChecking** | **string** | Specifies that only connections that have a return route in the routing table are accepted. The default value is disable. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Failsafe** | **string** | Enables a fail-safe mechanism that causes the active unit to fail over to a redundant unit when loss of traffic is detected on a VLAN, and traffic is not restored during the failover timeout period for that VLAN. The default action set with VLAN fail-safe is restart all. When the fail-safe mechanism is triggered, all the daemons are restarted and the unit fails over. The default value is disable. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**FailsafeTimeout** | **int64** | Specifies the number of seconds that an active unit can run without detecting network traffic on this VLAN before it initiates a failover. The default value is 90 seconds. | [optional] [default to 90]
**Partition** | **string** | Displays the administrative partition within which the vlan resides. | [optional] [default to null]
**FailsafeAction** | **string** | Specifies the action that you want the system to take when it does not detect any traffic on this VLAN, and the timeout has expired. | [optional] [default to null]
**AutoLasthop** | **string** | Specifies whether auto lasthop is enabled or not. The default is to use next levels default. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


