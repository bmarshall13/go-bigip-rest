# NetTunnelsTunnel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Profile** | **string** | Specifies the profile that you want to associate with the tunnel. The default value is gre. | [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**UsePmtu** | **string** | Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages. If enabled and the tunnel MTU is set to 0, the tunnel will use the PMTU information. If enabled and the tunnel MTU is fixed to a non-zero value, the tunnel will use the minimum of PMTU and MTU. If disabled, the tunnel will use fixed MTU or calculate its MTU using tunnel encapsulation configurations. | [optional] [default to null]
**TrafficGroup** | **string** | Specifies a traffic-group for use with the tunnel. Traffic group determines the ConfigSync behavior of the tunnel object. | [optional] [default to null]
**Key** | **int64** | The key field may represent different values depending on the type of the tunnel. E.g. it represents the Virtual Network Identifier(VNI) for VXLAN tunnels. The default value is 0. | [optional] [default to 0]
**IfIndex** | **int64** | Displays the index assigned to this tunnel. It is a unique identifier assigned for all objects displayed in the SNMP IF-MIB. | [optional] [default to 0]
**Mtu** | **int64** | Specifies the maximum transmission unit (MTU) of the tunnel. The default value is the MTU of the tunnel&#39;s underlying interface minus the the encapsulation overhead introduced by the tunneling protocol in use for the tunnel. | [optional] [default to 0]
**Transparent** | **string** | Enables or disables the tunnel to be transparent. If enabled, the user can inspect and/or manipulate the encapsulated traffic flowing through the BIG-IP. A transparent tunnel terminates a tunnel while presenting the illusion that the tunnel transits the device unperturbed i.e. the BIG-IP appears like an intermediate router that simply routes IP traffic through the device. The default value is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**RemoteAddress** | **string** | Specifies a remote IP address. This option is required. | [optional] [default to null]
**Tos** | **string** | Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets. The default value is preserve, which will cause the ToS octet from the encapsulated packet header to be copied into the ToS octet of the encapsulating packet header. The valid range is zero to 255 | [optional] [default to null]
**Partition** | **string** | Displays the admin-partition within which this component resides. | [optional] [default to null]
**IdleTimeout** | **int64** | Specifies an idle timeout for wildcard tunnels in seconds. This setting specifies the number of seconds that a wildcard tunnel connection is idle before the connection is eligible for deletion. The default value is 300 seconds. | [optional] [default to 300]
**SecondaryAddress** | **string** | Specifies a secondary non-floating IP address when the local-address is set to a floating address. Currently this setting is supported by NVGRE tunnels only. | [optional] [default to null]
**LocalAddress** | **string** | Specifies a local IP address. This option is required. | [optional] [default to null]
**Mode** | **string** | Specifies how the tunnel carries traffic. The default value is bidirectional. | [optional] [default to null]
**AutoLasthop** | **string** | Specifies whether auto lasthop is enabled or not. The default is to use next levels default. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


