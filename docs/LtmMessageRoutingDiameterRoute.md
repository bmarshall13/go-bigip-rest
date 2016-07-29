# LtmMessageRoutingDiameterRoute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**PeerSelectionMode** | **string** | Specifies how the system selects a peer to handle the Diameter traffic for this static route. The default is Sequential. Options are  1) Ratio: Peer selection is based on the ratio that is set for each peer that is in the Selected list.  2) Sequential: Peer selection is based on the order of the peers in the Selected list. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**DestinationRealm** | **string** | Specifies the destination realm matching the Destination-Realm AVP value in the message. The blank default value routes all destination-realms. | [optional] [default to null]
**OriginRealm** | **string** | Specifies the origin realm matching the Origin-Realm AVP value in the message. The blank default value routes all origin-realms. | [optional] [default to null]
**VirtualServer** | **string** | Specifies the virtual server from which the system receives client requests for this static route. If you do not select a virtual server, the system uses this static route to route Diameter messages originating from any client. | [optional] [default to null]
**ApplicationId** | **int64** | Specifies the identifier matching the application-id in the Diameter message. A value of 0 matches every application-id. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


