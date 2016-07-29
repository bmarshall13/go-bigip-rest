# LtmMessageRoutingDiameterProfileSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**VendorId** | **int64** | Specifies the vendor identification number assigned to the diameter server by the Internet Assigned Numbers Authority (IANA). The default value is 3375. | [optional] [default to null]
**DestHostRewrite** | **string** | Specifies the value used in rewriting the Destination-Host AVP on egress. | [optional] [default to null]
**AuthApplicationId** | **int64** | Name of the authorization application | [optional] [default to null]
**OriginHost** | **string** | Specifies an identifier for the originating server, for example, siteserver.f5.com. | [optional] [default to null]
**PersistTimeout** | **int64** | Indicates the timeout value for persistence entries in seconds. The default value is 180 seconds. This value should be greater than transaction timeout specified in the Diameter router profile. The lesser value of the two is used to create the persist record upon receipt of the first Diameter request message response. This value is updated upon receipt of subsequent responses. | [optional] [default to 180]
**DestRealmRewrite** | **string** | Specifies the value used in rewriting the Destination-Realm AVP on egress. | [optional] [default to null]
**WatchdogTimeout** | **int64** | Specifies the number of seconds that a connection can be idle before a device watchdog request is sent. The default value is 0 seconds, indicating that a device watchdog request will not be sent to a client or server. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile from which this profile inherits settings. The default is the system-supplied \&quot;diametersession\&quot; profile. | [optional] [default to null]
**OriginRealmRewrite** | **string** | Specifies the value used in rewriting the Origin-Realm AVP on egress. | [optional] [default to null]
**ProductName** | **string** | Specifies the vendor-assigned name for the product. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**PersistType** | **string** | Specifies the type of persistence. The default is AVP.  Valid persistence types:  1) None: Disables persistence.  2) AVP: Enables persistence as determined by the AVP within the message.  3) Custom: Enables persistence as determined by a custom key specified within an iRule. | [optional] [default to null]
**MaxWatchdogFailures** | **int64** | Specifies the maximum number of device watchdog failures that the traffic management system can receive before it tears down the connection. After the system receives this number of device watchdog failures, it closes the connection. The default value is 10. | [optional] [default to 10]
**Partition** | **string** |  | [optional] [default to null]
**OriginHostRewrite** | **string** | Specifies the value used in rewriting the Origin-Host AVP on egress. | [optional] [default to null]
**ResetOnTimeout** | **string** | Specifies, when checked (enabled), and watchdog failures exceed the specified number of maximum watchdog failures, that the system resets the connection. The default value is enabled. | [optional] [default to null]
**MaxMessageSize** | **int64** | Specifies the maximum number of bytes allowed for a message. The default value is 0, indicating that this restriction does not apply. | [optional] [default to null]
**HandshakeTimeout** | **int64** | Specifies the number of seconds before the handshake to a peer times out. The default value is 10 seconds. | [optional] [default to 10]
**PersistAvp** | **string** | Specifies the expression for the session-key that identifies the Diameter AVP. The expression can be an ASCII string or a 32-bit numeric ID, ranging from 1 through 4294967295. The default value is \&quot;session-id\&quot;. A grouped-avp can be specified using the expression grouped-avp-name[index]:Nested-avp1[index1]:Nested-avp2[index2], where nested-avp1 and nested-avp2 are the AVPs in the grouped AVP. Instead of an AVP name, AVP code can also be specified. | [optional] [default to null]
**AcctApplicationId** | **int64** | Name of the account application | [optional] [default to null]
**OriginRealm** | **string** | Specifies an identifier for the originating realm, for example, f5. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


