# ProfileSip

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**MaxRegistrations** | **int64** | Specifies the maximum number of registrations, the maximum allowable REGISTER messages can be recorded that the BIG-IP system accepts. The default value is 100. | [optional] [default to 100]
**Description** | **string** | User defined description. | [optional] [default to null]
**AlgEnable** | **string** | Indicates whether SIP ALG feature is enabled. The default value is disabled | [optional] [default to null]
**RegistrationTimeout** | **int64** | Indicates the timeout value for sip registration. The default value is 3600 seconds | [optional] [default to 3600]
**TerminateOnBye** | **string** | Enables or disables the termination of a connection when a BYE transaction finishes. Use this parameter with UDP connections only, not with TCP connections. The default value is enabled. | [optional] [default to null]
**SipSessionTimeout** | **int64** | Indicates the timeout value for sip session. The default value is 300 seconds | [optional] [default to 300]
**InsertRecordRouteHeader** | **string** | Enables or disables the insertion of a Record-Route header, which indicates the next hop for the following SIP request messages. The default value is disabled. | [optional] [default to null]
**Community** | **string** | A string that groups SIP profiles together if dialog-aware is enabled. | [optional] [default to null]
**MaxSize** | **int64** | Specifies the maximum SIP message size that the BIG-IP system accepts. The default value is 65535 bytes. | [optional] [default to 65535]
**SecureViaHeader** | **string** | Enables or disables the insertion of a Secure Via header, which indicates where the message originated. When you are using SSL/TLS (over TCP) to create a secure channel with the server node, use this setting to configure the BIG-IP system to insert a Secure Via header into SIP requests. The default value is disabled. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**UserViaHeader** | **string** | The header that is inserted if insert-via-header is enabled. | [optional] [default to null]
**MaxMediaSessions** | **int64** | Specifies the maximum number of SDP media sessions that the BIG-IP system accepts. The default value is 6. | [optional] [default to 6]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**MaxSessionsPerRegistration** | **int64** | Specifies the maximum number of calls or sessions can be made by a user for a single registration that the BIG-IP system accepts. The default value is 50. | [optional] [default to 50]
**Name** | **string** | Name of entity | [optional] [default to null]
**LogPublisher** | **string** | Configures the log publisher that handles events logging for this profile. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**RtpProxyStyle** | **string** | Indicates the style in which the RTP will proxy the data. When a dialog is established, the necessary SDP data needs to know where RTP flows are directed. It is provided with three options 1.) Use of a bidirectional related flow (symmetric, the default option), 2.) Use of ephemeral listeners to support fixed client IP, listener is restricted to connections coming from a particular source (restricted-by-ip-address option) and 3.) Use of ephemeral listeners to support wildcard, connections are allowed to come from anyway (any-location option) . | [optional] [default to null]
**DialogEstablishmentTimeout** | **int64** | Indicates the timeout value for dialog establishment. The default value is 10 seconds. | [optional] [default to 10]
**DialogAware** | **string** | Indicates whether snooping of SIP dialogs is enabled. | [optional] [default to null]
**InsertViaHeader** | **string** | Enables or disables the insertion of a Via header, which indicates where the message originated. The response message uses this routing information. The default value is disabled. | [optional] [default to null]
**Security** | **string** | Enables the use of enhanced SIP security checking. | [optional] [default to null]
**EnableSipFirewall** | **string** | Indicates whether SIP firewall capability is enabled. Default value is no | [optional] [default to null]
**LogProfile** | **string** | Configures the ALG log profile that controls logging. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


