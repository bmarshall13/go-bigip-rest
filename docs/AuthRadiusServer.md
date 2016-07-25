# AuthRadiusServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the partition within which the server resides. | [optional] [default to null]
**Server** | **string** | Specifies the host name or IP address of the RADIUS server. This option is required. | [default to null]
**Secret** | **string** | Sets the secret key used to encrypt and decrypt packets sent or received from the server. This option is required. | [default to null]
**Timeout** | **int64** | Specifies the timeout value in seconds. The default value is 3 seconds. | [optional] [default to 3]
**Port** | **int64** | Specifies the port for RADIUS authentication traffic. The default value is port 1812. | [optional] [default to 1812]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


