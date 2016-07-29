# LtmMessageRoutingSipProfileSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**MaxMsgSize** | **int64** | Indicates the maximum number acceptable SIP message size (in bytes). | [optional] [default to 65535]
**Description** | **string** | User defined description. | [optional] [default to null]
**MaxMsgHeaderCount** | **int64** | Indicates the maximum count of expected SIP message header fields. | [optional] [default to 64]
**InsertRecordRouteHeader** | **string** | Enables or disables the insertion of a record-route header in requests that establish dialog. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**LoopDetection** | **string** | Enables or disables loop-detection checking | [optional] [default to null]
**HonorVia** | **string** | Enables or disables honoring any via which is not inserted by the system for routing the response. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all of the settings and values from the specified parent profile. | [optional] [default to null]
**CustomVia** | **string** | Specifies the custom value for the sent-by field of via | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**MaxForwardsCheck** | **string** | Enables or disables checking on max-forwards. | [optional] [default to null]
**EnableSipFirewall** | **string** | Indicates whether SIP firewall capability is enabled. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**DoNotConnectBack** | **string** | Enables or disables whether a connection to a request originator is re-established (if it no longer exists) in order to deliver a response. | [optional] [default to null]
**MaxMsgHeaderSize** | **int64** | Indicates the maximum SIP message header size (in bytes). | [optional] [default to 16000]
**GenerateResponseOnFailure** | **string** | Enables or disables sending failure response messages such as 4xx, 5xx and 6xx, when a SIP request is being dropped. | [optional] [default to null]
**InsertViaHeader** | **string** | Enables or disables insertion of top via. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


