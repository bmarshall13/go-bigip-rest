# NetTrunk

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**MacAddress** | **string** | Specifies the media access control (MAC) address, which is associated with the trunk, in case-insensitive hexadecimal colon notation, for example, 00:0b:09:88:00:9a. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**StpReset** | **bool** | Resets STP, which forces a migration check. | [optional] [default to null]
**CfgMbrCount** | **int64** | Specifies the number of configured members. | [optional] [default to 0]
**WorkingMbrCount** | **int64** | Specifies the number of working members associated with this trunk. | [optional] [default to 0]
**Bandwidth** | **int64** | Specifies the operational bandwidth in Mobs. | [optional] [default to 0]
**LacpTimeout** | **string** | Specifies the rate at which the system sends the LACP control packets. | [optional] [default to null]
**LacpMode** | **string** | Specifies the operation mode for link aggregation control protocol (LACP), if LACP is enabled for the trunk. | [optional] [default to null]
**Lacp** | **string** | Specifies, when enabled, that the system supports the link aggregation control protocol (LACP), which monitors the trunk by exchanging control packets over the member links to determine the health of the links. If LACP detects a failure in a member link, it removes the link from the link aggregation. LACP is disabled by default, for backward compatibility. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**QinqEthertype** | **string** | Specifies the protocol identifier associated with the tagged mode of the trunk. | [optional] [default to null]
**Media** | **string** | Displays the media settings for the trunk. | [optional] [default to null]
**Id** | **int64** | Displays the ID of the trunk. | [optional] [default to 65535]
**DistributionHash** | **string** | Specifies the basis for the hash that the system uses as the frame distribution algorithm. The system uses the resulting hash to determine which interface to use for forwarding traffic. | [optional] [default to null]
**Stp** | **string** | Enables or disables STP. If you disable STP, the system does not transmit or receive STP, RSTP, or MSTP packets on the trunk, and STP has no control over forwarding or learning on the trunk. The default value is enabled. | [optional] [default to null]
**LinkSelectPolicy** | **string** | Sets the LACP policy that the trunk uses to determine which member link (interface) can handle new traffic. Note that link aggregation is allowed only when all the interfaces are operating at the same media speed and connected to the same partner aggregation system. When there is a mismatch among configured members due to configuration errors or topology changes (auto-negotiation), link selection policy determines which links become working members and form the aggregation. | [optional] [default to null]
**Type_** | **string** | Displays whether this trunk is managed by a VCMP hypervisor or not. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


