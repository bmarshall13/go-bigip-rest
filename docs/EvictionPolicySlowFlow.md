# EvictionPolicySlowFlow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Throttling** | **string** | Specifies whether to kill flows (throttling enabled) when limits are exceeded. | [optional] [default to null]
**GracePeriod** | **int64** | Specifies the minimum age that a flow is allowed to be before it is considered slow. | [optional] [default to 10]
**Enabled** | **string** | Specifies whether to enable or disable slow flow detection. | [optional] [default to null]
**Maximum** | **int64** | Specifies the maximum number of flows, either as a count or a percentage, that are allowed to remain open. | [optional] [default to 100]
**ThresholdBps** | **int64** | Specifies the transfer rate threshold (in bytes per second) under which flows are considered slow. | [optional] [default to 32]
**EvictionType** | **string** | Specifies whether the slow flow maximum is an absolute count of slow flows, or a percentage of slow flows on the context. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


