# SysSmtpServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**Username** | **string** | The user name that the SMTP server requires when validating a user. | [optional] [default to null]
**AuthenticationDisabled** | **bool** | Specifies that the SMTP server does not validate users. | [optional] [default to null]
**EncryptedConnection** | **string** | Specifies which type of encrypted connection the SMTP server requires in order to send mail. The default is none. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**SmtpServerHostName** | **string** | The SMTP server host name in the format of a fully qualified domain name. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**PasswordEncrypted** | **string** | The password that the SMTP server requires when validating a user, in an encrypted form. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**LocalHostName** | **string** | The host name used in SMTP headers in the format of a fully qualified domain name. This setting does not refer to the BIG-IP system&#39;s Hostname. | [optional] [default to null]
**FromAddress** | **string** | The email address that the email is being sent from. This is the \&quot;Reply-to\&quot; address that the recipient sees. | [optional] [default to null]
**AuthenticationEnabled** | **bool** | Specifies that the SMTP server validates users before allowing them to send mail. Specify this only if the SMTP server requires authentication. | [optional] [default to null]
**SmtpServerPort** | **int64** | Specifies the SMTP port number. The default is 25. | [optional] [default to 25]
**Password** | **string** | The password that the SMTP server requires when validating a user. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


