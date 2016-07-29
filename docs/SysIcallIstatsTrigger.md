# SysIcallIstatsTrigger

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Repeat** | **int64** | Specifies the interval to repeat event generation while the stat matches the istats-key, in seconds. | [optional] [default to 0]
**Description** | **string** | User defined text about the configuration item. | [optional] [default to null]
**EventName** | **string** | Specifies the name of the event to be generated. | [default to null]
**RangeMin** | **int64** | Minimum integer value for the statistic being monitored for the trigger to fire. The default is -2147483647. | [optional] [default to -2147483647]
**IstatsKey** | **string** | Specifies the string of name-value pairs that define the statistic to monitor. | [default to null]
**Duration** | **int64** | How long the counter needs to be in range before the trigger fires, in seconds. | [optional] [default to 0]
**RangeMax** | **int64** | Maximum integer value for the statistic being monitored for the trigger to fire. The default is 2147483647. | [optional] [default to 2147483647]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


