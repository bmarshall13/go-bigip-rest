# ProfileDhcpv4

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**DefaultLeaseTime** | **int64** | Provides the default value in seconds of DHCPv4 lease time in case it was missing in the client-server exchange. The default is 86400. | [optional] [default to 86400]
**IdleTimeout** | **string** | Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 60 seconds. | [optional] [default to null]
**TransactionTimeout** | **int64** | Specifies the DHCPv4 transaction timeout, in seconds. The transaction should complete within the timeout specified. If a transaction does not complete for any reason, it is removed. The default value is 45 seconds. | [optional] [default to 45]
**TtlDecValue** | **string** | Specifies the amount that the DHCP virtual will use to decrement the ttl for each outgoing DHCP packet. Default is by-1. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**MaxHops** | **int64** | Specifies the maximum number of relay agents that the DHCPv4 messages pass through before they are discarded. The default is 4. | [optional] [default to 4]
**Mode** | **string** | Specifies the operation mode of the DHCP virtual. If the virtual to run in relay mode, then it is acting as a standard DHCPv4 relay agent. The forwarding mode is similar to relay except that the virtual will not modify the standard fields, instead it will forward the message from client to server and vice-versa. The default is relay. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile | [optional] [default to null]
**TtlValue** | **int64** | Specifies the ttl absolute value that the user may want to set for each outgoing DHCP packet. Default is 0, and in this case, we use the ttl-dec-value field. | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


