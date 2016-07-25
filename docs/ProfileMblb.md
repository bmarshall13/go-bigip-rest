# ProfileMblb

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**ShutdownTimeout** | **int64** | Delays sending FIN when BIGIP receives the first FIN packet from either the client or the server. Value of 0 means send FIN immediately otherwise the minimum of tcp idle timeout and shutdown timeout is used. The default value is 5 seconds | [optional] [default to null]
**IsolateAbort** | **string** | whether to isolate abort event propagation | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**TagTtl** | **int64** | specifies the TTL (time to live) for TAGs | [optional] [default to null]
**IngressLow** | **int64** | specifies the low water mark for ingress message queue | [optional] [default to null]
**IsolateClient** | **string** | whether to isolate clientside shutdown event propagation | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**IngressHigh** | **int64** | specifies the high water mark for ingress message queue | [optional] [default to null]
**EgressLow** | **int64** | specifies the low water mark for egress message queue | [optional] [default to null]
**IsolateServer** | **string** | whether to isolate serverside shutdown event propagation | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**EgressHigh** | **int64** | specifies the high water mark for egress message queue | [optional] [default to null]
**MinConn** | **int64** | specifies the minimum number of serverside connections | [optional] [default to null]
**IsolateExpire** | **string** | whether to isolate expire event propagation | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


