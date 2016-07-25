# ProfileClientSslCertKeyChain

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Chain** | **string** | Specifies or builds a certificate chain file that a client can use to authenticate the profile. To use the default chain name, specify default. | [optional] [default to null]
**Cert** | **string** | Specifies the name of the certificate installed on the traffic management system for the purpose of terminating or initiating an SSL connection. | [optional] [default to null]
**Key** | **string** | Specifies the name of a key file that you generated and installed on the system. The default key name is default.key. | [optional] [default to null]
**Passphrase** | **string** | Specifies the key passphrase if required. | [optional] [default to null]
**OcspStaplingParams** | **string** | Specifies the OCSP Stapling Parameters object associated with this cert-key-chain object in a clientssl profile. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


