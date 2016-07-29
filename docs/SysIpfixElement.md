# SysIpfixElement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description | [optional] [default to null]
**DataType** | **string** | Specify the data-type of the IPFIX element. | [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**EnterpriseId** | **int64** | Specify the IPFIX element enterprise id between 0 and 4294967295. | [default to 0]
**Id** | **int64** | Specify the IPFIX element id between 1 and 65535. Values greater than 32767 will be considered NETFLOW-only IEs. | [default to 1]
**Size** | **int64** | Specify the size between 1 and 1900. Used only for octetarray and string data-type. Default is 0, which means variable. | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


