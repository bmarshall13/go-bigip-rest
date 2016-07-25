# ProfileSocks

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ProtocolVersions** | **string** | Specifies the SOCKS protocol versions that will be supported. | [optional] [default to null]
**DnsResolver** | **string** | Specifies the dns-resolver object that will be used to resolve hostnames in connect requests. | [optional] [default to null]
**TunnelName** | **string** | Specifies the tunnel that will be used for outbound proxy requests. This enables other virtual servers to receive connections initiated by the SOCKS service. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**RouteDomain** | **string** | Specifies the route-domain that will be used for outbound connect requests. | [optional] [default to null]
**DefaultConnectHandling** | **string** | Specifies the behavior of the proxy service for connect requests. If set to &#39;deny&#39;, connect requests will only be honored if there is another virtual server listening for the requested outbound connection. If set to &#39;allow&#39; outbound connections will be made regardless of other virtual servers. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


