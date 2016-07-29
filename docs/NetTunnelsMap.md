# NetTunnelsMap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Ip6Prefix** | **string** | Rule IPv6 prefix in CIDR notation for MAP domain CE (Customer Edge) devices. The default prefix length is 48. | [optional] [default to null]
**Partition** | **string** | Displays the admin-partition within which this component resides. | [optional] [default to null]
**PortOffset** | **int64** | Specifies port offset bits length of the MAP domain. The default is 6. | [optional] [default to 6]
**EaBitsLength** | **int64** | Specifies EA (Embedded Address) length of the MAP domain. The default is 32 (IPv4 prefix 24 bits + PSID 8 bits). | [optional] [default to 32]
**DefaultsFrom** | **string** | The profile from which to inherit settings. The default value is map. | [optional] [default to null]
**Ip4Prefix** | **string** | Rule IPv4 prefix in CIDR notation for MAP domain CE (Customer Edge) devices. The default prefix length is 8. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


