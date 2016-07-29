# LtmNat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Arp** | **string** | Enables or disables Address Resolution Protocol (ARP). | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the nat resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**TrafficGroup** | **string** | Specifies the traffic group for the NAT. The default is inherited from the containing folder. | [optional] [default to null]
**Disabled** | **bool** | Disables the NAT. | [optional] [default to null]
**TranslationAddress** | **string** | Translation/destination IP address. This may not be changed after the nat has been created. | [default to null]
**Unit** | **int64** |  | [optional] [default to 1]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**OriginatingAddress** | **string** | Specifies the IP address from which traffic is being initiated. | [default to null]
**VlansDisabled** | **bool** | Indicates the NAT is disabled on the list of VLANs. | [optional] [default to null]
**Enabled** | **bool** | Enables the NAT. | [optional] [default to null]
**VlansEnabled** | **bool** | Indicates the NAT is enabled on the list of VLANs. | [optional] [default to null]
**AutoLasthop** | **string** | Specifies whether to automatically map last hop for pools or not. The default is to use next level&#39;s default. | [optional] [default to null]
**InheritedTrafficGroup** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


