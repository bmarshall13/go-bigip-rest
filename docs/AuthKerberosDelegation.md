# AuthKerberosDelegation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**DebugLogging** | **string** | Specifies whether debug logging is enabled. The default value is disabled. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ClientPrincipal** | **string** | Specifies the principal that the client sees. This is usually a value such as HTTP/ fqdn . The client principal may be in a different domain from the server principal. This setting is required. There is no default value. | [default to null]
**ProtocolTransition** | **string** | Specifies whether associated virtual should transition client certificate authentication into Kerberos credentials. | [optional] [default to null]
**ServerPrincipal** | **string** | Specifies the principal of the back-end web server. This is usually a value such as HTTP/ fqdn of server . The server principal may be in a different domain from the client principal. This setting is required. There is no default value. | [default to null]
**Partition** | **string** | Displays the administrative partition within which the configuration resides. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


