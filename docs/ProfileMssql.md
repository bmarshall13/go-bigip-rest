# ProfileMssql

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ReadPool** | **string** | Specifies the pool of MS SQL database servers to which the system sends ready-only requests. | [optional] [default to null]
**Partition** | **string** | Specifies the administrative partition within which the profile resides. | [optional] [default to null]
**ReadWriteSplitByUser** | **string** | When enabled, the system decides which pool to send the client requests to by user name. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**UserList** | **string** | Specifies the users who have read-only access to the MS SQL if user-can-write-by-default is true, or the users who have write access to the MS SQL database if user-can-write-by-default is false. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**UserCanWriteByDefault** | **string** | Specifies whether users have write access by default. When set to true, all users have write access, except those added to the users list | [optional] [default to null]
**ReadWriteSplitByCommand** | **string** | When enabled, the system decides which pool to send the client requests to by the content in the message. | [optional] [default to null]
**WritePool** | **string** | Specifies the pool of MS SQL database servers to which the system sends requests that are not read-only | [optional] [default to null]
**WritePersistTimer** | **int64** | Specify how many minimum time in milliseconds the connection will be persisted to write-pool after connection switch to write pool. | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


