# LtmMessageRoutingGenericRoute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**DestinationAddress** | **string** | Specifies the destination address of the route. If this attribute is not present, it will be treated as a wildcard matching all message destination-addresses. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**PeerSelectionMode** | **string** | Specifies the method to use when selecting a peer from the provided list of peers. Possible values are: sequential and ratio. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**SourceAddress** | **string** | Specifies the source address of the route. If this attribute is not present, it will be treated as a wildcard matching all message source-addresses. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


