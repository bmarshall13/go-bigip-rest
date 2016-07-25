# GlobalSettingsTrafficControl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptIpSourceRoute** | **string** | Specifies whether the system accepts IPv4 packets with IP source route options that are destined for Traffic Management Microkernel (TMM). The default is disable.  To enable this option, you must also enable the accept ip options setting. | [optional] [default to null]
**RejectUnmatched** | **string** | Specifies, when enabled, that the system returns a TCP RESET or ICMP_UNREACH packet if no virtual servers on the system match the destination address of the incoming packet. When this setting is disabled, the system silently drops the unmatched packet. The default is enable. | [optional] [default to null]
**AllowIpSourceRoute** | **string** | Specifies whether the system allows IPv4 packets with IP source route options enabled to be routed through Traffic Management Microkernel (TMM). The default is disable.  To enable this option, you must also enable the accept ip options setting. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**PathMtuDiscovery** | **string** | Specifies, when enabled, that the system discovers the maximum transmission unit (MTU) that it can send over a path, without fragmenting TCP packets. The default is enable. | [optional] [default to null]
**AcceptIpOptions** | **string** | Specifies whether the system accepts IPv4 packets with IP options. The default is disable. | [optional] [default to null]
**MaxRejectRate** | **int64** | Specifies the maximum rate per second that the system issues reject packets (TCP RST or ICMP port unreach). The default value is 250 per second. The range is from 1 to 1000 per second. | [optional] [default to null]
**MinPathMtu** | **int64** | Specifies the minimum packet size that can traverse the path without suffering fragmentation, also known as path Maximum Transmission Unit (MTU). The default is 296. The range is from 68 to 1500. | [optional] [default to null]
**PortFindThresholdTrigger** | **int64** | Specifies the threshold warning&#39;s trigger which is the value of random port attempts when attempting to find an unused outbound port for a connection.  A warning is issued when the second trigger level is reached within the timeout period. The default is 8. The valid range is 1 to 12. | [optional] [default to null]
**ContinueMatching** | **string** | Specifies whether the system matches against a less-specific virtual server when the more-specific one is disabled. When continue matching is disabled, the system drops connections that request a disabled virtual server. In this case, the system rejects or drops packets depending on the setting of the reject unmatched option. | [optional] [default to null]
**MaxIcmpRate** | **int64** | Specifies the maximum rate per second at which the system issues Internet Control Message Protocol (ICMP) errors. The default is 100 errors per second. The range is from 0 (zero) to 2147483647 errors per second.  This option is useful for preventing ICMP message storms. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**PortFindThresholdWarning** | **string** | Specifies if the ephemeral port-exhaustion threshold warning is to be monitored. The default is enabled. | [optional] [default to null]
**PortFindThresholdTimeout** | **int64** | Specifies the threshold warning&#39;s timeout. This is the time in seconds since the last trigger value was hit and will throttle the output warnings from logging too often. The default is 30 seconds with range from 0 to 300. | [optional] [default to null]
**PortFindRandom** | **int64** | Specifies the maximum of ports to randomly search for outbound connections. The default is 16. The range is from 0 to 1024. | [optional] [default to null]
**PortFindLinear** | **int64** | Specifies the maximum of ports to linearly search for outbound connections. The default is 16. The range is from 0 to 61439. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


