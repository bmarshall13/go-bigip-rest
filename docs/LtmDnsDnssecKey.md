# LtmDnsDnssecKey

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**ExpirationPeriod** | **string** | Specifies the length of time before the key expires. The default value is 0 (zero).  This value must be greater than the value of the rollover-period option. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**BitWidth** | **int64** | Specifies the length of the key you want to generate. The default value is 1024. If the key is manually managed, MCPD will derive this value from the key file. | [optional] [default to 1024]
**Disabled** | **bool** | Specifies that the DNSSEC key is disabled. | [optional] [default to null]
**UseFips** | **string** | Specifies the type of FIPS-compliant hardware security module to use when storing, and signing with, the private key. The default value is none.  Optionally external or internal. | [optional] [default to null]
**Ttl** | **int64** | Specifies the number of seconds that a DNS server can cache the key. The default value is 86400. | [optional] [default to 86400]
**KeyFile** | **string** | Specifies the file containing the private key.  Fields certificate-file and key-file are required for manual DNSSEC key import. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**SignatureValidPeriod** | **string** | Specifies the length of time (seconds) that the signature of the key is valid.  The default value is 604800. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Algorithm** | **string** | Specifies the algorithm to use to generate the key.  The default value is RSASHA1. | [optional] [default to null]
**CertificateFile** | **string** | Specifies the file containing the public key.  Fields certificate-file and key-file are required for manual DNSSEC key import. | [optional] [default to null]
**KeyType** | **string** | Specifies whether the key is of type KSK or ZSK.  The default value is ZSK. | [optional] [default to null]
**Enabled** | **bool** | Specifies that the DNSSEC key is enabled. | [optional] [default to null]
**SignaturePubPeriod** | **string** | Specifies the length of time before we create a new signature.  The default value is 403200. | [optional] [default to null]
**RolloverPeriod** | **string** | Specifies the length of time before the key changes to another key. The default value is 0 (zero).  This value must be less than the value of the expiration-period option. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


