# NetIpsecIkePeer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeersIdType** | **string** | Specifies that address, fqdn, asn1dn, user-fqdn, or keyid-tag can be used as peers-id-type. | [optional] [default to null]
**MyIdValue** | **string** | Specifies the identifier value sent to the remote host to use in the phase 1 negotiation. | [optional] [default to null]
**Mode** | **string** | Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder. | [optional] [default to null]
**Prf** | **string** | Specifies the pseudo-random function used to derive keying material for all cryptographic operations. This attribute is only valid for IKEv2. Possible values are: sha1, sha256, sha384, or sha512 | [optional] [default to null]
**TrafficSelector** | **string** | Specifies the names of the traffic-selector objects associated with this ike-peer. | [optional] [default to null]
**Lifetime** | **int64** | Defines the lifetime of an IKE SA which will be proposed in the phase 1 negotiations. | [optional] [default to 1440]
**Phase1EncryptAlgorithm** | **string** | Specifies the encryption algorithm used for the isakmp phase 1 negotiation. This directive must be defined. Possible value is one of following: des, 3des, blowfish, cast128, aes, or camellia for Oakley. | [optional] [default to null]
**VerifyCert** | **string** | If set to true, the identifier sent by the remote host (as specified in its my_identifier statement) is compared with the credentials in the certificate used to authenticate the remote host as follows: Type asn1dn: The entire certificate subject name is compared with the identifier, e.g. \&quot;C&#x3D;XX, O&#x3D;YY, ...\&quot;. Type address, fqdn, or user_fqdn: The certificate&#39;s subjectAltName is compared with the identifier. If the two do not match the negotiation will fail. The default value is false, which is not to verify the identifier using the peer&#39;s certificate. | [optional] [default to null]
**RemoteAddress** | **string** | Specifies the IP address of the IKE remote node. | [optional] [default to null]
**PeersCertFile** | **string** | Specifies the peer&#39;s certificate. This is no longer needed in IKEv2. | [optional] [default to null]
**Passive** | **string** | Specify true if you do not want to be the initiator of the IKE negotiation with this ike-peer. | [optional] [default to null]
**NatTraversal** | **string** | Enables use of the NAT-Traversal IPsec extension (NAT-T). NAT-T allows one or both peers to reside behind a NAT gateway (that is, doing address- or port-translation). The presence of NAT gateways along the path is discovered during the phase 1 handshake, and if found, NAT-T is negotiated. When NAT-T is in charge, all ESP and AH packets of a given connection are encapsulated into UDP datagrams (port 4500, by default). The options are: on, off, and force. | [optional] [default to null]
**State** | **string** | Enables or disables this IKE remote node. | [optional] [default to null]
**Version** | **string** | Specifies which version of IKE to be used. The default value is v1. Possible values are: v1 or v2. | [optional] [default to null]
**Phase1PerfectForwardSecrecy** | **string** | Defines the group used for the Diffie-Hellman exponentiations to provide perfect forward secrecy. This directive must be defined. The group is one of following: modp768, modp1024, modp1536, modp2048, modp3072, modp4096, modp6144, or modp8192. | [optional] [default to null]
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**ReplayWindowSize** | **int64** | Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node | [optional] [default to 64]
**Description** | **string** | User defined description. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**DpdDelay** | **int64** | Specifies the number of seconds between Dead Peer Detection messages. | [optional] [default to 30]
**PeersCertType** | **string** | DEPRECATED - Specifies that the only peers-cert-type supported is certfile. This is no longer needed in IKEv2. | [optional] [default to null]
**PresharedKeyEncrypted** | **string** | Display the encrypted preshared-key for the IKE remote node. | [optional] [default to null]
**MyCertFile** | **string** | Specifies the name of ssl-crt object for the certificate file. | [optional] [default to null]
**CaCertFile** | **string** | Specifies the file name, which contains the certificates of the trusted root and intermediate certificate authorities. | [optional] [default to null]
**GeneratePolicy** | **string** | Enable or disable the generation of Security Policy Database entries(SPD) when the device is the the responder of the IKE remote node. | [optional] [default to null]
**PresharedKey** | **string** | Specifies the preshared key for ISAKMP SAs. This field is valid only when phase1-auth-method is pre-shared-key. | [optional] [default to null]
**MyCertKeyFile** | **string** | Specifies the name of ssl-key object for the certificate key file. | [optional] [default to null]
**CrlFile** | **string** | Specifies the file name of the Certificate Revocation List. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Phase1HashAlgorithm** | **string** | Defines the hash algorithm used for the isakmp phase 1 negotiation. This directive must be defined. The algorithm should be one of following: md5, sha1, sha256, sha384, or sha512 for Oakley. | [optional] [default to null]
**MyIdType** | **string** | Specifies the identifier type sent to the remote host to use in the phase 1 negotiation. | [optional] [default to null]
**ProxySupport** | **string** | If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs.  This attribute must be enabled, which is the default value. This field is used only for IKEv1. | [optional] [default to null]
**Phase1AuthMethod** | **string** | Defines the authentication method used for the phase 1 negotiation. Possible values are: pre-shared-key if using pre-shared passphrase, and dss, ecdsa-256, ecdsa-384, ecdsa-521 or rsa-signature if using X.509 certificate-based authentication. | [optional] [default to null]
**PeersIdValue** | **string** | Specifies the peer&#39;s identifier to be received. If it is not defined, then the IKE agent will not verify the peer&#39;s identifier in the ID payload transmitted from the peer. The usage of peers-id-type and peers-id-value is the same as my-id-type and my-id-value except that the individual component values of an asn1dn identifier may specified as * to match any value (e.g. \&quot;C&#x3D;XX, O&#x3D;MyOrg, OU&#x3D;*, CN&#x3D;Mine\&quot;). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


