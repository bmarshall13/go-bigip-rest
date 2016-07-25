# AuthProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**CredentialSource** | **string** | Specifies the credential source. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**CookieKey** | **string** | Specifies the key that is used to encrypt the cookie-name. The default value is \&quot;f5auth\&quot;. This setting applies to KRB Delegate profiles. | [optional] [default to null]
**Enabled** | **string** | Specifies whether the authentication profile is enabled. The default value is yes. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this component resides. | [optional] [default to null]
**IdleTimeout** | **int64** | Sets the idle timeout for the authentication profile. The default value is 300 seconds. | [optional] [default to 300]
**CookieName** | **string** | Specifies a unique session cookie assigned to each user. Each virtual server should use a different cookie name. The cookie-name is encrypted using the cookie-key. The default value is \&quot;abc123\&quot;. This setting applies to KRB Delegate profiles. | [optional] [default to null]
**Rule** | **string** | Specifies the name of the rule that corresponds to the authentication method you want to use with this profile. | [optional] [default to null]
**FileName** | **string** |  | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the name of the default authentication profile from which you want your custom profile to inherit settings. This setting is required. | [default to null]
**Configuration** | **string** | Specifies the name of a previously created authentication component. This setting is required. | [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


