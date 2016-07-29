# LtmAuthSslCcLdap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertmapUserSerial** | **string** | Specifies whether the system uses the client certificate&#39;s subject or serial number (in conjunction with the certificate&#39;s issuer) when trying to match an entry in the certificate map subtree. A value of yes uses the serial number. A value of no uses the subject. The default value is no. | [optional] [default to null]
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**UserBase** | **string** | Specifies the search base for the subtree used by the user and cert search methods. A typical search base is: ou&#x3D;people,dc&#x3D;company,dc&#x3D;com. Possible values are a user-specified string, and none. You must specify a user base when you create an SSL client certificate configuration object. | [optional] [default to null]
**Partition** | **string** | Displays the partition within which the server resides. | [optional] [default to null]
**Secure** | **string** | Enables or disables an attempt to use secure LDAP (LDAP over SSL). The alternative to using secure LDAP is to use insecure (clear text) LDAP. Secure LDAP is a consideration when the connection between the BIG-IP system and the LDAP server cannot be trusted. The default value is disabled. | [optional] [default to null]
**AdminPassword** | **string** | Specifies the password for the admin account. See the admin dn option above. Possible values are a user-specified string, and none. | [optional] [default to null]
**CertmapBase** | **string** | Specifies the search base for the subtree used by the certmap search method. A typical search base is: ou&#x3D;people,dc&#x3D;company,dc&#x3D;com. Possible values are a user-specified string, and none. | [optional] [default to null]
**ValidGroups** | **string** | Specifies a space-delimited list specifying the names of groups that the client must belong to in order to be authorized (matches against the group key in the group subtree). The client needs to be a member of only one of the groups in the list. Possible values are a user-specified string, or none. | [optional] [default to null]
**Servers** | **string** | Specifies a list of LDAP servers you want to search. Possible values are a user-specified list of servers, and none. You must specify a server when you create an SSL client certificate configuration object. | [default to null]
**CacheTimeout** | **int64** | Specifies the number of usable lifetime seconds of negotiable SSL session IDs. When this time expires, a client must negotiate a new session. Allowed values are:  number , immediate, and indefinite. The default value is 300 seconds. | [optional] [default to 300]
**UserKey** | **string** | Specifies the key that denotes a user ID in the LDAP database (for example, the common key for the user option is uid). Possible values are a user-specified string, and none. You must always specify a user key when you create an SSL client certificate configuration object. | [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**UserClass** | **string** | Specifies the object class in the LDAP database to which the user must belong in order to be authenticated. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**AdminDn** | **string** | Specifies the distinguished name of an account to which to bind, in order to perform searches. This search account is a read-only account used to do searches. The admin account can also be used as the search account. If no admin DN is specified, then no bind is attempted. This parameter is required only when an LDAP database does not allow anonymous searches. Possible values are a user-specified string, and none. | [optional] [default to null]
**RoleKey** | **string** | Specifies the name of the attribute in the LDAP database that specifies a user&#39;s authorization roles. This key is used only with the valid roles option. A typical role key might be authorizationRole. Possible values are a user-specified string, and none. | [optional] [default to null]
**CertmapKey** | **string** | Specifies the name of the certificate map found in the LDAP database. Used by the certmap search method. Possible values are a user-specified string, and none. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**GroupKey** | **string** | Specifies the name of the attribute in the LDAP database that specifies the group name in the group subtree. An example of a typical key is cn (common name for the group). Possible values are a user-specified string, and none. | [optional] [default to null]
**SearchType** | **string** | Specifies the type of LDAP search that is performed based on the client&#39;s certificate. | [optional] [default to null]
**GroupBase** | **string** | Specifies the search base for the subtree used by group searches. This parameter is only used when specifying the valid groups option. The typical search base is similar to: ou&#x3D;groups,dc&#x3D;company,dc&#x3D;com. Possible values are a user-specified string, and none. | [optional] [default to null]
**GroupMemberKey** | **string** | Specifies the name of the attribute in the LDAP database that specifies members (DNs) of a group. A typical key would be member. Possible values are a user-specified string, and none. | [optional] [default to null]
**CacheSize** | **int64** | Specifies the maximum size, in bytes, allowed for the SSL session cache. Setting this value to 0 disallows SSL session caching. The default value is 20000 bytes (that is 20KB). | [optional] [default to 20000]
**ValidRoles** | **string** | Specifies a space-delimited list specifying the valid roles that clients must have in order to be authorized. Possible values are a user-specified string, and none. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


