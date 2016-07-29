# NetRouterAdvertisement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**Managed** | **bool** | Indicates the Managed address configuration flag field in the router advertisement be set to 1. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Vlan** | **string** | Name of the VLAN used by this router advertisement. | [optional] [default to null]
**RouterLifetime** | **string** | Specifies the value to be used for the Router Lifetime field in the Router Advertisement. The default value is equal to 3 * max-interval. Valid values are between 600 and 9000. | [optional] [default to null]
**Disabled** | **bool** | Disables router advertisement for the VLAN. This is the default. | [optional] [default to null]
**RetransmitTimer** | **int64** | Specifies the value to be used for the Retransmit Timer field in the Router Advertisement. The default is 0. | [optional] [default to 0]
**Unmanaged** | **bool** | Indicates the Managed address configuration flag field in the router advertisement be set to 0. This is the default. | [optional] [default to null]
**Mtu** | **int64** | Set a specific maximum transmission unit (MTU) for the VLAN. The default value is 0. | [optional] [default to 0]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**OtherConfig** | **bool** | Indicates the Other Configuration flag field in the router advertisement be set to 1. | [optional] [default to null]
**MinInterval** | **int64** | Specifies the minimum time allowed between sending unsolicited multicast Router Advertisements from the interface, in seconds. The default value is 200. | [optional] [default to 200]
**Name** | **string** | Name of entity | [optional] [default to null]
**CurrentHopLimit** | **int64** | Defines the hop limit sent in the router advertisement. The range is from 0 to 255. The default value is 0. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**Enabled** | **bool** | Enables router advertisement for the VLAN. | [optional] [default to null]
**ReachableTime** | **int64** | Specifies the value to be used for the Reachable Time field in the Router Advertisement. The default is 0. | [optional] [default to 0]
**MaxInterval** | **int64** | Specifies the maximum time allowed between sending unsolicited multicast Router Advertisements from the interface, in seconds. The default value is 600. | [optional] [default to 600]
**NoOtherConfig** | **bool** | Indicates the Other Configuration flag field in the router advertisement be set to 0. This is the default. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


