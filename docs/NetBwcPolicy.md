# NetBwcPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**TrafficPriorityMap** | **string** | Policy Priority Traffic group to use as enforcement during congestion along with policy rate configured. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | Description for the policy. | [optional] [default to null]
**IpTos** | **string** | Specifies the Type of Service (ToS) level to use when outgoing packets cross threshold (Note: in such case, bwc just does only marking and not shaping). The default value is pass-through. | [optional] [default to null]
**LogPublisher** | **string** | Specifies the logging publisher. A new log-publisher object can be created via TMSH command tmsh create sys log-config publisher. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the bwc policy resides. | [optional] [default to null]
**Dynamic** | **string** | Policy type: If dynamic enabled or disabled. Policy type change modification is a disallowed configuration. | [optional] [default to null]
**LinkQos** | **string** | Specifies the Quality of Service (QoS) level to use when outgoing packets cross threshold (Note: in such case, bwc just does only marking and not shaping). The default value is pass-through. | [optional] [default to null]
**MaxRate** | **float32** | Maximum rate for this policy class. | [optional] [default to null]
**MaxUserRatePps** | **float32** | Maximum user pps for this policy class. | [optional] [default to null]
**Measure** | **string** | Enables/Disables bandwidth measurement for this policy. Available for only dynamic policies. | [optional] [default to null]
**LogPeriod** | **int64** | Specifies the frequency for bwc measurement to log to publisher. | [optional] [default to 2048]
**MaxUserRate** | **float32** | Maximum per user rate for this policy class. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


