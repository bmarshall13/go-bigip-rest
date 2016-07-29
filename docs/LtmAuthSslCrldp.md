# LtmAuthSslCrldp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**UseIssuer** | **string** | Specifies whether the system extracts the CRL distribution point from the client certificate. The default value is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the partition within which the server resides. | [optional] [default to null]
**CacheTimeout** | **int64** | Specifies the number of seconds that CRLs will be cached. The default value is 86400 seconds. | [optional] [default to 86400]
**ConnectionTimeout** | **int64** | Specifies the number of seconds before the connection times out. The default value is 15 seconds. | [optional] [default to 15]
**UpdateInterval** | **int64** | Specifies an update interval for CRL distribution points. The update interval for distribution points ensures that CRL status is checked at regular intervals, regardless of the CRL timeout value. This helps to prevent CRL information from becoming outdated before the BIG-IP system checks the status of a certificate. The default value is 0 (zero), which indicates an internal default value is active. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


