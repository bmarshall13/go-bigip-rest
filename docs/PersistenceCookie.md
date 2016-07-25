# PersistenceCookie

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**CookieEncryptionPassphrase** | **string** | Specifies a passphrase to be used for cookie encryption. | [optional] [default to null]
**Secure** | **string** | Specifies whether the secure attribute should be enabled or disabled for the inserted cookies. The default value is enabled. | [optional] [default to null]
**OverrideConnectionLimit** | **string** | Specifies, when enabled, that the pool member connection limits are not enforced for persisted clients. Per-virtual connection limits remain hard limits and are not disabled. The default value is disabled. | [optional] [default to null]
**MatchAcrossServices** | **string** | Specifies, when enabled, that all persistent connections from a client IP address, which go to the same virtual IP address, also go to the same node. The default value is disabled. | [optional] [default to null]
**CookieName** | **string** | Specifies a unique name for the profile. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing profile from which the system imports settings for the new profile. | [optional] [default to null]
**Mirror** | **string** | Specifies whether the system mirrors persistence records to the high-availability peer. This option is applicable only when the value of the method option is hash. The default value is disabled. | [optional] [default to null]
**AlwaysSend** | **string** | Specifies, when enabled, that the cookie persistence entry will be sent to the client on every response, rather than only on the first response.  The default value is disabled. | [optional] [default to null]
**HashLength** | **int64** | Specifies the cookie hash length. The length is the number of bytes to use when calculating the hash value. The default value is 0 bytes. | [optional] [default to 0]
**MatchAcrossPools** | **string** | Specifies, when enabled, that the system can use any pool that contains this persistence record. The default value is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**CookieEncryption** | **string** | Specifies the way in which cookie format will be used: \&quot;disabled\&quot;: generate old format,unencrypted, \&quot;preferred\&quot;: generate encrypted cookie but accept both encrypted and old format, and \&quot;required\&quot;: cookie format must be encrypted. Default is required. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the profile resides. | [optional] [default to null]
**MatchAcrossVirtuals** | **string** | Specifies, when enabled, that all persistent connections from the same client IP address go to the same node. The default value is disabled. | [optional] [default to null]
**Mode** | **string** |  | [optional] [default to null]
**Timeout** | **string** | Specifies the duration of the persistence entries. The default value is 180 seconds. | [optional] [default to null]
**HashOffset** | **int64** | Specifies the cookie hash offset. The offset is the number of bytes in the cookie to skip before calculating the hash value. The default value is 0 bytes. | [optional] [default to 0]
**Method** | **string** | Specifies the type of cookie processing that the system uses. The default value is insert. | [optional] [default to null]
**Httponly** | **string** | Specifies whether the httponly attribute should be enabled or disabled for the inserted cookies. The default value is enabled. | [optional] [default to null]
**Expiration** | **string** | Specifies the cookie expiration date in the format d:h:m:s, h:m:s, m:s or seconds. Hours 0-23, minutes 0-59, seconds 0-59. The time period must be less than 24856 days. You can use \&quot;session-cookie\&quot; (0 seconds) to indicate that the cookie expires when the browser closes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


