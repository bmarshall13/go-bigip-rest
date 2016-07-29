# LtmProfileDns

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Dns64AdditionalSectionRewrite** | **string** | Set DNS64 additional section rewriting | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Dns64** | **string** | Set DNS64 mapping mode | [optional] [default to null]
**EnableDnsExpress** | **string** | Indicates whether the dns-express service is enabled. The service handles zone transfers from the primary DNS server. | [optional] [default to null]
**RapidResponseLastAction** | **string** | When Rapid Response is enabled, what to do with non-matching DNS queries. | [optional] [default to null]
**EnableLogging** | **string** | Indicates whether to perform DNS packet logging. | [optional] [default to null]
**Dns64Prefix** | **string** | IPv6 prefix for DNS64 mapping | [optional] [default to null]
**ProcessXfr** | **string** | Indicates whether the system answers zone transfer requests for a DNS zone created on the system. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**EnableDnsFirewall** | **string** | Indicates whether DNS firewall capability is enabled. | [optional] [default to null]
**AvrDnsstatSampleRate** | **int64** | Sets AVR DNS statistics sampling rate. 0: no query will be sent to the analytics database. 1: every query will be sent. N: every Nth query will be sent and the analytics database will count it N times. The default value is 0. When sampling rate is greater than one, the statistics will be inaccurate if the traffic volume is low. However, when the traffic volume is high, the system performance will benefit from sampling and the inaccuracy will be negligible. The sampling rate does not apply to DNS firewall statistics. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]
**EnableGtm** | **string** | Indicates whether to allow the BIG-IP global traffic management system to handle DNS resolution for DNS queries and responses that contain wide IP names. The default value is yes. | [optional] [default to null]
**ProcessRd** | **string** | Indicates whether to process clientside DNS packets with Recursion Desired set in the header. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**EnableHardwareQueryValidation** | **string** | On supported platforms, indicates whether the hardware will accelerate query validation. | [optional] [default to null]
**UnhandledQueryAction** | **string** | Last action to take if query name doesn&#39;t match. | [optional] [default to null]
**UseLocalBind** | **string** | Forward non-GTM and non-dns-express requests to local BIND. | [optional] [default to null]
**DnsSecurity** | **string** | Indicates the DNS security profile the system uses. | [optional] [default to null]
**EnableHardwareResponseCache** | **string** | On supported platforms, indicates whether the hardware will cache responses. | [optional] [default to null]
**EnableCache** | **string** | Indicates whether to allow queries to be answered non-authoritatively by a DNS cache. | [optional] [default to null]
**Cache** | **string** | DNS cache used to generate non-authoritative responses. | [optional] [default to null]
**EnableDnssec** | **string** | Indicates whether to perform DNS Security Extensions. | [optional] [default to null]
**EnableRapidResponse** | **string** | Indicates whether to allow queries to be answered by Rapid Response DNS. | [optional] [default to null]
**LogProfile** | **string** | Selects the desired logging profile. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


