# ProfileFtp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**InheritParentProfile** | **string** | Enables the FTP data channel to inherit the TCP profile used by the control channel. If disabled, the data channel uses FastL4 (BigProto) only. The default is unchecked (disabled). | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**LogPublisher** | **string** | Configures the log publisher that handles events logging for this profile. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**LogProfile** | **string** | Configures the ALG log profile that controls logging. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**TranslateExtended** | **string** | This setting is enabled by default, and thus, automatically translates RFC 2428 extended requests EPSV and EPRT to PASV and PORT when communicating with IPv4 servers. | [optional] [default to null]
**Security** | **string** | Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the system is licensed for the BIG-IP Application Security Manager. The default value is disabled. | [optional] [default to null]
**AllowFtps** | **string** | Allow explicit FTPS negotiation. The default is disabled. | [optional] [default to null]
**Port** | **int64** | Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data. | [optional] [default to 20]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


