# ProfileRadius

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Clients** | **string** | The list of host and network addresses from which clients can connect. Any entry in the list can be a host or network address: 10.10.10.0/24 or 10.10.10.10. A value of \&quot;none\&quot; indicates that any client can connect. The default value is \&quot;none\&quot;. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**PemProtocolProfileRadius** | **string** | Specifies PEM protocol profile to be used when subscriber discovery is enabled. PEM protocol profile defines mapping of RADIUS AVPs to subscriber ID and other PEM subscriber session attributes. | [optional] [default to null]
**PersistAvp** | **string** | Specifies the name of the RADIUS attribute used to persist on. It can be an ASCII string or numeric code (1-255). Acceptable strings are from RFC 2865 section 5. A value of \&quot;none\&quot; indicates that persistence is disabled. The default value is \&quot;none\&quot;. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. The default value is radius. | [optional] [default to null]
**SubscriberDiscovery** | **string** | Enables PEM subscriber discovery based on the content of RADIUS packets. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


