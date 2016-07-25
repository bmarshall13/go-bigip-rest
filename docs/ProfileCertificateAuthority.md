# ProfileCertificateAuthority

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | Displays the application service to which the object belongs. The default value is none. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**CrlFile** | **string** | Specifies the certificate revocation list file name. You can use default for the default certificate revocation file name. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**AuthenticateDepth** | **int64** | Specifies the authenticate depth. This is the client certificate chain maximum traversal depth. | [optional] [default to null]
**Partition** | **string** | Specifies the administrative partition within which the profile resides. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the name of the default certificate authority profile from which you want your custom profile to inherit settings. This setting is required. | [optional] [default to null]
**LocationSpecific** | **string** |  | [optional] [default to null]
**UpdateCrl** | **string** | Automatically updates the CRL file. | [optional] [default to null]
**CaFile** | **string** | Specifies the certificate authority file name or, you can use default for the default certificate authority file name. Configures certificate verification by specifying a list of client or server certificate authorities that the traffic management system trusts. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


