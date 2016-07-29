# NetWccpServices

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**RedirectionMethod** | **string** | Specifies the method the router uses to redirect traffic: GRE or L2. The default value is gre. | [optional] [default to null]
**Protocol** | **string** | Specifies the protocol of the traffic to be redirected: TCP or UDP. The default value is tcp. | [optional] [default to null]
**HashFields** | **string** | Specifies to the router which traffic attributes to use to determine which BIG-IP system it should forward traffic to for load balancing: destination IP address (dest-ip), destination port (dest-port), source IP address (src-ip), and/or source port (src-port). | [optional] [default to null]
**Weight** | **int64** | Specifies the relative importance of this traffic in a load balancing environment. The range is from 1 to 100. The default value is 50. | [optional] [default to 50]
**Routers** | **string** | Specifies the IP addresses of the WCCP-enabled routers that redirect traffic. | [optional] [default to null]
**Password** | **string** | Specifies the password for MD5 authentication or none. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**ReturnMethod** | **string** | Specifies the method used to return pass-through traffic to the router: GRE or L2. The default value is gre. | [optional] [default to null]
**TunnelRemoteAddresses** | **string** | Specifies the Router Identifier IP address of the router that redirects traffic. | [optional] [default to null]
**Ports** | **int64** | Specifies one or more ports (up to 8) on which traffic is redirected. | [optional] [default to null]
**Priority** | **int64** | Specifies the precedence of the service group relative to the other service groups. The range is from 1 to 255. The default value is 100. | [optional] [default to 100]
**TunnelLocalAddress** | **string** | Specifies an IP address on the BIG-IP system to which the WCCP-enabled routers should redirect traffic. Specify a self IP address of an external VLAN on the BIG-IP system. | [optional] [default to null]
**PortType** | **string** | Specifies whether the WCCP interception of traffic is based on the destination port (dest) or source port (source), or is not specified (none). The default value is none. | [optional] [default to null]
**TrafficAssign** | **string** | Specifies whether load balancing is achieved by a hash algorithm or a mask. If you specify hash, specify one or more attributes using the option hash-fields. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


