# LtmAuthOcspResponder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | **string** | Specifies whether the system verifies an OCSP response signature or the nonce values. The default value is enabled. | [optional] [default to null]
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**SignOther** | **string** | Adds a list of additional certificates to an OCSP request. | [optional] [default to null]
**VaFile** | **string** | Specifies the name of the file containing explicitly-trusted responder certificates. This parameter is needed in the event that the responder is not covered by the certificates already loaded into the responder&#39;s CA store. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Chain** | **string** | Specifies whether the system constructs a chain from certificates in the OCSP response. The default value is enabled. | [optional] [default to null]
**TrustOther** | **string** | Instructs the BIG-IP local traffic management system to trust the certificates specified with the verify-other option. The default value is disabled. | [optional] [default to null]
**Signer** | **string** | Specifies a certificate used to sign an OCSP request. If the certificate is specified but the key is not specified, then the private key is read from the same file as the certificate. If neither the certificate nor the key is specified, then the request is not signed. If the certificate is not specified and the key is specified, then the configuration is considered to be invalid. | [optional] [default to null]
**Intern** | **string** | Specifies whether the system ignores certificates contained in an OCSP response when searching for the signer&#39;s certificate. To use this setting, the signer&#39;s certificate must be specified with either the verify-other or va-file option. The default value is enabled. | [optional] [default to null]
**CheckCerts** | **string** | Enables or disables verification of an OCSP response certificate. Use this option for debugging purposes only. The default value is enabled. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Verify** | **string** | Enables or disables verification of an OCSP response signature or the nonce values. Used for debugging purposes only. The default value is enabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**VerifyCert** | **string** | Specifies that the system makes additional checks to see if the signer&#39;s certificate is authorized to provide the necessary status information. Use this option for testing purposes only. The default value is enabled. | [optional] [default to null]
**VerifyOther** | **string** | Specifies the name of the file used to search for an OCSP response signing certificate when the certificate has been omitted from the response. | [optional] [default to null]
**SignKeyPassPhrase** | **string** | Specifies the passphrase that the system uses to encrypt the sign key.  The default value is none. | [optional] [default to null]
**IgnoreAia** | **string** | Specifies whether the system ignores the URL contained in the certificate&#39;s AIA fields, and always uses the URL specified by the responder instead. The default value is disabled. | [optional] [default to null]
**CertIdDigest** | **string** | Specifies a specific algorithm identifier, either sha1 or md5. sha1 is newer and provides more security with a 160 bit hash length. md5 is older and has only a 128 bit hash length. The default values is sha1. The cert ID is part of the OCSP protocol. The OCSP client (in this case, the BIG-IP system) calculates the cert ID using a hash of the Issuer and serial number for the certificate that it is trying to verify. | [optional] [default to null]
**Partition** | **string** | Displays the partition within which the server resides. | [optional] [default to null]
**Explicit** | **string** | Specifies that the BIG-IP local traffic management system explicitly trusts that the OCSP response signer&#39;s certificate is authorized for OCSP response signing. If the signer&#39;s certificate does not contain the OCSP signing extension, specification of this option causes a response to be untrusted. The default value is enabled. | [optional] [default to null]
**SignKey** | **string** | Specifies the key that the system uses to sign an OCSP request. The default value is none. | [optional] [default to null]
**Url** | **string** | Specifies the URL used to contact the OCSP service on the responder. When using the ocsp responder command, you must specify a URL. | [optional] [default to null]
**CaPath** | **string** | Specifies the name of the path containing trusted CA certificates used to verify the signature on the OCSP response. | [optional] [default to null]
**StatusAge** | **int64** | Specifies the age of the status of the OCSP responder. The default value is 0 (zero). | [optional] [default to 0]
**VerifySig** | **string** | Specifies that the system checks the signature on the OCSP response. Use this option for testing purposes only. The default value is enabled. | [optional] [default to null]
**SignDigest** | **string** | Specifies the algorithm for signing the request, using the signing certificate and key. This parameter has no meaning if request signing is not in effect (that is, both the request signing certificate and request signing key parameters are empty). This parameter is required only when request signing is in effect. The default value is sha1. | [optional] [default to null]
**CaFile** | **string** | Specifies the name of the file containing trusted CA certificates used to verify the signature on the OCSP response. | [optional] [default to null]
**AllowCerts** | **string** | Enables or disables the addition of certificates to an OCSP request. The default value is enabled. | [optional] [default to null]
**ValidityPeriod** | **int64** | Specifies the number of seconds used to specify an acceptable error range. This option is used when the OCSP responder clock and a client clock are not synchronized, which could cause a certificate status check to fail. This value must be a positive number. The default value is 300 seconds. | [optional] [default to 300]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


