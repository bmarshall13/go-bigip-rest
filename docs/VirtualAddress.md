# VirtualAddress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Arp** | **string** | Enables or disables ARP for the specified virtual address. The default value is enabled. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Enabled** | **string** | Specifies whether a virtual server IP address is enabled. The default value is yes. | [optional] [default to null]
**TrafficGroup** | **string** | Specifies the traffic group for the virtual address.  The default traffic group is inherited from the containing folder. | [optional] [default to null]
**ConnectionLimit** | **int64** | Sets a concurrent connection limit in seconds for one or more virtual servers. The default value is 0 seconds. | [optional] [default to 0]
**Address** | **string** |  | [optional] [default to null]
**Floating** | **string** |  | [optional] [default to null]
**IcmpEcho** | **string** | Enables or disables ICMP echo replies for the specified virtual address. The default value is enabled. | [optional] [default to null]
**Unit** | **int64** |  | [optional] [default to 1]
**RouteAdvertisement** | **string** | Enables or disables route advertisement for the specified virtual address. The default value is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**AutoDelete** | **string** | Indicates if the virtual address will be deleted automatically on deletion of the last associated virtual server or not. The default value is true. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**Mask** | **string** | Sets the netmask or one or more network virtual servers only. This setting is required for network virtual servers. | [optional] [default to null]
**ServerScope** | **string** | Specifies the server that uses the specified virtual address. | [optional] [default to null]
**InheritedTrafficGroup** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


