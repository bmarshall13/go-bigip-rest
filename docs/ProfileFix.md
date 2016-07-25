# ProfileFix

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**FullLogonParsing** | **string** | Enable or disable logon message is always fully parsed | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ResponseParsing** | **string** | Enable or disable response parsing which parses the messages from FIX server | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**MessageLogPublisher** | **string** | Specifies the publisher for message logging | [optional] [default to null]
**QuickParsing** | **string** | Enable or disable quick parsing which parses the basic standard fields and validates message length and checksum | [optional] [default to null]
**SenderTagClass** | **string** | Specifies the tag substitution map between sender id and tag substitution data group | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**ReportLogPublisher** | **string** | Specifies the publisher for error message and status report | [optional] [default to null]
**ErrorAction** | **string** | Specifies the error handling method | [optional] [default to null]
**StatisticsSampleInterval** | **int64** | Specifies the sample interval in seconds of the message rate | [optional] [default to 20]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


