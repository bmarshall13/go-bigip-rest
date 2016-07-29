# NetRoute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gw** | **string** | Specifies a gateway address for the route. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the route resides. | [optional] [default to null]
**Blackhole** | **bool** | Drops traffic addressed to the subnet. | [optional] [default to null]
**Mtu** | **int64** | Sets a specific maximum transmission unit (MTU). | [optional] [default to null]
**TmInterface** | **string** | Specifies a VLAN for the route. This can be a VLAN or VLAN group. | [optional] [default to null]
**Pool** | **string** | Specifies a gateway pool, which allows multiple, load-balanced gateways to be used for the route. | [optional] [default to null]
**Network** | **string** | The destination subnet and netmask for the route. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


