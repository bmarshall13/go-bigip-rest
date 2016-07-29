# NetBwcPolicyCategories

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**TrafficPriorityMap** | **string** | Policy Priority Traffic group to use as enforcement during congestion along with category configured. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | Description for the policy category. | [optional] [default to null]
**IpTos** | **string** | Specifies the Type of Service (ToS) level to use when outgoing packets cross threshold (Note: in such case, bwc just does only marking and not shaping).  The default value is pass-through. | [optional] [default to null]
**LinkQos** | **string** | Specifies the Quality of Service (QoS) level to use when outgoing packets cross threshold (Note: in such case, bwc just does only marking and not shaping). . The default value is pass-through. | [optional] [default to null]
**MaxCatRatePercentage** | **int64** | Policy Category - Percentage of max-user-rate for this category. This becomes marking threshold when marking is done instead of enforcement. | [optional] [default to null]
**MaxCatRate** | **int64** | Policy Category - Maximum rate for this category. This becomes marking threshold when marking is done instead of enforcement. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


