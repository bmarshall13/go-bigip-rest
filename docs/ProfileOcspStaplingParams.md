# ProfileOcspStaplingParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**SignerKeyPassphrase** | **string** | Specifies the passphrase of the key used for signing the OCSP request. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**DnsResolver** | **string** | Specifies the DNS resolver object used for fetching the OCSP response. | [optional] [default to null]
**StrictRespCertCheck** | **string** | If enabled, the responder&#39;s certificate is checked for OCSP signingextension. By default, it is disabled. | [optional] [default to null]
**CacheTimeout** | **string** | Specifies the lifetime of the OCSP response in the cache, in seconds.The actual time period for which the response is cached is the minimum of the response validity period and the cache-timeout. The default value is indefinite, indicating that the response validity period takes precedence. | [optional] [default to null]
**ClockSkew** | **int64** | Specifies the tolerable maximum absolute difference in the clocks of  the responder and the BIG-IP, in seconds. The default value is 300. | [optional] [default to 300]
**SignHash** | **string** | Specifies the hash algorithm used for signing the OCSP request.The default value is SHA256. | [optional] [default to null]
**ProxyServerPool** | **string** | Specifies the proxy server pool used for fetching the OCSP response. | [optional] [default to null]
**UseProxyServer** | **string** | Specifies whether the proxy server pool or the DNS resolver should be used for making the connection to the OCSP responder. | [optional] [default to null]
**ResponderUrl** | **string** | Specifies the absolute URL that overrides the OCSP responder URL obtained from the certificate&#39;s AIA extension(s). This should be a HTTP or HTTPS based URL. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**CacheErrorTimeout** | **int64** | Specifies the lifetime of an error response in the cache, in seconds.The default value is 3600 or one hour. | [optional] [default to 3600]
**Timeout** | **int64** | Specifies the time interval (in seconds) that the BIG-IP waits for before aborting the connection to the OCSP responder. The default value is 8. The timeout should be less than the handshake timeout of the clientssl profile that the OCSP Stapling Parameters object is associated with. | [optional] [default to 8]
**TrustedCa** | **string** | Specifies the certificate-authority that signs the responder&#39;s certificate. | [optional] [default to null]
**SignerCert** | **string** | Specifies the certificate corresponding to the key used for signing the OCSP request. | [optional] [default to null]
**StatusAge** | **int64** | Specifies the maximum allowed lag time for the &#39;thisUpdate&#39; time in the OCSP response that the BIG-IP accepts. If this maximum is exceeded, the response is dropped. If this value is set to &#39;0&#39;, this validation is skipped. The default value is 86400 seconds. | [optional] [default to 86400]
**SignerKey** | **string** | Specifies the key used for signing the OCSP request. | [optional] [default to null]
**TrustedResponders** | **string** | Specifies the certificate(s) used for validating the OCSP response when the responder&#39;s certificate has been omitted from the response. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


