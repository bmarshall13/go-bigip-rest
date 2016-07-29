# LtmAuthLdap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**SslClientCert** | **string** | Specifies the name of a SSL client certificate. The default value is none. | [optional] [default to null]
**Servers** | **string** | Specifies the LDAP servers that the system must use to obtain authentication information. You must specify a server when you create an LDAP configuration object. | [default to null]
**SslCiphers** | **string** | Specifies SSL ciphers.The default value is none. | [optional] [default to null]
**SslCaCertFile** | **string** | Specifies the name of an SSL CA certificate. The default value is none. | [optional] [default to null]
**GroupDn** | **string** | Specifies the group distinguished name. The system uses this option for authorizing client traffic. | [optional] [default to null]
**Port** | **int64** | Specifies the port name or number for the LDAP service. Port 389 is typically used for non-SSL and port 636 is used for an SSL-enabled LDAP service. | [optional] [default to 389]
**BindTimeout** | **int64** | Specifies a bind timeout limit, in seconds. The default value is 30 seconds. | [optional] [default to 30]
**SslClientKey** | **string** | Specifies the name of a SSL client key. The default value is none. | [optional] [default to null]
**Version** | **int64** | Specifies the version number of the LDAP application. The default value is 3. | [optional] [default to 3]
**BindPw** | **string** | Specifies the password for the search account created on the LDAP server. This option is required if you use a bind DN. | [optional] [default to null]
**SearchBaseDn** | **string** | Specifies the search base distinguished name. The default value is none. | [optional] [default to null]
**LoginAttribute** | **string** | Specifies a logon attribute. Normally, the value of this option is uid; however, if the server is a Microsoft Windows Active Directory server, the value must be the account name samaccountname (not case-sensitive). | [optional] [default to null]
**SearchTimeout** | **int64** | Specifies the search timeout, in seconds. The default value is 30 seconds. | [optional] [default to 30]
**Description** | **string** | User defined description. | [optional] [default to null]
**Warnings** | **string** | Enables or disables warning messages. The default value is enabled. | [optional] [default to null]
**CheckRolesGroup** | **string** | Specifies whether to verify a user&#39;s group membership given in the remote-role definitions, formatted as \&quot;*member*of&#x3D; group-dn \&quot;. | [optional] [default to null]
**SslCheckPeer** | **string** | Specifies that the system checks an SSL peer. The default value is disabled. | [optional] [default to null]
**IgnoreUnknownUser** | **string** | Specifies whether the system ignores an unknown user. The default value is disabled. | [optional] [default to null]
**Ssl** | **string** | Enables or disables SSL. The default value is disabled. Note that when you use the command line interface to enable SSL for an LDAP service, the system does not change the service port number from 389 to 636, as is required. To change the port number from the command line, use the service option for this component, for example, ldap [name] ssl enabled service 636. | [optional] [default to null]
**IgnoreAuthInfoUnavail** | **string** | Specifies whether the system ignores authentication information if it is not available. The default value is disabled. | [optional] [default to null]
**Scope** | **string** | Specifies the search scope. The default value is sub. | [optional] [default to null]
**GroupMemberAttribute** | **string** | Specifies a group member attribute. The system uses this option for authorizing client traffic. | [optional] [default to null]
**BindDn** | **string** | Specifies the distinguished name of an account to which to bind, in order to perform searches. This search account is a read-only account used to do searches. The admin account can be used as the search account. If no admin DN is specified, then no bind is attempted. This option is only required when a site does not allow anonymous searches. If the remote server is a Microsoft Windows Active Directory server, the distinguished name must be in the form of an email address. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Partition** | **string** | Displays the partition within which the server resides. | [optional] [default to null]
**IdleTimeout** | **int64** | Specifies the idle timeout, in seconds, for connections. The default value is 3600 seconds. | [optional] [default to 3600]
**Filter** | **string** | Specifies a filter. Use this option for authorizing client traffic. | [optional] [default to null]
**UserTemplate** | **string** | Specifies a user template for the LDAP application to use for authentication. The default value is none. | [optional] [default to null]
**Debug** | **string** | Enables or disables syslog-ng debugging information at LOG DEBUG level. The default value is disabled. F5 Networks does not recommend using this option for a normal configuration. | [optional] [default to null]
**CheckHostAttr** | **string** | Confirms the password for the bind distinguished name. This option is optional. The default value is disabled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


