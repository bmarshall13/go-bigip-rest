# MonitorDiameter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**VendorId** | **string** | The IANA SMI Network Management Private Enterprise Codes value assigned to the vendor of the Diameter application. | [optional] [default to null]
**Destination** | **string** |  | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**HostIpAddress** | **string** | The Host-IP-Address is used to inform a Diameter peer of the sender&#39;s IP address. | [optional] [default to null]
**UpInterval** | **int64** | Specifies how often in seconds that the system issues the monitor check when the node is up. The default value is the same as the (down) interval. | [optional] [default to null]
**VendorSpecificAcctApplicationId** | **string** | Used to advertise support of a vendor-specific Diameter Application.  Exactly one of the Auth-Application-Id and Acct-Application-Id AVPs MAY be present. | [optional] [default to null]
**AuthApplicationId** | **string** | Used to advertise support of the Authentication and Authorization portion of an application.  Exactly one of the Auth-Application-Id and Acct-Application-Id AVPs MAY be present. | [optional] [default to null]
**OriginHost** | **string** | Identifies the endpoint that originated the Diameter message. | [optional] [default to null]
**ManualResume** | **string** | Specifies whether the system automatically changes the status of a resource to up at the next successful monitor check. The default value of the manual-resume option is disabled. Note that if you set the manual-resume option to enabled, you must manually mark the resource as up before the system can use it for load balancing connections. | [optional] [default to null]
**VendorSpecificAuthApplicationId** | **string** | Used to advertise support of a vendor-specific Diameter Application.  Exactly one of the Auth-Application-Id and Acct-Application-Id AVPs MAY be present. | [optional] [default to null]
**Interval** | **int64** | Specifies the frequency at which the system issues the monitor check. The default value is 10 seconds. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the existing monitor from which the system imports settings for the new monitor. | [optional] [default to null]
**ProductName** | **string** | The vendor assigned name for the product.  The Product-Name SHOULD remain constant across firmware revisions for the same product. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the monitor resides. | [optional] [default to null]
**VendorSpecificVendorId** | **string** | The vendor id for a vendor-specific Diameter Application. | [optional] [default to null]
**TimeUntilUp** | **int64** | Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0. | [optional] [default to null]
**OriginRealm** | **string** | The Realm of the originator of any Diameter message | [optional] [default to null]
**Timeout** | **int64** | Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 31 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire. | [optional] [default to null]
**AcctApplicationId** | **string** | Used to advertise support of the Accounting portion of an application. Exactly one of the Auth-Application-Id and Acct-Application-Id AVPs MAY be present. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


