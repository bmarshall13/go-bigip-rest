# NetStpGlobals

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelloTime** | **int64** | Specifies the time interval in seconds between the periodic transmissions that communicate spanning tree information to the adjacent bridges in the network. The default value is 2 seconds, and the valid range is 1 to 10. The default hello time is optimal in virtually all cases. F5 recommends that you do not change the hello time. | [optional] [default to 2]
**TransmitHold** | **int64** | Specifies the absolute limit on the number of spanning tree protocol packets the traffic management system may transmit on a port in any hello time interval. It is used to ensure that spanning tree packets do not unduly load the network even in unstable situations. The default value is 6 packets, and the valid range is 1 to 10 packets. | [optional] [default to 6]
**ConfigName** | **string** | Specifies the configuration name (1 - 32 characters in length) only when the spanning tree mode is MSTP. The default configuration name is a string representation of a globally-unique MAC address belonging to the traffic management system.The MSTP standard introduces the concept of spanning tree regions, which are groups of adjacent bridges with identical configuration names, configuration revision levels, and assignments of VLANs to spanning tree instances. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**MaxAge** | **int64** | Specifies the number of seconds for which spanning tree information received from other bridges is considered valid. The default value is 20 seconds, and the valid range is 6 to 40 seconds. | [optional] [default to 20]
**MaxHops** | **int64** | Specifies the maximum number of hops an MSTP packet may travel before it is discarded. Use this option only when the spanning tree mode is MSTP. The number of hops must be in the range of 1 to 255 hops. The default number of hops is 20. | [optional] [default to 20]
**ConfigRevision** | **int64** | Specifies the revision level of the MSTP configuration only when the spanning tree mode is MSTP. The specified number must be in the range 0 to 65535. The default value is 0 (zero). | [optional] [default to 0]
**FwdDelay** | **int64** | In the original Spanning Tree Protocol, the forward delay parameter controlled the number of seconds for which an interface was blocked from forwarding network traffic after a reconfiguration of the spanning tree topology. This parameter has no effect when RSTP or MSTP are used, as long as all bridges in the spanning tree use the RSTP or MSTP protocol. If any legacy STP bridges are present, then neighboring bridges must fall back to the old protocol, whose reconfiguration time is affected by the forward delay value. The default forward delay value is 15, and the valid range is 4 to 30. | [optional] [default to 15]
**Mode** | **string** | Specifies the spanning tree modes. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


