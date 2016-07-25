# SnatTranslation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Arp** | **string** | Indicates whether or not the system responds to ARP requests or sends gratuitous ARPs. The default value is enabled. | [optional] [default to null]
**IpIdleTimeout** | **string** | Specifies the number of seconds that IP connections initiated using a SNAT address are allowed to remain idle before being automatically disconnected. The default value is indefinite. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**TcpIdleTimeout** | **string** | Specifies the number of seconds that TCP connections initiated using a SNAT address are allowed to remain idle before being automatically disconnected. The default value is indefinite. | [optional] [default to null]
**TrafficGroup** | **string** | Specifies the traffic group of the SNAT. The default is inherited from the containing folder. | [optional] [default to null]
**Disabled** | **bool** | Disables SNAT translation on the system. | [optional] [default to null]
**UdpIdleTimeout** | **string** | Specifies the number of seconds that UDP connections initiated using a SNAT address are allowed to remain idle before being automatically disconnected. The default value is indefinite. | [optional] [default to null]
**ConnectionLimit** | **int64** | Specifies the number of connections a translation address must reach before it no longer initiates a connection. The default value of 0 disables this option. | [optional] [default to 0]
**Address** | **string** | The translation IP address. | [optional] [default to null]
**Unit** | **int64** |  | [optional] [default to 1]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the snat resides. | [optional] [default to null]
**Enabled** | **bool** | Enables SNAT translation on the system. This is the default setting. | [optional] [default to null]
**InheritedTrafficGroup** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


