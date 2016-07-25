# LsnPool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**ClientConnectionLimit** | **int64** | Configures the number of connections allowed per client. Default value of 0 disables this option. | [optional] [default to 0]
**Description** | **string** | User defined description. | [optional] [default to null]
**HairpinMode** | **string** | Specifies whether hairpinning for incoming connections is enabled or disabled on LSN translated IP addresses. | [optional] [default to null]
**TranslationPortRange** | **string** | Configures the port range used for translated addresses. | [optional] [default to null]
**IcmpEcho** | **string** | Enable or disable ICMP echo response on translation IP addresses. | [optional] [default to null]
**EgressInterfacesDisabled** | **bool** | Specifies that LSN translated traffic may not exit on the interfaces specified in egress-interfaces list. | [optional] [default to null]
**RouteAdvertisement** | **string** | Enable or disable route advertisements for LSN translation IP addresses. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**LogPublisher** | **string** | Configures the log publisher that handles LSN translation events logging for this LSN pool. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**InboundConnections** | **string** | Specifies whether inbound connections are enabled, and if they are created automatically or explicitly. | [optional] [default to null]
**EgressInterfacesEnabled** | **bool** | Specifies that LSN translated traffic may exit only on the interfaces specified in egress-interfaces list. | [optional] [default to null]
**Mode** | **string** | Configures the translation mode of operation for this LSN pool. | [optional] [default to null]
**LogProfile** | **string** | Configures the lsn log profile that controls logging. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


