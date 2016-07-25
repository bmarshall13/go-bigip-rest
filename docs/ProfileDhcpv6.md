# ProfileDhcpv6

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**DefaultLeaseTime** | **int64** | Provides the default value in seconds of DHCPv6 lease time in case it was missing in the client-server exchange. The default is 86400. | [optional] [default to 86400]
**TransactionTimeout** | **int64** | Specifies the DHCPv6 transaction timeout, in seconds. The transaction should complete within the timeout specified. If a transaction does not complete for any reason, it is removed. The default value is 45 seconds. | [optional] [default to 45]
**IdleTimeout** | **string** | Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 60 seconds. | [optional] [default to null]
**Mode** | **string** | Specifies the operation mode of the DHCP virtual. If the virtual to run in relay mode, then it is acting as a standard DHCPv6 relay agent. The forwarding mode is similar to relay except that the virtual will not modify the standard fields, instead it will forward the message from client to server and vice-versa. The default is relay. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


