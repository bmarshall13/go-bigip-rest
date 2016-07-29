# LtmAuthRadius

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Retries** | **int64** | Specifies the number of authentication retries that the BIG-IP local traffic management system allows before authentication fails. The default value is 3. | [optional] [default to 3]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**AccountingBug** | **string** | Enables or disables validation of the accounting response vector. This option should be necessary only on older servers. The default value is disabled. | [optional] [default to null]
**Partition** | **string** | Displays the partition within which the component resides. | [optional] [default to null]
**ServiceType** | **string** | Specifies the type of service requested from the RADIUS server.  The default value is authenticate-only | [optional] [default to null]
**ClientId** | **string** | Sends a NAS-Identifier RADIUS attribute with string bar. If you do not specify a value for the client-id option, the system uses the pluggable authentication module (PAM) service type. You can disable this feature by specifying a blank client ID. | [optional] [default to null]
**Debug** | **string** | Enables or disables syslog-ng debugging information at LOG DEBUG level. Not recommended for normal use. The default value is disabled. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


