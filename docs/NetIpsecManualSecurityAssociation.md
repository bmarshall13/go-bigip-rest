# NetIpsecManualSecurityAssociation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**AuthKeyEncrypted** | **string** | Displays the encrypted key for the authentication algorithm. | [optional] [default to null]
**DestinationAddress** | **string** | Specifies the destination of the security association. | [optional] [default to null]
**EncryptKeyEncrypted** | **string** | Displays the encrypted key for the encryption algorithm. | [optional] [default to null]
**AuthAlgorithm** | **string** | Specifies an authentication algorithm. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**EncryptAlgorithm** | **string** | Specifies an encryption algorithm. | [optional] [default to null]
**IpsecPolicy** | **string** | Specifies the ipsec-policy associated with this manual-security-association. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**AuthKey** | **string** | Specifies the key for the authentication algorithm. | [optional] [default to null]
**Spi** | **int64** | Specifies the Security Parameters Index. If this is the Security Association(SA) for the outbound traffic, make sure it matches the SPI of the inbound SA configured on the remote site and vice versa. SPI values between 0 and 255 are reserved for the future use by IANA and cannot be used. | [optional] [default to null]
**Protocol** | **string** | Specifies the IPsec protocol: Encapsulating Security Payload (ESP) or Authentication Header (AH). | [optional] [default to null]
**EncryptKey** | **string** | Specifies the key for the encryption algorithm. | [optional] [default to null]
**SourceAddress** | **string** | Specifies the source address of the security association. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


