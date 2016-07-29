# NetTunnelsV6rd

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Ipv4prefix** | **string** | As an extension not mentioned in the RFC5969, it specifies the IPv4 prefix for the Customer-Edge (CE) devices of a 6RD domain at a Border-Relay (BR) in case that the subnet prefixes used by the 6RD devices do not share the same IPv4 prefix. If they do, there is no need to configure this parameter. The default value is B 0.0.0.0 . | [optional] [default to null]
**V6rdprefixlen** | **int64** | Specifies the IPv6 prefix length of the 6rd domain. The default is 56. | [optional] [default to 56]
**Partition** | **string** | Displays the admin-partition within which this component resides. | [optional] [default to null]
**Ipv4prefixlen** | **int64** | Also noted as IPv4MaskLen in RFC5969, it specifies the number of identical high-order bits shared by all CE and BR IPv4 addresses in a given 6RD domain. The valid range is from zero to 32. It is a required value for create. It defaults to zero, i.e. the full ipv4 address must be encapsulated. | [optional] [default to 0]
**V6rdprefix** | **string** | Specifies the IPv6 prefix for 6rd domain. | [optional] [default to null]
**DefaultsFrom** | **string** | The profile from which to inherit settings. The default value is v6rd. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


