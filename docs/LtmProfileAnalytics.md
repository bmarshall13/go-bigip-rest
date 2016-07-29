# LtmProfileAnalytics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationEmailAddresses** | **string** | Specifies which email addresses receive alerts by email when notification-by-email is enabled. | [optional] [default to null]
**CollectPageLoadTime** | **string** | Enables or disables the collection of the page load time statistics. | [optional] [default to null]
**NotificationBySnmp** | **string** | Enables or disables sending the analytics alerts by SNMP traps (notification-by-syslog must be enabled). | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. | [optional] [default to null]
**CapturedTrafficInternalLogging** | **string** | Enables or disables the internal logging of captured traffic. | [optional] [default to null]
**CollectResponseCodes** | **string** | Enables or disables the collection of HTTP response codes returned by the servers. | [optional] [default to null]
**CollectServerLatency** | **string** | Enables or disables the collection of server latency statistics. Deprecated: server latency is always collected. | [optional] [default to null]
**CapturedTrafficExternalLogging** | **string** | Enables or disables the external logging of captured traffic. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**CollectMethods** | **string** | Enables or disables the collection of HTTP methods statistics. | [optional] [default to null]
**SessionTimeout** | **int64** |  | [optional] [default to 1800]
**CollectIp** | **string** | Enables or disables the collection of client IPs statistics. | [optional] [default to null]
**CollectGeo** | **string** | Enables or disables the collection of the names of the countries from where the traffic was sent. | [optional] [default to null]
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**SessionTimeoutMinutes** | **string** | Specifies the number of minutes of user non-activity before the system considers the session to be over. | [optional] [default to null]
**PublishIruleStatistics** | **string** | Enables or disables publishing analytics statistics for iRules. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ExternalLoggingPublisher** | **string** | Specifies the external logging publisher used to send statistical data to one or more destinations. | [optional] [default to null]
**CollectedStatsExternalLogging** | **string** | Enables or disables the external logging of the collected statistics. | [optional] [default to null]
**NotificationBySyslog** | **string** | Enables or disables logging of the analytics alerts into the Syslog. | [optional] [default to null]
**CollectMaxTpsAndThroughput** | **string** | Enables or disables the collection of maximum TPS and throughput for all collected entities. | [optional] [default to null]
**SmtpConfig** | **string** | The SMTP configuration to be used with analytics | [optional] [default to null]
**Sampling** | **string** | Enables or disables transaction sampling. This attribute can be set in the default profile only. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**CollectUrl** | **string** | Enables or disables the collection of URL statistics. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**CollectSubnets** | **string** | Enables or disables the collection of client side subnets. | [optional] [default to null]
**SamplingRatio** | **int64** |  | [optional] [default to 2]
**CollectUserSessions** | **string** | Enables or disables the collection of user sessions. | [optional] [default to null]
**NotificationByEmail** | **string** | Enables or disables sending the analytics alerts by email. | [optional] [default to null]
**CollectUserAgent** | **string** | Enables or disables the collection of user agents statistics. | [optional] [default to null]
**CollectHttpThroughput** | **string** | Enables or disables the collection of throughput statistics. Deprecated: HTTP throughput is always collected. | [optional] [default to null]
**CollectedStatsInternalLogging** | **string** | Enables or disables the internal logging of the collected statistics. | [optional] [default to null]
**SessionCookieSecurity** | **string** | Specifies the condition for adding a secure attribute to the session cookie. The default value is ssl-only. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


