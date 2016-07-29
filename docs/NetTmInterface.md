# NetTmInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaFixed** | **string** | Specifies the settings for a fixed (non-pluggable) interface. Use this option only with a combo port to specify the media type for the fixed interface, when it is not the preferred port. | [optional] [default to null]
**MacAddress** | **string** | Displays the MAC address (6-byte Ethernet address in hexadecimal colon notation, for example, 00:0b:09:88:00:9a) of the interface. This is the hardware address. | [optional] [default to null]
**LldpTlvmap** | **int64** |  | [optional] [default to 130943]
**Vendor** | **string** | Displays the name of the vendor of the pluggable unit on an interface. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**StpAutoEdgePort** | **string** | Sets STP automatic edge port detection for the interface. The default value is true. When automatic edge port detection is set to true for an interface, the system monitors the interface for incoming STP, RSTP, or MSTP packets. If no such packets are received for a sufficient period of time (about three seconds), the interface is automatically given edge port status. When automatic edge port detection is set to false for an interface, the system never gives the interface edge port status automatically. Any STP setting set on a per-interface basis applies to all spanning tree instances. | [optional] [default to null]
**StpReset** | **bool** | Resets STP, which forces a migration check. | [optional] [default to null]
**PreferPort** | **string** | Indicates which side of a combo port the interface uses, if both sides have the potential for an external link. The default value for a combo port is sfp. Do not use this option for non-combo ports. | [optional] [default to null]
**Stp** | **string** | Enables or disables STP. If you disable STP, no STP, RSTP, or MSTP packets are transmitted or received on the interface or trunk, and spanning tree has no control over forwarding or learning on the port or the trunk. The default value is enabled. | [optional] [default to null]
**Bundle** | **string** |  | [optional] [default to null]
**Disabled** | **bool** | Disables the specified interfaces from passing traffic. | [optional] [default to null]
**MediaSfp** | **string** | Specifies the settings for an SFP (pluggable) interface. Use this option only with a combo port to specify the media type for the SFP interface, when it is not the preferred port. | [optional] [default to null]
**StpEdgePort** | **string** | Specifies whether the interface connects to an end station instead of another spanning tree bridge. The default value is true. | [optional] [default to null]
**IfIndex** | **int64** | Displays the index assigned to this interface. It is a unique identifier assigned for all objects displayed in the SNMP IF-MIB. | [optional] [default to 0]
**Serial** | **string** | Displays the serial number of the pluggable unit on an interface. | [optional] [default to null]
**MediaMax** | **string** | Displays the maximum media value for the interface. | [optional] [default to null]
**StpLinkType** | **string** | Specifies the STP link type for the interface. The default value is auto. The spanning tree system includes important optimizations that can only be used on point-to-point links, that is, on links which connect just two bridges. If these optimizations are used on shared links, incorrect or unstable behavior may result. By default, the implementation assumes that full-duplex links are point-to-point and that half-duplex links are shared. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**QinqEthertype** | **string** | Specifies the protocol identifier associated with the tagged mode of the interface. | [optional] [default to null]
**Media** | **string** | Specifies the settings for the interface. When you set the media option, the system automatically sets the media-sfp or media-fixed option based on whether the interface is a small form factor pluggable (SFP) interface, or for combo ports whether SFP is the preferred port. | [optional] [default to null]
**Enabled** | **bool** | Enables the specified interfaces to pass traffic. | [optional] [default to null]
**Mtu** | **int64** | Displays the maximum Transmission Unit (MTU) of the interface, which is the maximum number of bytes in a frame without IP fragmentation. | [optional] [default to 1500]
**LldpAdmin** | **string** |  | [optional] [default to null]
**FlowControl** | **string** | Specifies how the system controls the sending of PAUSE frames. The default value is tx-rx. | [optional] [default to null]
**ForceGigabitFiber** | **string** | Enables or disables forcing of gigabit fiber media. If this is enabled for a gigabit fiber interface, the media setting will be forced, and no auto-negotiation will be performed. If it is disabled, auto-negotiation will be performed with just a single gigabit fiber option advertised. | [optional] [default to null]
**MediaActive** | **string** | Displays the current media settings for the interface. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


