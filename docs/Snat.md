# Snat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the snat resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Snatpool** | **string** | Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used. | [optional] [default to null]
**VlansDisabled** | **bool** | Disables the SNAT on all VLANs. | [optional] [default to null]
**AutoLasthop** | **string** | Specifies whether to automatically map last hop for pools or not. The default is to use next level&#39;s default. | [optional] [default to null]
**VlansEnabled** | **bool** | Enables the SNAT on all VLANs. | [optional] [default to null]
**Mirror** | **string** | Enables or disables mirroring of SNAT connections. | [optional] [default to null]
**SourcePort** | **string** | Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses. | [optional] [default to null]
**Translation** | **string** | Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used. | [optional] [default to null]
**Automap** | **bool** | Specifies that the system translate the source IP address to an available self IP address when establishing connections through the virtual IP. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


