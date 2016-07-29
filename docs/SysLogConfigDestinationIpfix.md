# SysLogConfigDestinationIpfix

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the log-destination object belongs | [optional] [default to null]
**TemplateRetransmitInterval** | **int64** | Specify a time interval in seconds after which active IPFIX Templates should be re-sent to the pool (30 seconds is the default) | [optional] [default to 30]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | A user defined description for this logging destination | [optional] [default to null]
**PoolName** | **string** | Specify the LTM pool of nodes (IPFIX Collectors) to receive IPFIX messages | [optional] [default to null]
**TransportProfile** | **string** | Specify the name of a transport-protocol profile (udp or tcp) that this logging destination uses to send IPFIX messages | [optional] [default to null]
**TemplateDeleteDelay** | **int64** | Specify a time interval in seconds after which deleted IPFIX Templates may be re-used (5 seconds is the default) - not implemented | [optional] [default to 5]
**ProtocolVersion** | **string** | Specify the protocol version to be used: netflow-9 or ipfix (ipfix is the default) | [optional] [default to null]
**ServersslProfile** | **string** | Specify the name of a serverssl profile that this logging destination uses to send encrypted IPFIX messages | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


