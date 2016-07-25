# ProfileRequestLog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**ResponseLogErrorProtocol** | **string** | HighSpeedLogging protocol to use when logging. | [optional] [default to null]
**LogRequestLoggingErrors** | **string** | Enables secondary logging should the primary lack sufficient available bandwidth.  This mechanism is best used to send an alert to a completely separate destination. | [optional] [default to null]
**RequestLogTemplate** | **string** | The template to use when generating log messages. Shell style escapes (eg $foo and/or ${foo}) are used to import transaction-specific values. | [optional] [default to null]
**RequestLogErrorPool** | **string** | Name of the pool from which to select log servers. | [optional] [default to null]
**ResponseLogPool** | **string** | Name of the pool from which to select log servers. | [optional] [default to null]
**LogResponseByDefault** | **string** | Response logging may be overridden via iRule. This field determines the default response action. | [optional] [default to null]
**ResponseLogTemplate** | **string** | The template to use when generating log messages. Shell style escapes (eg $foo and/or ${foo}) are used to import transaction-specific values. | [optional] [default to null]
**ProxyCloseOnError** | **string** | If enabled, the logging profile will close the connection after sending its proxy-response. | [optional] [default to null]
**RequestLogProtocol** | **string** | HighSpeedLogging protocol to use when logging. | [optional] [default to null]
**LogResponseLoggingErrors** | **string** | Enables secondary logging should the primary lack sufficient available bandwidth.  This mechanism is best used to send an alert to a completely separate destination. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**ResponseLogErrorTemplate** | **string** | The template to use when generating log messages. Shell style escapes (eg $foo and/or ${foo}) are used to import transaction-specific values. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**ResponseLogErrorPool** | **string** | Name of the pool from which to select log servers. | [optional] [default to null]
**ProxyRespondOnLoggingError** | **string** | Should the logging fail, this feature allows the logging profile to respond directly, eg with an HTTP 502. | [optional] [default to null]
**RequestLogPool** | **string** | Name of the pool from which to select log servers. | [optional] [default to null]
**RequestLogErrorTemplate** | **string** | The template to use when generating log messages. Shell style escapes (eg $foo and/or ${foo}) are used to import transaction-specific values. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**ProxyResponse** | **string** | The response to send on logging errors. | [optional] [default to null]
**RequestLogging** | **string** | Enables or disables logging before the request is proxied. | [optional] [default to null]
**ResponseLogProtocol** | **string** | HighSpeedLogging protocol to use when logging. | [optional] [default to null]
**RequestLogErrorProtocol** | **string** | HighSpeedLogging protocol to use when logging. | [optional] [default to null]
**ResponseLogging** | **string** | Enables or disables logging before the response is returned to the client. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


