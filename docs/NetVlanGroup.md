# NetVlanGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**BridgeInStandby** | **string** | When enabled, specifies that the VLAN group forwards packets, even when the system is the standby unit in a redundant system. Note that this setting is designed for deployments in which the VLAN group exists on only one of the units. If that does not match your configuration, using this setting may cause adverse effects. The default value is disabled. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the vlan-group resides. | [optional] [default to null]
**BridgeTraffic** | **string** | When enabled, specifies that the VLAN group forwards all frames, including non-IP traffic. The default value is disable. | [optional] [default to null]
**MigrationKeepalive** | **string** | Specifies whether the system will send keepalive frames (TCP keepalives and empty UDP packets depending on the connection type) when a node is moved from one vlan-group member to another vlan-group member for all existing connections that the system has to that node. | [optional] [default to null]
**BridgeMulticast** | **string** | When enabled, allows bridging of non-Internet Protocol (IP) Address Resolution Protocol (ARP) multicast frames across a VLAN group. An example of when you might want to use this option is when you are implementing the Spanning Tree Protocol (STP). | [optional] [default to null]
**Mode** | **string** | Specifies the level of exposure of remote MAC addresses within VLAN groups. The default value is translucent. | [optional] [default to null]
**IfIndex** | **int64** | Displays the index assigned to this VLAN group. It is a unique identifier assigned for all objects displayed in the SNMP IF-MIB. | [optional] [default to 0]
**AutoLasthop** | **string** | Specifies whether auto lasthop is enabled or not. The default is to use next levels default. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


