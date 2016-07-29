# NetIpsecTrafficSelector

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**DestinationAddress** | **string** | Specifies the destination IP address of the traffic to be matched. | [optional] [default to null]
**IpsecPolicy** | **string** | Specifies the name of the IPsec policy to be enforced on the matched traffic. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**Direction** | **string** | Specifies the direction of traffic to be protected with IPsec. If the direction is both, use source-address and destination-address with respect to the outbound direction. The default value is both. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Action** | **string** | Specifies how the system handles traffic that matches the criteria in the traffic selector. Only protect is currently supported. | [optional] [default to null]
**DestinationPort** | **int64** | Specifies the destination port for which you want the traffic to be matched. | [optional] [default to null]
**SourcePort** | **int64** | Specifies the source port for which you want the traffic to be matched. | [optional] [default to null]
**IpProtocol** | **int64** | Specifies the IP protocol for which you want the traffic to be matched. | [optional] [default to 255]
**SourceAddress** | **string** | Specifies the source IP address of the traffic to be matched. | [optional] [default to null]
**Order** | **int64** | Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


