# GlobalSettingsConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalFlowEvictionPolicy** | **string** | Assigns the flow eviction policy to use when memory limits are approached. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**VlanKeyedConn** | **string** | Enables or disables VLAN-keyed connections. You use VLAN-keyed connections when traffic for the same connection must pass through the system several times, on multiple pairs of VLANs (or in different VLAN groups). The default setting is enable. | [optional] [default to null]
**SyncookiesThreshold** | **int64** | Specifies the number of new or untrusted TCP connections that can be established before the system activates the SYN Cookies authentication method for subsequent TCP connections.  The default value is 16384. | [optional] [default to null]
**AdaptiveReaperLowater** | **int64** | *IMPORTANT* This command has been deprecated (as of 11.6.0). Please use ltm eviction-policy instead. Specifies, in percent, the memory usage at which the system silently purges stale connections, without sending reset packets (RST) to the client. If the memory usage remains above the low-water mark after the purge, then the system starts purging established connections closest to their service timeout. The default setting is 85. To disable the adaptive reaper, set the low-water mark to 100. | [optional] [default to null]
**AdaptiveReaperHiwater** | **int64** | *IMPORTANT* This command has been deprecated (as of 11.6.0). Please use ltm eviction-policy instead. Specifies, in a percentage, the memory usage at which the system stops establishing new connections. Once the system meets the reaper high-water mark, the system does not establish new connections until the memory usage drops below the reaper low-water mark. The default setting is 95. To disable the adaptive reaper, set the high-water mark to 100. The adaptive reaper settings help mitigate the effects of a denial-of-service attack. | [optional] [default to null]
**AutoLastHop** | **string** | Specifies that the system automatically maps the last hop for pools. The default is enable. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


