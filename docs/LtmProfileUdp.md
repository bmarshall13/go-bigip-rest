# LtmProfileUdp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**IpTtlMode** | **string** | Describe the outgoing UDP packet TTL mode. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**IdleTimeout** | **string** | Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 60 seconds. | [optional] [default to null]
**NoChecksum** | **string** | Enables or disables checksum processing. Note that if the datagram is IPv6, the system always performs checksum processing. The default value is disabled. | [optional] [default to null]
**IpTosToClient** | **string** | Specifies the Type of Service level that the traffic management system assigns to UDP packets when sending them to clients. | [optional] [default to null]
**AllowNoPayload** | **string** | Provides the ability to allow the passage of datagrams that contain header information, but no essential data. The default value is disabled. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**ProxyMss** | **string** | Make the system use the same max segment size on both ends. | [optional] [default to null]
**DatagramLoadBalancing** | **string** | Provides the ability to load balance UDP datagram by datagram. The default value is disabled. | [optional] [default to null]
**LinkQosToClient** | **string** | Specifies the Quality of Service level that the system assigns to UDP packets when sending them to clients. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


