# ProfilePcp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ListeningPort** | **int64** | PCP server listening  port  for the PCP profile. Set to 5351; read-only. | [optional] [default to 5351]
**AnnounceMulticast** | **int64** | Specifies the count of multicast ANNOUNCE packets whenever the system sends one [0 - 50]. Default 10. | [optional] [default to 10]
**MapRecycleDelay** | **int64** | Specifies the delay before re-using an unused translation address (after the lifetime ends for the address&#39;s mapping) [0 - 600]. Default 60. | [optional] [default to 60]
**MinMappingLifetime** | **int64** | Specifies the minimum lifetime for a PCP mapping [120 - 3600]s. Default 600s. | [optional] [default to 600]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. The default value is pcp. | [optional] [default to null]
**MulticastPort** | **int64** | PCP Server multicast  port  for the PCP profile. Set to 5350; read-only. | [optional] [default to 5350]
**PeerOperAllowed** | **string** | Allows Peer operation for this PCP profile. Disabled; read-only. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**MaxMappingLifetime** | **int64** | Specifies the maximum lifetime for a PCP mapping [min-mapping-lifetime - 86400]s. Default 86400s. | [optional] [default to 86400]
**Name** | **string** | Name of entity | [optional] [default to null]
**AnnounceAfterFailover** | **string** | Enable or disable multicast ANNOUNCE responses after failover. Default disable. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the PCP profile resides. | [optional] [default to null]
**Rule** | **string** | Specifies the name of the iRule that will be used with this profile. | [optional] [default to null]
**MapFilterLimit** | **int64** | Specifies the maximum filters (external IPs that can exclusively use a given PCP mapping) per PCP client [1 - 16]. Default 1. | [optional] [default to 1]
**MapLimitPerClient** | **int64** | Specifies the maximum simultaneous PCP mappings per client [1 - 65535]. Default 65535. | [optional] [default to 65535]
**ThirdPartyOption** | **string** | Allows the third party option (ability for clients to add PCP mappings on behalf of third parties). Default disable. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


