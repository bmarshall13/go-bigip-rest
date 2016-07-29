# NetIpsecIpsecPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**IkePhase2PerfectForwardSecrecy** | **string** | Defines the group of Diffie-Hellman exponentiations. This attribute is only valid when IKE is used to negotiate security associations. The value &#39;none&#39; indicates that the PFS is disabled for phase2 SA negotiations. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**IkePhase2AuthAlgorithm** | **string** | Specifies an payload authentication algorithm for ESP. This attribute is only valid when IKE is used to negotiate Security Associations. The possible options are: aes-gcm128, aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256, sha256, sha384, sha512 and sha1. The default value is aes-gcm128. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**IkePhase2LifetimeKilobytes** | **int64** | Specifies the lifetime duration in kilobytes, for the dynamically-negotiated security associations (SA). This attribute is only valid when IKE is used to negotiate security associations. A value of &#39;0&#39; means the SA will not re-key based on the number of bytes encrypted/decrypted. The minimum recommended value is 1000 kilobytes. This value is not negotiated between peers. | [optional] [default to 0]
**TunnelRemoteAddress** | **string** | Specifies the IP address of the remote IPsec tunnel endpoint. This option is only valid when mode is tunnel mode. | [optional] [default to null]
**Ipcomp** | **string** | Specifies the compression algorithm for IPComp. | [optional] [default to null]
**Mode** | **string** | Specifies a security protocol mode for use. The options are: transport, tunnel, isession and interface. | [optional] [default to null]
**TunnelLocalAddress** | **string** | Specifies the IP address of the local IPsec tunnel endpoint. This option is only valid when mode is tunnel mode. | [optional] [default to null]
**IkePhase2Lifetime** | **int64** | Specifies the lifetime duration in minutes, for the dynamically-negotiated security associations (SA). This attribute is only valid when IKE is used to negotiate security associations. | [optional] [default to 1440]
**Protocol** | **string** | Specifies the IPsec protocol: Encapsulating Security Payload (ESP) or Authentication Header (AH). | [optional] [default to null]
**IkePhase2EncryptAlgorithm** | **string** | Specifies an encryption algorithm for ESP. This attribute is only valid when IKE is used to negotiate security associations. The default value is B aes-gcm128 . | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


