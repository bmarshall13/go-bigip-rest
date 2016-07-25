# GlobalSettingsGeneral

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**GratuitousArpRate** | **int64** | Specifies how fast gratuitous ARPs can be sent. If it is 0, then gratuitous ARPs are sent without pause. Otherwise, it specifies how many gratuitous ARPs can be sent every second. The default value is 0. | [optional] [default to null]
**MaintenanceMode** | **string** | Specifies, when enabled, that the unit is in maintenance mode. In maintenance mode, the system stops accepting new connections and slowly finishes off existing connections. | [optional] [default to null]
**ShareSingleMac** | **string** | Specifies the Media Access Control address (MAC address) that the system assigns to a VLAN. The value of unique indicates that a VLAN uses a mac address from a global mac address pool assigned to each hardware platform. The global value indicates that all of the VLANs on the system use the same MAC address.  The vmw-compat value indicates that the MAC address of a vlan is allocated in a manner compatible with VMware(tm) vSwitch and restricts VLANs to a single interface, with no trunks allowed.  Changing the value of this feature requires a manual restart of all daemons | [optional] [default to null]
**SnatPacketForward** | **string** | Enables or disables SNAT packet forwarding. The default is disable. | [optional] [default to null]
**L2CacheTimeout** | **int64** | Specifies, in seconds, the amount of time that records remain in the Layer 2 forwarding table, when the MAC address of the record is no longer detected on the network. The default is 300 seconds. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


